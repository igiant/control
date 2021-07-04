package control

import "encoding/json"

type RouteType string

const (
	RouteSystem RouteType = "RouteSystem"
	RouteStatic RouteType = "RouteStatic"
	RouteVpn    RouteType = "RouteVpn"
)

type Route struct {
	Enabled bool      `json:"enabled"` // valid only for RouteStatic
	Name    string    `json:"name"`    // valid only for RouteStatic
	Type    RouteType `json:"type"`    // Type defines Descriptions: RouteSystem: 'System route'
	//               RouteVpn:  'VPN route'
	Network       IpAddress     `json:"network"`
	Mask          IpAddress     `json:"mask"`      // used for IPv4
	PrefixLen     int           `json:"prefixLen"` // used for IPv6
	Gateway       IpAddress     `json:"gateway"`
	InterfaceType InterfaceType `json:"interfaceType"` // @see InterfaceManager, used values: Ethernet, Ras, VpnTunnel, VpnServer
	InterfaceId   IdReference   `json:"interfaceId"`   // invalid - interface is no more in the configuration
	// @note: not used for interfaceType == VpnServer (it means 'VPN Server' in gui)
	Metric int `json:"metric"`
}

type RouteList []Route

// RoutingTableGet - Gets routing table from system.
// Return
//	routes - a list of routes currently stroed and used by system.
func (s *ServerConnection) RoutingTableGet() (RouteList, error) {
	data, err := s.CallRaw("RoutingTable.get", nil)
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

// RoutingTableGetStaticRoutes - Gets static routes.
// Return
//	routes - a list of routes currently stroed and used by Control.
func (s *ServerConnection) RoutingTableGetStaticRoutes() (RouteList, error) {
	data, err := s.CallRaw("RoutingTable.getStaticRoutes", nil)
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

// RoutingTableSetStaticRoutes - Set static routes
// Parameters
//	routes - A list of routes that should be stored in configuration.
// Return
//	errors - list of errors
func (s *ServerConnection) RoutingTableSetStaticRoutes(routes RouteList) (ErrorList, error) {
	params := struct {
		Routes RouteList `json:"routes"`
	}{routes}
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
