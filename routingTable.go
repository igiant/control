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


// RoutingTableSetStaticRoutes - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	routes - A list of routes that should be stored in configuration.
// Return
//	errors - list of errors \n
func (s *ServerConnection) RoutingTableSetStaticRoutes(routes RouteList, boolean [Opt]) (ErrorList, error) {
	params := struct {
		Routes RouteList `json:"routes"`
		Boolean [Opt] `json:"boolean"`
	}{routes, boolean}
	data, err := s.CallRaw("RoutingTable.setStaticRoutes", params)
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
