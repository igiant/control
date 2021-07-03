package control

import "encoding/json"

type InterfaceType string

const (
	Ethernet  InterfaceType = "Ethernet"
	Ras       InterfaceType = "Ras"
	DialIn    InterfaceType = "DialIn"
	VpnServer InterfaceType = "VpnServer"
	VpnTunnel InterfaceType = "VpnTunnel"
)

type InterfaceModeType string

const (
	InterfaceModeManual    InterfaceModeType = "InterfaceModeManual"
	InterfaceModeAutomatic InterfaceModeType = "InterfaceModeAutomatic"
	InterfaceModeLinkLocal InterfaceModeType = "InterfaceModeLinkLocal"
)

type InterfaceEncapType string

const (
	InterfaceEncapNative InterfaceEncapType = "InterfaceEncapNative"
	InterfaceEncapPppoe  InterfaceEncapType = "InterfaceEncapPppoe"
)

type ConnectivityType string

const (
	Persistent    ConnectivityType = "Persistent"
	DialOnDemand  ConnectivityType = "DialOnDemand"
	Failover      ConnectivityType = "Failover"
	LoadBalancing ConnectivityType = "LoadBalancing"
)

type RasType string

const (
	PPPoE RasType = "PPPoE"
	PPTP  RasType = "PPTP"
	L2TP  RasType = "L2TP"
)

type ConnectivityConfig struct {
	Mode ConnectivityType `json:"mode"`
	/*@{ Failover, LoadBalancing */
	ProbeHosts OptionalString `json:"probeHosts"`
	/*@}*/
	/*@{ Failover */
	ReconnectTunnelsWhenPrimaryGoesBack bool `json:"reconnectTunnelsWhenPrimaryGoesBack"`
	LazyFailover                        bool `json:"lazyFailover"`
	/*@}*/
}

type MppeType string

const (
	MppeDisabled   MppeType = "MppeDisabled"
	MppeEnabled    MppeType = "MppeEnabled"
	Mppe128Enabled MppeType = "Mppe128Enabled"
)

type RasConfig struct {
	Dead              bool              `json:"dead"`
	EntryName         string            `json:"entryName"`
	UseOwnCredentials bool              `json:"useOwnCredentials"`
	Credentials       CredentialsConfig `json:"credentials"`
	Timeout           OptionalLong      `json:"timeout"`
	ConnectTime       OptionalEntity    `json:"connectTime"`
	NoConnectTime     OptionalEntity    `json:"noConnectTime"`
	BdScriptEnabled   bool              `json:"bdScriptEnabled"`
	AdScriptEnabled   bool              `json:"adScriptEnabled"`
	BhScriptEnabled   bool              `json:"bhScriptEnabled"`
	AhScriptEnabled   bool              `json:"ahScriptEnabled"`
	RasType           RasType           `json:"rasType"`
	PppoeIfaceId      string            `json:"pppoeIfaceId"`
	Server            string            `json:"server"`
	PapEnabled        bool              `json:"papEnabled"`
	ChapEnabled       bool              `json:"chapEnabled"`
	MschapEnabled     bool              `json:"mschapEnabled"`
	Mschapv2Enabled   bool              `json:"mschapv2Enabled"`
	Mppe              MppeType          `json:"mppe"`
	MppeStateful      bool              `json:"mppeStateful"`
}

type VpnRoute struct {
	Id          KId       `json:"id"`
	Enabled     bool      `json:"enabled"`
	Description string    `json:"description"`
	Network     IpAddress `json:"network"`
	Mask        IpAddress `json:"mask"`
}

type VpnRouteList []VpnRoute

type VpnServerConfig struct {
	/*@{ Kerio VPN */
	KerioVpnEnabled     bool        `json:"kerioVpnEnabled"`
	KerioVpnCertificate IdReference `json:"kerioVpnCertificate"`
	Port                int         `json:"port"`
	DefaultRoute        bool        `json:"defaultRoute"`
	/*@}*/
	/*@{ IPsec VPN */
	IpsecVpnEnabled     bool           `json:"ipsecVpnEnabled"`
	Mschapv2Enabled     bool           `json:"mschapv2Enabled"`
	IpsecVpnCertificate IdReference    `json:"ipsecVpnCertificate"`
	CipherIke           string         `json:"cipherIke"` // read-only
	CipherEsp           string         `json:"cipherEsp"` // read-only
	UseCertificate      bool           `json:"useCertificate"`
	Psk                 OptionalString `json:"psk"`
	/*@}*/
	Routes                 VpnRouteList `json:"routes"`
	Network                IpAddress    `json:"network"`
	Mask                   IpAddress    `json:"mask"`
	LocalDns               bool         `json:"localDns"`
	PrimaryDns             IpAddress    `json:"primaryDns"`
	SecondaryDns           IpAddress    `json:"secondaryDns"`
	AutodetectDomainSuffix bool         `json:"autodetectDomainSuffix"`
	DomainSuffix           string       `json:"domainSuffix"`
	LocalWins              bool         `json:"localWins"`
	PrimaryWins            IpAddress    `json:"primaryWins"`
	SecondaryWins          IpAddress    `json:"secondaryWins"`
}

type CertificateDn struct {
	Certificate IdReference `json:"certificate"`
	Value       string      `json:"value"`
}

type CertificateDnList []CertificateDn

type IpsecPeerIdConfig struct {
	DefaultLocalIdValue string            `json:"defaultLocalIdValue"`
	DefaultCipherIke    string            `json:"defaultCipherIke"`
	DefaultCipherEsp    string            `json:"defaultCipherEsp"`
	CertificateDnValues CertificateDnList `json:"certificateDnValues"` // values for IpsecPeerIdCertDn, based on choosen certificate
}

type VpnTunnelConfig struct {
	Type         VpnType        `json:"type"`
	Peer         OptionalString `json:"peer"`        // hostname or ip, passive if disabled
	LocalRoutes  VpnRouteList   `json:"localRoutes"` // IPsec only
	RemoteRoutes VpnRouteList   `json:"remoteRoutes"`
	/*@{ Kerio VPN */
	RemoteFingerprint        string `json:"remoteFingerprint"`
	UseRemoteAutomaticRoutes bool   `json:"useRemoteAutomaticRoutes"`
	UseRemoteCustomRoutes    bool   `json:"useRemoteCustomRoutes"`
	/*@}*/
	/*@{ IPsec VPN */
	Psk                     OptionalString `json:"psk"`         // use certificate if disabled
	Certificate             IdReference    `json:"certificate"` // empty ID for "Remote certificate"
	CipherIke               string         `json:"cipherIke"`   // read-only
	CipherEsp               string         `json:"cipherEsp"`   // read-only
	LocalIdValue            string         `json:"localIdValue"`
	RemoteIdValue           string         `json:"remoteIdValue"`
	UseLocalAutomaticRoutes bool           `json:"useLocalAutomaticRoutes"`
	UseLocalCustomRoutes    bool           `json:"useLocalCustomRoutes"`
	/*@}*/
	/*@}*/
}

type InterfaceGroupType string

const (
	Other    InterfaceGroupType = "Other"
	Guest    InterfaceGroupType = "Guest"
	Vpn      InterfaceGroupType = "Vpn"
	Trusted  InterfaceGroupType = "Trusted"
	Internet InterfaceGroupType = "Internet"
)

type InterfaceStatusType string

const (
	Up                InterfaceStatusType = "Up"
	Down              InterfaceStatusType = "Down"
	Connecting        InterfaceStatusType = "Connecting"
	Disconnecting     InterfaceStatusType = "Disconnecting"
	CableDisconnected InterfaceStatusType = "CableDisconnected"
	ErrorType         InterfaceStatusType = "Error"
	Backup            InterfaceStatusType = "Backup"
)

type FailoverRoleType string

const (
	None      FailoverRoleType = "None"
	Primary   FailoverRoleType = "Primary"
	Secondary FailoverRoleType = "Secondary"
)

type BandwidthUnit string

const (
	BandwidthUnitBits      BandwidthUnit = "BandwidthUnitBits"
	BandwidthUnitBytes     BandwidthUnit = "BandwidthUnitBytes"
	BandwidthUnitKilobits  BandwidthUnit = "BandwidthUnitKilobits"
	BandwidthUnitKiloBytes BandwidthUnit = "BandwidthUnitKiloBytes"
	BandwidthUnitMegabits  BandwidthUnit = "BandwidthUnitMegabits"
	BandwidthUnitMegaBytes BandwidthUnit = "BandwidthUnitMegaBytes"
	BandwidthUnitPercent   BandwidthUnit = "BandwidthUnitPercent"
)

// InterfaceConnectivityParameters - mode-dependent data
type InterfaceConnectivityParameters struct {
	/*@{ Failover */
	FailoverRole FailoverRoleType `json:"failoverRole"`
	/*@}*/
	/*@{ OnDemand */
	OnDemand bool `json:"onDemand"`
	/*@}*/
	/*@{ Balancing */
	LoadBalancingWeight OptionalLong `json:"loadBalancingWeight"`
	/*@}*/
}

type InterfaceFlags struct {
	Deletable     bool `json:"deletable"`
	Dialable      bool `json:"dialable"`
	Hangable      bool `json:"hangable"`
	VirtualSwitch bool `json:"virtualSwitch"`
	Wifi          bool `json:"wifi"`
	Vlan          bool `json:"vlan"`
}

type DetailsConfig struct {
	Localizable        bool               `json:"localizable"`
	FixedMessage       string             `json:"fixedMessage"`
	LocalizableMessage LocalizableMessage `json:"localizableMessage"`
}

type IpAddressMask struct {
	Ip         IpAddress `json:"ip"` // can't name it ipAddress :-(
	SubnetMask IpAddress `json:"subnetMask"`
}

type IpAddressMaskList []IpAddressMask

type Ip6AddressMask struct {
	Ip           Ip6Address `json:"ip"`
	PrefixLength int        `json:"prefixLength"`
}

type Ip6AddressMaskList []Ip6AddressMask

type Interface struct {
	Enabled           bool          `json:"enabled"`
	Type              InterfaceType `json:"type"`
	Status            StoreStatus   `json:"status"`
	DhcpServerEnabled bool          `json:"dhcpServerEnabled"`
	/*@{ grid columns. they are not common subset from interface types */
	Id         KId                 `json:"id"`
	Group      InterfaceGroupType  `json:"group"`
	Name       string              `json:"name"`
	LinkStatus InterfaceStatusType `json:"linkStatus"`
	Details    DetailsConfig       `json:"details"`
	Mac        string              `json:"mac"`
	SystemName string              `json:"systemName"`
	/*@{ IPv4 */
	Ip4Enabled          bool              `json:"ip4Enabled"`
	Mode                InterfaceModeType `json:"mode"`
	Ip                  IpAddress         `json:"ip"`
	SubnetMask          IpAddress         `json:"subnetMask"`
	SecondaryAddresses  IpAddressMaskList `json:"secondaryAddresses"`
	DnsAutodetected     bool              `json:"dnsAutodetected"`
	DnsServers          string            `json:"dnsServers"`
	GatewayAutodetected bool              `json:"gatewayAutodetected"`
	Gateway             IpAddress         `json:"gateway"`
	/*@}*/
	/*@{ IPv6 */
	Ip6Enabled                  bool               `json:"ip6Enabled"`
	Ip6Mode                     InterfaceModeType  `json:"ip6Mode"`
	Ip6Addresses                Ip6AddressMaskList `json:"ip6Addresses"`
	LinkIp6Address              Ip6Address         `json:"linkIp6Address"`
	Ip6Gateway                  IpAddress          `json:"ip6Gateway"`
	RoutedIp6PrefixAutodetected bool               `json:"routedIp6PrefixAutodetected"`
	RoutedIp6Prefix             string             `json:"routedIp6Prefix"`
	/*@}*/
	/*@}*/
	ConnectivityParameters InterfaceConnectivityParameters `json:"connectivityParameters"`
	/*@{ engine on linux */
	Encap       InterfaceEncapType `json:"encap"`
	MtuOverride OptionalLong       `json:"mtuOverride"`
	MacOverride OptionalString     `json:"macOverride"`
	/*@}*/
	/* medium-dependent data */
	/*@{ RAS */
	Ras RasConfig `json:"ras"`
	/*@}*/
	/*@{ VPN Server */
	Server VpnServerConfig `json:"server"`
	/*@}*/
	/*@{ VPN Tunnel */
	Tunnel VpnTunnelConfig `json:"tunnel"`
	/*@}*/
	Flags InterfaceFlags `json:"flags"`
	/*@{ engine on HW Box */
	Ports KIdList `json:"ports"`
	Stp   bool    `json:"stp"`
	/*@}*/
	/*@{ for flags.vlan */
	VlanId int `json:"vlanId"`
	/*@}*/
}

type InterfaceList []Interface

type ConnectivityStatus string

const (
	ConnectivityOk       ConnectivityStatus = "ConnectivityOk"
	ConnectivityChecking ConnectivityStatus = "ConnectivityChecking"
	ConnectivityError    ConnectivityStatus = "ConnectivityError"
)

type IpCollisionList []KIdList

// InterfacesGet - Obtain list of interfaces.
// When sortByGroup is true and sorting is 'name', sorting order is 'group', 'type', 'name'
func (s *ServerConnection) InterfacesGet(query SearchQuery, sortByGroup bool) (InterfaceList, int, error) {
	params := struct {
		Query       SearchQuery `json:"query"`
		SortByGroup bool        `json:"sortByGroup"`
	}{query, sortByGroup}
	data, err := s.CallRaw("Interfaces.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       InterfaceList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// InterfacesCreate - Creates new interface (Only one interface can be created at a time) - VPN Tunnel or RAS on Ape/Box
// Parameters
//	list - list of interfaces desired to be created (must contain exactly one item)
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) InterfacesCreate(list InterfaceList) (ErrorList, CreateResultList, error) {
	params := struct {
		List InterfaceList `json:"list"`
	}{list}
	data, err := s.CallRaw("Interfaces.create", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// InterfacesSet - Update interface's details.
// Parameters
//	ids - list of IDs of interfaces to modify
//	details - details to set to every interface lister in ids parameter
// Return
//	errors - list of errors
func (s *ServerConnection) InterfacesSet(ids KIdList, details Interface) (ErrorList, error) {
	params := struct {
		Ids     KIdList   `json:"ids"`
		Details Interface `json:"details"`
	}{ids, details}
	data, err := s.CallRaw("Interfaces.set", params)
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

// InterfacesRemove - Delete Interface configuration - VPN Tunnel or RAS on Ape/Box
// Parameters
//	ids - list of IDs of interfaces to modify
// Return
//	errors - list of errors
func (s *ServerConnection) InterfacesRemove(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Interfaces.remove", params)
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

// InterfacesCheckIpCollision - Checks collision of all interfaces IP + VPN Server network
// Return
//  collisions - list of ip collision
func (s *ServerConnection) InterfacesCheckIpCollision() (IpCollisionList, error) {
	data, err := s.CallRaw("Interfaces.checkIpCollision", nil)
	if err != nil {
		return nil, err
	}
	collisions := struct {
		Result struct {
			Collisions IpCollisionList `json:"collisions"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &collisions)
	return collisions.Result.Collisions, err
}

// InterfacesGetWarnings - Checks Link Load Balancing warnings
// Return
//  warnings - list of notification type
func (s *ServerConnection) InterfacesGetWarnings() (NotificationTypeList, error) {
	data, err := s.CallRaw("Interfaces.getWarnings", nil)
	if err != nil {
		return nil, err
	}
	warnings := struct {
		Result struct {
			Warnings NotificationTypeList `json:"warnings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &warnings)
	return warnings.Result.Warnings, err
}

// InterfacesGetConnectivityConfig - Returns Connectivity config values
// Return
//	config - Connectivity config values
func (s *ServerConnection) InterfacesGetConnectivityConfig() (*ConnectivityConfig, error) {
	data, err := s.CallRaw("Interfaces.getConnectivityConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ConnectivityConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// InterfacesSetConnectivityConfig - Stores Connectivity config values
// Parameters
//	config - Connectivity config values
func (s *ServerConnection) InterfacesSetConnectivityConfig(config ConnectivityConfig) error {
	params := struct {
		Config ConnectivityConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("Interfaces.setConnectivityConfig", params)
	return err
}

// InterfacesStartConnectivityTest - Initiates testing of connectivity
func (s *ServerConnection) InterfacesStartConnectivityTest() error {
	_, err := s.CallRaw("Interfaces.startConnectivityTest", nil)
	return err
}

// InterfacesConnectivityTestStatus - Returns progress of connectivity test
// Return
//	status - actual status
func (s *ServerConnection) InterfacesConnectivityTestStatus() (*ConnectivityStatus, error) {
	data, err := s.CallRaw("Interfaces.connectivityTestStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status ConnectivityStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// InterfacesCancelConnectivityTest - Cancels testing of connectivity nad sets status to ConnectivityError
func (s *ServerConnection) InterfacesCancelConnectivityTest() error {
	_, err := s.CallRaw("Interfaces.cancelConnectivityTest", nil)
	return err
}

// InterfacesDial - Dial interface. Works only for disconnected RAS. Action is taken immediatelly, without apply.
func (s *ServerConnection) InterfacesDial(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("Interfaces.dial", params)
	return err
}

// InterfacesHangup - Hangup interface. Works only for connected RAS. Action is taken immediatelly, without apply.
// Parameters
//	id - interface id
func (s *ServerConnection) InterfacesHangup(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("Interfaces.hangup", params)
	return err
}

// InterfacesGetIpsecPeerIdConfig - Returns (defaults/read-only) values to be displayed on VPN Tunnel IPsec dialog as peer ID config
// Return
//	config - values to be displayed on VPN Tunnel IPsec dialog as peer ID config
func (s *ServerConnection) InterfacesGetIpsecPeerIdConfig() (*IpsecPeerIdConfig, error) {
	data, err := s.CallRaw("Interfaces.getIpsecPeerIdConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config IpsecPeerIdConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// InterfacesApply - write changes cached in manager to configuration
// Parameters
//	revertTimeout how many seconds to wait for confirmation until revert is performed
// Return
//	errors - list of errors
func (s *ServerConnection) InterfacesApply(revertTimeout int) (ErrorList, error) {
	params := struct {
		RevertTimeout int `json:"revertTimeout"`
	}{revertTimeout}
	data, err := s.CallRaw("Interfaces.apply", params)
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

// InterfacesReset - discard changes cached in manager
func (s *ServerConnection) InterfacesReset() error {
	_, err := s.CallRaw("Interfaces.reset", nil)
	return err
}
