package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	liquidweb "github.com/liquidweb/liquidweb-go"
)

// Config is the configuration for the API client.
type Config struct {
	Username  string
	Password  string
	URL       *url.URL
	Timeout   int
	SecureTLS bool
}

// NewConfig builds a new Config.
func NewConfig(username string, password string, apiURL string, timeout int, secureTLS bool) (*Config, error) {
	if len(username) == 0 {
		return nil, fmt.Errorf("lwApi.username is missing from config")
	}
	if len(password) == 0 {
		return nil, fmt.Errorf("lwApi.password is missing from config")
	}
	if len(apiURL) == 0 {
		return nil, fmt.Errorf("lwApi.url is missing from config")
	}

	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	apiTimeout := timeout
	// Set a default timeout if not set.
	if apiTimeout == 0 {
		apiTimeout = 20
	}

	config := &Config{
		Username:  username,
		Password:  password,
		URL:       parsedURL,
		Timeout:   apiTimeout,
		SecureTLS: secureTLS,
	}

	return config, nil
}

// Client provides the HTTP backend.
type Client struct {
	config     *Config
	httpClient *http.Client
}

// NewClient returns a prepared API client.
func NewClient(config *Config) *Client {
	httpClient := &http.Client{Timeout: time.Duration(config.Timeout) * time.Second}

	if !config.SecureTLS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// Call takes a path, such as "network/zone/details" and a params structure.
// It is recommended that the params be a map[string]interface{}, but you can use
// anything that serializes to the right json structure.
// A `interface{}` and an error are returned, in typical go fashion.
//
// Example:
//	args := map[string]interface{}{
//		"uniq_id": "ABC123",
//	}
//	got, gotErr := apiClient.Call("bleed/asset/details", args)
//	if gotErr != nil {
//		panic(gotErr)
//	}
func (client *Client) Call(method string, params interface{}, into interface{}) error {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return err
	}

	var raw map[string]interface{}
	if err = json.Unmarshal(bsRb, &raw); err != nil {
		return err
	}
	errorClass, ok := raw["error_class"]
	if ok {
		errorClassStr := errorClass.(string)
		if errorClassStr != "" {
			return liquidweb.LWAPIError{
				ErrorClass:   errorClassStr,
				ErrorFullMsg: raw["full_message"].(string),
				ErrorMsg:     raw["error"].(string),
			}
		}
	}

	// Response should be valid, decode it.
	return json.Unmarshal(bsRb, &into)
}

// CallRaw is just like Call, except it returns the raw json as a byte slice. However, in contrast to
// Call, CallRaw does *not* check the API response for LiquidWeb specific exceptions as defined in
// the type LWAPIError. As such, if calling this function directly, you must check for LiquidWeb specific
// exceptions yourself.
//
// Example:
//	args := map[string]interface{}{
//		"uniq_id": "ABC123",
//	}
//	got, gotErr := apiClient.CallRaw("bleed/asset/details", args)
//	if gotErr != nil {
//		panic(gotErr)
//	}
//	// Check got now for LiquidWeb specific exceptions, as described above.
func (client *Client) CallRaw(method string, params interface{}) ([]byte, error) {
	//  api wants the "params" prefix key. Do it here so consumers dont have
	// to do this every time.
	args := map[string]interface{}{
		"params": params,
	}
	encodedArgs, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	// formulate the HTTP POST request
	uri := fmt.Sprintf("%s/%s", client.config.URL, method)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(encodedArgs))
	if err != nil {
		return nil, err
	}

	// HTTP basic auth
	req.SetBasicAuth(client.config.Username, client.config.Password)

	// make the POST request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad HTTP response code [%d] from [%s]", resp.StatusCode, uri)
	}

	// read the response body into a byte slice
	return ioutil.ReadAll(resp.Body)
}
