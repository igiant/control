package control

import "encoding/json"

// IpServiceReference - service reference used in various policies
type IpServiceReference struct {
	Id      KId    `json:"id"`
	Name    string `json:"name"`
	IsGroup bool   `json:"isGroup"`
	Invalid bool   `json:"invalid"`
}

// Different types of service's port definition

// PortComparator - names based on WAM's CompareOperator
type PortComparator string

const (
	Any         PortComparator = "Any"         // port not specified
	Equal       PortComparator = "Equal"       // '=' - equal to
	LessThan    PortComparator = "LessThan"    // '<' - lower than
	GreaterThan PortComparator = "GreaterThan" // '>' - greater than
	Range       PortComparator = "Range"       // from/to values
	List        PortComparator = "List"        // list of values
)

// PortList - List of ports (port must be < 65536)
type PortList []int

// PortCondition - Port - used to define condition for ports (i.e. one port, from/to range or greater/less than)
type PortCondition struct {
	Comparator PortComparator `json:"comparator"` // does the list contains either single port (equal/greater/etc.), range or list of ports
	Ports      PortList       `json:"ports"`
	/* - for comparator Equal, NotEqual, GreaterThan, LessThan => one port number for the condition
	- for comparator Range => two port numbers defining the from/to range (inclusive)
	- for comparator List => list of exact port numbers
	- for comparator Any => list should be empty (if not, just ignore the values)
	*/
}

const ipProtoIcmp int = 1

const ipProtoTcp int = 6

const ipProtoUdp int = 17

const ipProtoOther int = 128

const ipProtoTcpUdp int = 129

// IpService - basic structure for Service's properties
type IpService struct {
	Id          KId         `json:"id"` // never updated in store
	Status      StoreStatus `json:"status"`
	Name        string      `json:"name"`        // unique name for the service, max 23 chars
	Description string      `json:"description"` // brief description of the service, max 63 chars, can be empty
	Protocol    int         `json:"protocol"`    // ICMP 1, TCP - 6, UDP - 17, other - 128, TCP_UDP - 129
	Group       bool        `json:"group"`
	/*@{ TCP, UDP, TCP_UDP */
	SrcPort PortCondition `json:"srcPort"` // port(s) on client-side
	DstPort PortCondition `json:"dstPort"` // port(s) on server-side
	/*@}*/
	/*@{ TCP, UDP */
	Inspector string `json:"inspector"` // name of Protocol Inspector, @see InspectorManager
	/*@}*/
	/*@{ other */
	ProtoNumber int `json:"protoNumber"` // 1-255
	/*@}*/
	/*@{ ICMP */
	IcmpTypes StringList `json:"icmpTypes"` // "any" or list of numbers
	/*@}*/
	/*@{ group */
	Members IdReferenceList `json:"members"` // for group = true list of member IpService ids
	/*@}*/
}

type IpServiceList []IpService

// IpServicesGet - Get the list of services
//	query - conditions and limits. Included from weblib. Kerio Control engine implementation notes:
// Return
//	list - list of services and it's details
//	totalItems - count of all services on server (before the start/limit applied)
func (s *ServerConnection) IpServicesGet(query SearchQuery) (IpServiceList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("IpServices.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       IpServiceList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// IpServicesCreate - Add new services
//	services - details for new services. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors
//	result - list of IDs assigned to each item
func (s *ServerConnection) IpServicesCreate(services IpServiceList) (ErrorList, CreateResultList, error) {
	params := struct {
		Services IpServiceList `json:"services"`
	}{services}
	data, err := s.CallRaw("IpServices.create", params)
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

// IpServicesSet - Update existing services
//	serviceIds - ids of services to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other fields must be filled and they are written to all services specified by serviceIds.
// Return
//	errors - list of errors
func (s *ServerConnection) IpServicesSet(serviceIds StringList, details IpService) (ErrorList, error) {
	params := struct {
		ServiceIds StringList `json:"serviceIds"`
		Details    IpService  `json:"details"`
	}{serviceIds, details}
	data, err := s.CallRaw("IpServices.set", params)
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

// IpServicesRemove - Remove services
//	serviceIds - ids of services that should be removed
// Return
//	errors - list of errors TODO write particular values
func (s *ServerConnection) IpServicesRemove(serviceIds StringList) (ErrorList, error) {
	params := struct {
		ServiceIds StringList `json:"serviceIds"`
	}{serviceIds}
	data, err := s.CallRaw("IpServices.remove", params)
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

// IpServicesApply - Write changes cached in manager to configuration
// Return
//	errors - list of errors
func (s *ServerConnection) IpServicesApply() (ErrorList, error) {
	data, err := s.CallRaw("IpServices.apply", nil)
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

// IpServicesReset - Discard changes cached in manager
func (s *ServerConnection) IpServicesReset() error {
	_, err := s.CallRaw("IpServices.reset", nil)
	return err
}
