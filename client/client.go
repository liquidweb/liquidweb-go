package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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
	httpClient := &http.Client{Timeout: time.Duration(time.Duration(config.Timeout) * time.Second)}

	if !config.SecureTLS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient.Transport = tr
	}
	client := &Client{
		config:     config,
		httpClient: httpClient}

	return client
}

// Call takes a path, such as "network/zone/details" and a params structure.
// It is recommended that the params be a map[string]interface{}, but you can use
// anything that serializes to the right json structure.
// A `interface{}` and an error are returned, in typical go fasion.
//
// Example:
//	args := map[string]interface{}{
//		"uniq_id": "ABC123",
//	}
//	got, gotErr := apiClient.Call("bleed/asset/details", args)
//	if gotErr != nil {
//		panic(gotErr)
//	}
func (client *Client) Call(method string, params interface{}) (interface{}, error) {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return nil, err
	}

	// json decode into interface
	var decodedResp interface{}
	if jsonDecodeErr := json.Unmarshal(bsRb, &decodedResp); jsonDecodeErr != nil {
		return nil, jsonDecodeErr
	}
	mapDecodedResp, ok := decodedResp.(map[string]interface{})
	if !ok {
		return nil, errors.New("endpoint did not return the expected JSON structure")
	}
	errorClass, ok := mapDecodedResp["error_class"]
	if ok {
		errorClassStr := errorClass.(string)
		if errorClassStr != "" {
			return nil, LWAPIError{
				ErrorClass:   errorClassStr,
				ErrorFullMsg: mapDecodedResp["full_message"].(string),
				ErrorMsg:     mapDecodedResp["error"].(string),
			}
		}
	}
	// no LW errors so return the decoded response
	return decodedResp, nil
}

// CallInto is like call, but instead of returning an interface you pass it a
// struct which is filled, much like the json.Unmarshal function.  The struct
// you pass must satisfy the LWAPIRes interface.  If you embed the LWAPIError
// struct from this package into your struct, this will be taken care of for you.
//
// Example:
//	type ZoneDetails struct {
//		lwApi.LWAPIError
//		AvlZone     string   `json:"availability_zone"`
//		Desc        string   `json:"description"`
//		GatewayDevs []string `json:"gateway_devices"`
//		HvType      string   `json:"hv_type"`
//		ID          int      `json:"id"`
//		Legacy      int      `json:"legacy"`
//		Name        string   `json:"name"`
//		Status      string   `json:"status"`
//		SourceHVs   []string `json:"valid_source_hvs"`
//	}
//	var zone ZoneDetails
//	err = apiClient.CallInto("network/zone/details", paramers, &zone)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Got struct %#v\n", zone)
//
func (client *Client) CallInto(method string, params interface{}, into LWAPIRes) error {
	bsRb, err := client.CallRaw(method, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bsRb, into)
	if err != nil {
		return err
	}

	if into.HadError() {
		// the LWAPIRes satisfies the Error interface, so we can just return it on
		// error.
		return into
	}

	return nil
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
	// to do this everytime.
	args := map[string]interface{}{
		"params": params,
	}
	encodedArgs, encodeErr := json.Marshal(args)
	if encodeErr != nil {
		return nil, encodeErr
	}
	// formulate the HTTP POST request
	url := fmt.Sprintf("%s/%s", client.config.URL, method)
	req, reqErr := http.NewRequest("POST", url, bytes.NewReader(encodedArgs))
	if reqErr != nil {
		return nil, reqErr
	}
	// HTTP basic auth
	req.SetBasicAuth(client.config.Username, client.config.Password)
	// make the POST request
	resp, doErr := client.httpClient.Do(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad HTTP response code [%d] from [%s]", resp.StatusCode, url)
	}
	// read the response body into a byte slice
	bsRb, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return bsRb, nil
}
