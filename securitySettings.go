package control

import "encoding/json"

// AntiSpoofingConfig - Miscellaneous part
type AntiSpoofingConfig struct {
	Enabled bool `json:"enabled"`
	Log     bool `json:"log"`
}

type UpnpConfig struct {
	Enabled        bool `json:"enabled"`
	LogPackets     bool `json:"logPackets"`
	LogConnections bool `json:"logConnections"`
}

type IPv6Config struct {
	BlockNative           bool           `json:"blockNative"`
	BlockTunnels          bool           `json:"blockTunnels"`
	AddressGroupException OptionalEntity `json:"addressGroupException"`
}

type DhcpScopes struct {
	BlockUnknownIp bool           `json:"blockUnknownIp"`
	Log            bool           `json:"log"`
	BlockScopeId   OptionalString `json:"blockScopeId"`
}

// ConnectionLimit - @deprecated use ConnLimitSettings instead
type ConnectionLimit struct {
	Enabled    bool                `json:"enabled"`
	Soft       int                 `json:"soft"`
	Value      int                 `json:"value"` // hard
	Rate       OptionalLong        `json:"rate"`
	Exclusions OptionalIdReference `json:"exclusions"`
	ExclSoft   int                 `json:"exclSoft"`
	ExclValue  int                 `json:"exclValue"` // hard
	ExclRate   OptionalLong        `json:"exclRate"`
}

type MiscSettingsConfig struct {
	AntiSpoofing AntiSpoofingConfig `json:"antiSpoofing"`
	// [Deprecated]
	ConnectionLimit ConnectionLimit `json:"connectionLimit"` // @deprecated use ConnLimit instead
	Upnp            UpnpConfig      `json:"upnp"`
	Ipv6            IPv6Config      `json:"ipv6"`
	DhcpScopes      DhcpScopes      `json:"dhcpScopes"`
}

// MacFilterActionType - Mac Filter part
type MacFilterActionType string

const (
	MacFilterDeny  MacFilterActionType = "MacFilterDeny"
	MacFilterAllow MacFilterActionType = "MacFilterAllow"
)

type MacAccessItem struct {
	MacAddress  string `json:"macAddress"`
	Description string `json:"description"`
}

type MacAccessList []MacAccessItem

type MacFilterConfig struct {
	Enabled         bool                `json:"enabled"`
	Interfaces      IdReferenceList     `json:"interfaces"`
	MacFilterAction MacFilterActionType `json:"macFilterAction"`
	MacAccessItems  MacAccessList       `json:"macAccessItems"`
	AllowReserved   bool                `json:"allowReserved"`
}

// ZeroConfigItemType - Zero-config network
type ZeroConfigItemType string

const (
	ZeroConfigVpnClients ZeroConfigItemType = "ZeroConfigVpnClients"
	ZeroConfigVpnTunnel  ZeroConfigItemType = "ZeroConfigVpnTunnel"
	ZeroConfigInterface  ZeroConfigItemType = "ZeroConfigInterface"
)

// ZeroConfigItem - One VPN in rule
type ZeroConfigItem struct {
	Type ZeroConfigItemType `json:"type"`
	Item IdReference        `json:"item"`
}

type ZeroConfigList []ZeroConfigItem

type ZeroConfigNetwork struct {
	Enabled bool           `json:"enabled"`
	Items   ZeroConfigList `json:"items"`
}

type SecuritySettingsConfig struct {
	MiscSettings      MiscSettingsConfig `json:"miscSettings"`
	MacFilter         MacFilterConfig    `json:"macFilter"`
	ZeroConfigNetwork ZeroConfigNetwork  `json:"zeroConfigNetwork"`
}

// SecuritySettingsGet - Returns actual values for Security Settings configuration in WebAdmin
// Return
//	config - structure containig security settings such as macfilter action, name, mac list and belonging interfaces.
func (s *ServerConnection) SecuritySettingsGet() (*SecuritySettingsConfig, error) {
	data, err := s.CallRaw("SecuritySettings.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config SecuritySettingsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// SecuritySettingsSet - Stores Security Settings configuration
// Parameters
//	config - structure containing security settings to be stored such as mac-filter action, name, mac list and belonging interfaces.
// Return
//	errors - list of errors \n
func (s *ServerConnection) SecuritySettingsSet(config SecuritySettingsConfig) (ErrorList, error) {
	params := struct {
		Config SecuritySettingsConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("SecuritySettings.set", params)
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
