package control

type RouteType string
const (
	RouteSystem RouteType = "RouteSystem" 
	RouteStatic RouteType = "RouteStatic" 
	RouteVpn RouteType = "RouteVpn" 
)

type Route struct {
	Enabled bool `json:"enabled"` // valid only for RouteStatic
	Name string `json:"name"` // valid only for RouteStatic
	Type RouteType `json:"type"` // Type defines Descriptions: RouteSystem: 'System route'
	//               RouteVpn:  'VPN route'
	Network IpAddress `json:"network"` 
	Mask IpAddress `json:"mask"` // used for IPv4
	PrefixLen int `json:"prefixLen"` // used for IPv6
	Gateway IpAddress `json:"gateway"` 
	InterfaceType InterfaceType `json:"interfaceType"` // @see InterfaceManager, used values: Ethernet, Ras, VpnTunnel, VpnServer
	InterfaceId IdReference `json:"interfaceId"` // invalid - interface is no more in the configuration
	// @note: not used for interfaceType == VpnServer (it means 'VPN Server' in gui)
	Metric int `json:"metric"` 
}

type RouteList []Route


// RoutingTableGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	routes - a list of routes currently stroed and used by system.
func (s *ServerConnection) RoutingTableGet(boolean [Opt]) (RouteList, error) {
	params := struct {
		Boolean [Opt] `json:"boolean"`
	}{boolean}
	data, err := s.CallRaw("RoutingTable.get", params)
	if err != nil {
		return nil, err
	}
	routes := struct {
		Result struct {
			Routes RouteList `json:"routes"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &routes)
	return routes.Result.Routes, err
}


// RoutingTableGetStaticRoutes - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	routes - a list of routes currently stroed and used by Control.
func (s *ServerConnection) RoutingTableGetStaticRoutes(boolean [Opt]) (RouteList, error) {
	params := struct {
		Boolean [Opt] `json:"boolean"`
	}{boolean}
	data, err := s.CallRaw("RoutingTable.getStaticRoutes", params)
	if err != nil {
		return nil, err
	}
	routes := struct {
		Result struct {
			Routes RouteList `json:"routes"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &routes)
	return routes.Result.Routes, err
}

type UpnpConfig struct {
	Enabled bool `json:"enabled"` 
	LogPackets bool `json:"logPackets"` 
	LogConnections bool `json:"logConnections"` 
}

type IPv6Config struct {
	BlockNative bool `json:"blockNative"` 
	BlockTunnels bool `json:"blockTunnels"` 
	AddressGroupException OptionalEntity `json:"addressGroupException"` 
}

type DhcpScopes struct {
	BlockUnknownIp bool `json:"blockUnknownIp"` 
	Log bool `json:"log"` 
	BlockScopeId OptionalString `json:"blockScopeId"` 
}

// ConnectionLimit - @deprecated use ConnLimitSettings instead
type ConnectionLimit struct {
	Enabled bool `json:"enabled"` 
	Soft int `json:"soft"` 
	Value int `json:"value"` // hard
	Rate OptionalLong `json:"rate"` 
	Exclusions OptionalIdReference `json:"exclusions"` 
	ExclSoft int `json:"exclSoft"` 
	ExclValue int `json:"exclValue"` // hard
	ExclRate OptionalLong `json:"exclRate"` 
}

type MiscSettingsConfig struct {
	AntiSpoofing AntiSpoofingConfig `json:"antiSpoofing"` 
	// [Deprecated]
	ConnectionLimit ConnectionLimit `json:"connectionLimit"` // @deprecated use ConnLimit instead
	Upnp UpnpConfig `json:"upnp"` 
	Ipv6 IPv6Config `json:"ipv6"` 
	DhcpScopes DhcpScopes `json:"dhcpScopes"` 
}

// MacFilterActionType - Mac Filter part
type MacFilterActionType string
const (
	MacFilterDeny MacFilterActionType = "MacFilterDeny" 
	MacFilterAllow MacFilterActionType = "MacFilterAllow" 
)

type MacAccessItem struct {
	MacAddress string `json:"macAddress"` 
	Description string `json:"description"` 
}

type MacAccessList []MacAccessItem

type MacFilterConfig struct {
	Enabled bool `json:"enabled"` 
	Interfaces IdReferenceList `json:"interfaces"` 
	MacFilterAction MacFilterActionType `json:"macFilterAction"` 
	MacAccessItems MacAccessList `json:"macAccessItems"` 
	AllowReserved bool `json:"allowReserved"` 
}

// ZeroConfigItemType - Zero-config network
type ZeroConfigItemType string
const (
	ZeroConfigVpnClients ZeroConfigItemType = "ZeroConfigVpnClients" 
	ZeroConfigVpnTunnel ZeroConfigItemType = "ZeroConfigVpnTunnel" 
	ZeroConfigInterface ZeroConfigItemType = "ZeroConfigInterface" 
)

// ZeroConfigItem - One VPN in rule
type ZeroConfigItem struct {
	Type ZeroConfigItemType `json:"type"` 
	Item IdReference `json:"item"` 
}

type ZeroConfigList []ZeroConfigItem

type ZeroConfigNetwork struct {
	Enabled bool `json:"enabled"` 
	Items ZeroConfigList `json:"items"` 
}

type SecuritySettingsConfig struct {
	MiscSettings MiscSettingsConfig `json:"miscSettings"` 
	MacFilter MacFilterConfig `json:"macFilter"` 
	ZeroConfigNetwork ZeroConfigNetwork `json:"zeroConfigNetwork"` 
}


// SecuritySettingsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
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


// SecuritySettingsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - structure containig security settings to be stored such as macfilter action, name, mac list and belonging interfaces.
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
