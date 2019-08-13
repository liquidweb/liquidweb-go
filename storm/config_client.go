package storm

import liquidweb "git.liquidweb.com/masre/liquidweb-go"

// ConfigBackend is the interface for storm configs.
type ConfigBackend interface {
	Details() (*Config, error)
	List(ConfigListParams) (*ConfigList, error)
}

// ConfigClient is the API client for storm configs.
type ConfigClient struct {
	Backend liquidweb.Backend
}

// Details fetches the details for a storm config.
func (c *ConfigClient) Details() (*Config, error) {
	return &Config{}, nil
}

// List fetches a list of storm configs.
func (c *ConfigClient) List(params ConfigListParams) (*ConfigList, error) {
	configList := &ConfigList{}

	err := c.Backend.Call("v1/Storm/Config/list", params, configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}
