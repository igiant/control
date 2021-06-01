package control

import "encoding/json"

type WanInterfaceConfig struct {
	Id                  KId                `json:"id"` // not used on Box
	Encap               InterfaceEncapType `json:"encap"`
	Mode                InterfaceModeType  `json:"mode"`
	Ip                  IpAddress          `json:"ip"`
	SubnetMask          IpAddress          `json:"subnetMask"`
	GatewayAutodetected bool               `json:"gatewayAutodetected"`
	Gateway             IpAddress          `json:"gateway"`
	DnsAutodetected     bool               `json:"dnsAutodetected"`
	DnsServers          string             `json:"dnsServers"`          // ipaddr;ipaddr
	LoadBalancingWeight OptionalLong       `json:"loadBalancingWeight"` // balancing
	Credentials         CredentialsConfig  `json:"credentials"`         // for Pppoe mode
}

type WanInterfaceConfigList []WanInterfaceConfig

type LanInterfaceConfig struct {
	Id                KId       `json:"id"` // not used on Box
	Ip                IpAddress `json:"ip"`
	SubnetMask        IpAddress `json:"subnetMask"`
	DhcpServerEnabled bool      `json:"dhcpServerEnabled"`
}

type ConnectivityAssistantConfig struct {
	Type ConnectivityType       `json:"type"` // only Persistent, Failover, LoadBalancing
	Wans WanInterfaceConfigList `json:"wans"`
	Lan  LanInterfaceConfig     `json:"lan"`
	// Wifi WifiConfig             `json:"wifi"`
}

// ConnectivityAssistantSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - input configuration
//	revertTimeout - If client doesn't confirm config to this timeout, configuration is reverted (0 - revert disabled)
// Return
//	errors - list of errors \n
func (s *ServerConnection) ConnectivityAssistantSet(config ConnectivityAssistantConfig, revertTimeout int) (ErrorList, error) {
	params := struct {
		Config        ConnectivityAssistantConfig `json:"config"`
		RevertTimeout int                         `json:"revertTimeout"`
	}{config, revertTimeout}
	data, err := s.CallRaw("ConnectivityAssistant.set", params)
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
