package control

// SourceNatMode - Mode of source address NAT
type SourceNatMode string

const (
	NatDefault   SourceNatMode = "NatDefault"
	NatInterface SourceNatMode = "NatInterface"
	NatIpAddress SourceNatMode = "NatIpAddress"
)

// NatBalancing - Balancing mode of source address NAT
type NatBalancing string

const (
	BalancingPerHost       NatBalancing = "BalancingPerHost"
	BalancingPerConnection NatBalancing = "BalancingPerConnection"
)

// TrafficIpVersion - IP version
type TrafficIpVersion string

const (
	Ipv4  TrafficIpVersion = "Ipv4"
	Ipv6  TrafficIpVersion = "Ipv6"
	IpAll TrafficIpVersion = "IpAll"
)

// InterfaceConditionType - Type of interface in rule
type InterfaceConditionType string

const (
	InterfaceInternet InterfaceConditionType = "InterfaceInternet"
	InterfaceTrusted  InterfaceConditionType = "InterfaceTrusted"
	InterfaceGuest    InterfaceConditionType = "InterfaceGuest"
	InterfaceSelected InterfaceConditionType = "InterfaceSelected"
)

// InterfaceCondition - One interface in rule
type InterfaceCondition struct {
	Type              InterfaceConditionType `json:"type"`
	InterfaceType     InterfaceType          `json:"interfaceType"`     // @see InterfaceManager, used values: Ethernet, Ras
	SelectedInterface IdReference            `json:"selectedInterface"` // invalid - interface is no more in the configuration (nothing)
	Enabled           bool                   `json:"enabled"`           // interface is present, but disabled/down
}

// VpnConditionType - Type of VPN in rule
type VpnConditionType string

const (
	IncomingClient VpnConditionType = "IncomingClient"
	SelectedTunnel VpnConditionType = "SelectedTunnel"
	AllTunnels     VpnConditionType = "AllTunnels"
)

// VpnCondition - One VPN in rule
type VpnCondition struct {
	Type    VpnConditionType `json:"type"`
	Tunnel  IdReference      `json:"tunnel"`  // invalid - tunnel is no more in the configuration (nothing)
	Enabled bool             `json:"enabled"` // tunnel is present, but disabled/down
}

// TrafficEntityType - Type of Traffic Entity in TrafficEntityList
type TrafficEntityType string

const (
	TrafficEntityHost         TrafficEntityType = "TrafficEntityHost"
	TrafficEntityNetwork      TrafficEntityType = "TrafficEntityNetwork"
	TrafficEntityRange        TrafficEntityType = "TrafficEntityRange"
	TrafficEntityAddressGroup TrafficEntityType = "TrafficEntityAddressGroup"
	TrafficEntityPrefix       TrafficEntityType = "TrafficEntityPrefix"
	TrafficEntityInterface    TrafficEntityType = "TrafficEntityInterface"
	TrafficEntityVpn          TrafficEntityType = "TrafficEntityVpn"
	TrafficEntityUsers        TrafficEntityType = "TrafficEntityUsers"
)

// TrafficEntity - One entity if there is list of entities in rule's Source or Destination
type TrafficEntity struct {
	Type TrafficEntityType `json:"type"`
	/*@{ host */
	Host string `json:"host"`
	/*@}*/
	/*@{ network, range */
	Addr1 IpAddress `json:"addr1"`
	Addr2 IpAddress `json:"addr2"`
	/*@}*/
	/*@{ IP address group */
	AddressGroup IdReference `json:"addressGroup"`
	/*@}*/
	/*@{ interface */
	InterfaceCondition InterfaceCondition `json:"interfaceCondition"`
	/*@}*/
	/*@{ vpn */
	VpnCondition VpnCondition `json:"vpnCondition"`
	/*@}*/
	/*@{ users */
	UserType UserConditionType `json:"userType"` // @see Users.idl, used values: AuthenticatedUsers, SelectedUsers
	User     UserReference     `json:"user"`     // @see UserManager
	/*@}*/
}

// TrafficEntityList - All entities in rule's Source or Destination
type TrafficEntityList []TrafficEntity

// TrafficCondition - Rule's Source or Destination
type TrafficCondition struct {
	Type     RuleConditionType `json:"type"`
	Firewall bool              `json:"firewall"`
	Entities TrafficEntityList `json:"entities"`
}

// TrafficServiceEntity - One service if there is list of services
type TrafficServiceEntity struct {
	DefinedService bool               `json:"definedService"`
	Service        IpServiceReference `json:"service"`
	Protocol       int                `json:"protocol"` // TCP - 6, UDP - 17 @see IpServiceManager
	Port           PortCondition      `json:"port"`
}

// TrafficServiceEntityList - List of services
type TrafficServiceEntityList []TrafficServiceEntity

// TrafficService - Rule's Services properties
type TrafficService struct {
	Type    RuleConditionType        `json:"type"`
	Entries TrafficServiceEntityList `json:"entries"`
}

// List of logEnabled values, order:
// 1. logPackets
// LogEnabled - 2. logConnections
type LogEnabled []bool

// TrafficRule - One traffic policy rule
type TrafficRule struct {
	Id KId `json:"id"`
	/*@{ name */
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	/*@{ rule Sourc, Destination */
	Source      TrafficCondition `json:"source"`
	Destination TrafficCondition `json:"destination"`
	/*@{ service */
	Service TrafficService `json:"service"`
	/*@{ IP verison */
	IpVersion TrafficIpVersion `json:"ipVersion"`
	/*@{ action */
	Action       RuleAction   `json:"action"` // @see common; possible values: allow, deny, drop
	LogEnabled   LogEnabled   `json:"logEnabled"`
	GraphEnabled bool         `json:"graphEnabled"`
	Dscp         OptionalLong `json:"dscp"`
	/*@{ NAT IP version */
	NatIpv4Only bool `json:"natIpv4Only"`
	/*@{ source NAT */
	EnableSourceNat        bool          `json:"enableSourceNat"`
	NatMode                SourceNatMode `json:"natMode"`
	AllowReverseConnection bool          `json:"allowReverseConnection"`
	/*@{ properties of NatDefault */
	Balancing NatBalancing `json:"balancing"`
	/*@{ properties of NatInterface */
	NatInterface  IdReference `json:"natInterface"`
	AllowFailover bool        `json:"allowFailover"`
	/*@{ properties of NatIpAddress */
	IpAddress   string `json:"ipAddress"`
	Ipv6Address string `json:"ipv6Address"`
	/*@{ destination NAT */
	EnableDestinationNat bool         `json:"enableDestinationNat"`
	TranslatedHost       string       `json:"translatedHost"`
	TranslatedIpv6Host   string       `json:"translatedIpv6Host"`
	TranslatedPort       OptionalLong `json:"translatedPort"`
	/*@{ valid time */
	ValidTimeRange IdReference `json:"validTimeRange"`
	/*@{ protocol inspector */
	Inspector string   `json:"inspector"` // name of Protocol Inspector, @see InspectorManager + values: default, none
	LastUsed  TimeSpan `json:"lastUsed"`  // last time when connection matched, read-only
}

// TrafficRuleList - All traffic policy rules
type TrafficRuleList []TrafficRule

type TrafficPolicyFilter struct {
	SourceIp      IpAddress `json:"sourceIp"`
	DestinationIp IpAddress `json:"destinationIp"`
	Port          int       `json:"port"`
}

// Manager for Traffic Policy

// TrafficPolicyGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	list - list of Traffic Policy rules
//	totalItems - count of all rules in Traffic Policy
func (s *ServerConnection) TrafficPolicyGet() (TrafficRuleList, int, error) {
	data, err := s.CallRaw("TrafficPolicy.get", nil)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       TrafficRuleList `json:"list"`
			TotalItems int             `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// TrafficPolicySet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	rules - list of Traffic Policy rules
//	defaultRule - properties of default rule
// Return
//	errors - list of errors \n
func (s *ServerConnection) TrafficPolicySet(rules TrafficRuleList, defaultRule TrafficRule) (ErrorList, error) {
	params := struct {
		Rules       TrafficRuleList `json:"rules"`
		DefaultRule TrafficRule     `json:"defaultRule"`
	}{rules, defaultRule}
	data, err := s.CallRaw("TrafficPolicy.set", params)
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

// TrafficPolicyGetCollisions - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) TrafficPolicyGetCollisions() (CollisionList, error) {
	data, err := s.CallRaw("TrafficPolicy.getCollisions", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List CollisionList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// TrafficPolicyGetDefaultRule - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	rule - properties of default rule
func (s *ServerConnection) TrafficPolicyGetDefaultRule() (*TrafficRule, error) {
	data, err := s.CallRaw("TrafficPolicy.getDefaultRule", nil)
	if err != nil {
		return nil, err
	}
	rule := struct {
		Result struct {
			Rule TrafficRule `json:"rule"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &rule)
	return &rule.Result.Rule, err
}

// TrafficPolicyFilterRules - 8001  Invalid parameters. - "Invalid parameters."
// Parameters
//	condition - Filter parameters. Empty parameter (0 for numbers) in condition means 'any'.
func (s *ServerConnection) TrafficPolicyFilterRules(condition TrafficPolicyFilter) (KIdList, error) {
	params := struct {
		Condition TrafficPolicyFilter `json:"condition"`
	}{condition}
	data, err := s.CallRaw("TrafficPolicy.filterRules", params)
	if err != nil {
		return nil, err
	}
	idList := struct {
		Result struct {
			IdList KIdList `json:"idList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &idList)
	return idList.Result.IdList, err
}

// TrafficPolicyNormalizeTrafficEntity - 8001  Invalid parameters. - "Invalid parameters."
// Parameters
//	input - TrafficEntity
func (s *ServerConnection) TrafficPolicyNormalizeTrafficEntity(input TrafficEntity) (ErrorList, *TrafficEntity, error) {
	params := struct {
		Input TrafficEntity `json:"input"`
	}{input}
	data, err := s.CallRaw("TrafficPolicy.normalizeTrafficEntity", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList     `json:"errors"`
			Result TrafficEntity `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, &errors.Result.Result, err
}
