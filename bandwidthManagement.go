package control

import "encoding/json"

type BMConditionType string

const (
	BMConditionTrafficType BMConditionType = "BMConditionTrafficType"
	BMConditionQuota       BMConditionType = "BMConditionQuota"
	BMConditionLargeData   BMConditionType = "BMConditionLargeData"
	BMConditionTrafficRule BMConditionType = "BMConditionTrafficRule"
	BMConditionContentRule BMConditionType = "BMConditionContentRule"
	BMConditionService     BMConditionType = "BMConditionService"
	BMConditionDscp        BMConditionType = "BMConditionDscp"
	BMConditionUsers       BMConditionType = "BMConditionUsers"
	BMConditionInvalid     BMConditionType = "BMConditionInvalid"
	BMContentRuleType      BMConditionType = "BMContentRuleType"
	BMConditionGuests      BMConditionType = "BMConditionGuests"
	BMConditionApplication BMConditionType = "BMConditionApplication"
)

type BMTrafficType string

const (
	BMTrafficEmail            BMTrafficType = "BMTrafficEmail"
	BMTrafficFtp              BMTrafficType = "BMTrafficFtp"
	BMTrafficInstantMessaging BMTrafficType = "BMTrafficInstantMessaging"
	BMTrafficMultimedia       BMTrafficType = "BMTrafficMultimedia"
	BMTrafficP2p              BMTrafficType = "BMTrafficP2p"
	BMTrafficRemoteAccess     BMTrafficType = "BMTrafficRemoteAccess"
	BMTrafficSip              BMTrafficType = "BMTrafficSip"
	BMTrafficVpn              BMTrafficType = "BMTrafficVpn"
	BMTrafficWeb              BMTrafficType = "BMTrafficWeb"
)

type BMCondition struct {
	Type BMConditionType `json:"type"`
	/*@{ TrafficRule, ContentRules */
	ValueId IdReference `json:"valueId"`
	/*@}*/
	/*@{ Service */
	Service IpServiceReference `json:"service"`
	/*@}*/
	/*@{ Dscp */
	Dscp int `json:"dscp"`
	/*@}*/
	/*@{ TrafficType */
	TrafficType BMTrafficType `json:"trafficType"`
	/*@}*/
	/*@{ Users */
	User UserReference `json:"user"` // @see UserManager
	/*@}*/
	/*@{ Application */
	AppId int `json:"appId"`
	/*@}*/
}

type BMConditionList []BMCondition

type BandwidthSetting struct {
	Enabled bool          `json:"enabled"`
	Value   int           `json:"value"`
	Unit    BandwidthUnit `json:"unit"`
}

type BMRule struct {
	Id               KId              `json:"id"`
	Enabled          bool             `json:"enabled"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Color            string           `json:"color"`
	Traffic          BMConditionList  `json:"traffic"`
	ReservedDownload BandwidthSetting `json:"reservedDownload"`
	ReservedUpload   BandwidthSetting `json:"reservedUpload"`
	MaximumDownload  BandwidthSetting `json:"maximumDownload"`
	MaximumUpload    BandwidthSetting `json:"maximumUpload"`
	InterfaceId      IdReference      `json:"interfaceId"`
	ValidTimeRange   IdReference      `json:"validTimeRange"`
	Chart            bool             `json:"chart"`
}

type BMRuleList []BMRule

type InternetBandwidthData struct {
	Speed int           `json:"speed"` ///> maximum speed of the link (defined in Interfaces); zero means "undefined"
	Unit  BandwidthUnit `json:"unit"`  ///> unit for the speed value
}

type InternetBandwidth struct {
	Id       KId                   `json:"id"`
	Name     string                `json:"name"`   ///> name of the interface
	Type     InterfaceType         `json:"type"`   ///> (e.g. ethernet, ras, etc.)
	Online   bool                  `json:"online"` ///> false = interface is offline (values download and upload should be ignored)
	Download InternetBandwidthData `json:"download"`
	Upload   InternetBandwidthData `json:"upload"`
}

type InternetBandwidthList []InternetBandwidth

type BandwidthManagementConfig struct {
	DecryptVpnTunnels bool       `json:"decryptVpnTunnels"` ///>Traffic in VPN tunnels will be matched against rules decrypted
	Rules             BMRuleList `json:"rules"`
}

// BandwidthManagementGet - Get the list of Bandwidth Management rules
// Return
//	config - Bandwidth Management rules
func (s *ServerConnection) BandwidthManagementGet() (*BandwidthManagementConfig, error) {
	data, err := s.CallRaw("BandwidthManagement.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config BandwidthManagementConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// BandwidthManagementSet - Stores the list of Bandwidth Management rules
// Parameters
//	config - Bandwidth Management rules
// Return
//	errors - list of errors \n
func (s *ServerConnection) BandwidthManagementSet(config BandwidthManagementConfig) (ErrorList, error) {
	params := struct {
		Config BandwidthManagementConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("BandwidthManagement.set", params)
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

// BandwidthManagementGetBandwidth - Returns list of Internet interfaces and their current usage
// Return
//	list - list of interfaces (sorted by name); empty if there are no Internet interfaces
func (s *ServerConnection) BandwidthManagementGetBandwidth() (InternetBandwidthList, error) {
	data, err := s.CallRaw("BandwidthManagement.getBandwidth", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List InternetBandwidthList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// BandwidthManagementSetBandwidth - Stores the list of Bandwidth Management rules
// Parameters
//	list - list of Bandwidth Management rules
// Return
//	errors - list of errors \n
func (s *ServerConnection) BandwidthManagementSetBandwidth(list InternetBandwidthList) (ErrorList, error) {
	params := struct {
		List InternetBandwidthList `json:"list"`
	}{list}
	data, err := s.CallRaw("BandwidthManagement.setBandwidth", params)
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
