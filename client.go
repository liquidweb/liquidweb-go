package storm

import "net/http"

// A Client holds the packages *viper.Viper and *http.Client. To get a *Client, call New.
type Client struct {
	config     *viper.Viper
	httpClient *http.Client
}
