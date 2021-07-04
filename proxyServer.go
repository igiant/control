package control

import "encoding/json"

type ParentProxyConfig struct {
	Enabled     bool              `json:"enabled"`
	Server      string            `json:"server"`
	Port        int               `json:"port"`
	AuthEnabled bool              `json:"authEnabled"`
	Credentials CredentialsConfig `json:"credentials"`
}

type ProxyServerConfig struct {
	Enabled                bool              `json:"enabled"`
	Port                   int               `json:"port"`
	AllowAllPorts          bool              `json:"allowAllPorts"`
	ParentProxy            ParentProxyConfig `json:"parentProxy"`
	AutomaticScriptDirect  bool              `json:"automaticScriptDirect"`
	AutomaticScriptEnabled bool              `json:"automaticScriptEnabled"`
}

// ProxyServerGet - Gets Proxy server configuration
// Return
//	config - current configuration
func (s *ServerConnection) ProxyServerGet() (*ProxyServerConfig, error) {
	data, err := s.CallRaw("ProxyServer.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ProxyServerConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ProxyServerSet - Sets Proxy server configuration
//	config - new configuration
func (s *ServerConnection) ProxyServerSet(config ProxyServerConfig) error {
	params := struct {
		Config ProxyServerConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ProxyServer.set", params)
	return err
}
