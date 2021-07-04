package control

import "encoding/json"

type HttpsServerMode string

const (
	HttpsServerModeDisabled           HttpsServerMode = "HttpsServerModeDisabled"
	HttpsServerModeDefaultCertificate HttpsServerMode = "HttpsServerModeDefaultCertificate"
	HttpsServerModeCustomCertificate  HttpsServerMode = "HttpsServerModeCustomCertificate"
)

type ReverseProxyRule struct {
	/*@{ server match */
	ServerHostname    string          `json:"serverHostname"`
	ServerHttp        bool            `json:"serverHttp"`           // server on standard HTTP port 80
	HttpsMode         HttpsServerMode `json:"httpsMode;//"`         // server on standard HTTPS port 443
	CustomCertificate IdReference     `json:"customCertificate;//"` // HTTPS server certificate ID
	/*@{ target connection */
	TargetServer string `json:"targetServer;//"` // hostname/IPv4/IPv6 + optionally ":<port>"
	TargetHttps  bool   `json:"targetHttps"`
	/*@{ actions */
	Antivirus bool `json:"antivirus"`
	/*@{ misc */
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	Id          KId    `json:"id"`
}

type ReverseProxyRuleList []ReverseProxyRule

type ReverseProxyConfig struct {
	Enabled            bool                 `json:"enabled"`
	DefaultCertificate IdReference          `json:"defaultCertificate"`
	Rules              ReverseProxyRuleList `json:"rules"`
}

// ReverseProxyGet - Get ReverseProxy config
// Return
//	config - reverse proxy config (enabled, default cert.)
func (s *ServerConnection) ReverseProxyGet() (*ReverseProxyConfig, error) {
	data, err := s.CallRaw("ReverseProxy.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ReverseProxyConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ReverseProxySet - Set ReverseProxy config
// Parameters
//	config - reverse proxy config (enabled, default cert.)
// Return
//	errors - list of errors TODO Write particular values
func (s *ServerConnection) ReverseProxySet(config ReverseProxyConfig) (ErrorList, error) {
	params := struct {
		Config ReverseProxyConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("ReverseProxy.set", params)
	if err != nil {
		return nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, err
}
