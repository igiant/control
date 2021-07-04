package control

import "encoding/json"

type DynamicDnsStatus string

const (
	DynamicDnsOk     DynamicDnsStatus = "DynamicDnsOk"
	DynamicDnsError  DynamicDnsStatus = "DynamicDnsError"
	DynamicDnsUpdate DynamicDnsStatus = "DynamicDnsUpdate"
)

type DynamicDnsAddressType string

const (
	DynamicDnsAdressIface  DynamicDnsAddressType = "DynamicDnsAdressIface"
	DynamicDnsAdressDetect DynamicDnsAddressType = "DynamicDnsAdressDetect"
	DynamicDnsAdressCustom DynamicDnsAddressType = "DynamicDnsAdressCustom"
)

type DynamicDnsConfig struct {
	Enabled     bool                  `json:"enabled"`
	Provider    string                `json:"provider"`
	Hostname    string                `json:"hostname"`
	Credentials CredentialsConfig     `json:"credentials"`
	AddressType DynamicDnsAddressType `json:"addressType"`
	CustomIface IdReference           `json:"customIface"`
}

// DynamicDnsGet - Returns DynDNS configuration
// Return
//	config - configuration values
func (s *ServerConnection) DynamicDnsGet() (*DynamicDnsConfig, error) {
	data, err := s.CallRaw("DynamicDns.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config DynamicDnsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// DynamicDnsSet - Stores DynDNS configuration
// Parameters
//	config - configuration values
func (s *ServerConnection) DynamicDnsSet(config DynamicDnsConfig) error {
	params := struct {
		Config DynamicDnsConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("DynamicDns.set", params)
	return err
}

// DynamicDnsUpdate - Performs synchronous DynDNS update
func (s *ServerConnection) DynamicDnsUpdate() error {
	_, err := s.CallRaw("DynamicDns.update", nil)
	return err
}

// DynamicDnsGetStatus - Returns status of DynDNS update
// Return
//	message - list of errors
//	status - actual status of DynDNS update
func (s *ServerConnection) DynamicDnsGetStatus() (*LocalizableMessage, *DynamicDnsStatus, error) {
	data, err := s.CallRaw("DynamicDns.getStatus", nil)
	if err != nil {
		return nil, nil, err
	}
	message := struct {
		Result struct {
			Message LocalizableMessage `json:"message"`
			Status  DynamicDnsStatus   `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &message)
	return &message.Result.Message, &message.Result.Status, err
}

// DynamicDnsGetProviders - Returns list of supported providers
// Return
//	providers - return values
func (s *ServerConnection) DynamicDnsGetProviders() (StringList, error) {
	data, err := s.CallRaw("DynamicDns.getProviders", nil)
	if err != nil {
		return nil, err
	}
	providers := struct {
		Result struct {
			Providers StringList `json:"providers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &providers)
	return providers.Result.Providers, err
}
