package control

import "encoding/json"

type DhcpExclusion struct {
	Description string    `json:"description"`
	IpStart     IpAddress `json:"ipStart"`
	IpEnd       IpAddress `json:"ipEnd"`
}

type DhcpExclusionList []DhcpExclusion

type DhcpOptionType string

const (
	DhcpBool         DhcpOptionType = "DhcpBool"
	DhcpInt8         DhcpOptionType = "DhcpInt8"
	DhcpInt16        DhcpOptionType = "DhcpInt16"
	DhcpInt32        DhcpOptionType = "DhcpInt32"
	DhcpIpAddr       DhcpOptionType = "DhcpIpAddr"
	DhcpString       DhcpOptionType = "DhcpString"
	DhcpHex          DhcpOptionType = "DhcpHex"
	DhcpTimeSigned   DhcpOptionType = "DhcpTimeSigned"
	DhcpTimeUnsigned DhcpOptionType = "DhcpTimeUnsigned"
	DhcpInt8List     DhcpOptionType = "DhcpInt8List"
	DhcpInt16List    DhcpOptionType = "DhcpInt16List"
	DhcpInt32List    DhcpOptionType = "DhcpInt32List"
	DhcpIpAddrList   DhcpOptionType = "DhcpIpAddrList"
	DhcpIpPairList   DhcpOptionType = "DhcpIpPairList"
	DhcpIpMaskList   DhcpOptionType = "DhcpIpMaskList"
	DhcpIpMaskIpList DhcpOptionType = "DhcpIpMaskIpList"
)

type IpListList []IpAddressList

type DhcpOption struct {
	Type     DhcpOptionType `json:"type"`
	OptionId int            `json:"optionId"`
	Name     string         `json:"name"`
	/**
	* @note: Value format:
	* DHOTIpAddr - address in dot notation (192.168.0.12)
	* DHOTHex - pairs of characters 0-9, a-f (12ef980ad8)
	* DHOTTimexxx - number of seconds (3600) (negative only for DHOTTimeSigned)
	* xxxList - values separated by ; (xxx;xx;x;xxxx)
	* DHOTBool - "0" / "1"
	 */
	/*@{ DHOTBool, DHOTInt8, DHOTInt16, DHOTInt32, DHOTIpAddr, DHOTString, DHOTHex
	, DHOTTimeSigned, DHOTTimeUnsigned, DHOTInt8List, DHOTInt16List, DHOTInt32List, DHOTIpAddrList, */
	Value string `json:"value"`
	/*@}*/
	/*@{ DHOTIpPairList, DHOTIpMaskList, DHOTIpMaskIpList */
	IpListList IpListList `json:"ipListList"`
	/*@}*/
}

type DhcpOptionList []DhcpOption

type DhcpScope struct {
	Id         KId               `json:"id"`
	Status     StoreStatus       `json:"status"`
	Enabled    bool              `json:"enabled"`
	Name       string            `json:"name"`
	IpStart    IpAddress         `json:"ipStart"`
	IpEnd      IpAddress         `json:"ipEnd"`
	IpMask     IpAddress         `json:"ipMask"`
	Exclusions DhcpExclusionList `json:"exclusions"`
	Options    DhcpOptionList    `json:"options"`
}

type DhcpScopeList []DhcpScope

type DhcpLeaseType string

const (
	DhcpTypeReservation DhcpLeaseType = "DhcpTypeReservation"
	DhcpTypeLease       DhcpLeaseType = "DhcpTypeLease"
)

type DhcpLease struct {
	Id      KId         `json:"id"`
	LeaseId KId         `json:"leaseId"` /// for internal purposes
	ScopeId KId         `json:"scopeId"`
	Status  StoreStatus `json:"status"`
	/** do not change in gui */
	Type             DhcpLeaseType  `json:"type"`
	Leased           bool           `json:"leased"`
	IsRas            bool           `json:"isRas"`
	CardManufacturer string         `json:"cardManufacturer"`
	IpAddress        IpAddress      `json:"ipAddress"`
	Name             string         `json:"name"`
	MacDefined       bool           `json:"macDefined"`
	MacAddress       string         `json:"macAddress"`
	HostName         string         `json:"hostName"`
	UserName         string         `json:"userName"`
	ExpirationDate   Date           `json:"expirationDate"` // @see SharedStructures.idl shared in lib
	ExpirationTime   Time           `json:"expirationTime"` // @see SharedStructures.idl shared in lib
	RequestDate      Date           `json:"requestDate"`    // @see SharedStructures.idl shared in lib
	RequestTime      Time           `json:"requestTime"`    // @see SharedStructures.idl shared in lib
	Options          DhcpOptionList `json:"options"`
}

type DhcpLeaseList []DhcpLease

type DhcpModeType string

const (
	DhcpAutomatic DhcpModeType = "DhcpAutomatic"
	DhcpManual    DhcpModeType = "DhcpManual"
)

type DhcpConfig struct {
	Enabled bool `json:"enabled"`
}

type DhcpMode struct {
	Type DhcpModeType `json:"type"`
}

// DhcpGet - get the list of scopes
// Parameters
//	query - conditions and limits
// Return
//	list - list of scopes and it's details
//	totalItems - count of all scopes on server (before the start/limit applied)
func (s *ServerConnection) DhcpGet(query SearchQuery) (DhcpScopeList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Dhcp.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       DhcpScopeList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DhcpCreate - add new scope
// Parameters
//	scopes - details for new scopes. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors
//	result - list of IDs assigned to each item
func (s *ServerConnection) DhcpCreate(scopes DhcpScopeList) (ErrorList, CreateResultList, error) {
	params := struct {
		Scopes DhcpScopeList `json:"scopes"`
	}{scopes}
	data, err := s.CallRaw("Dhcp.create", params)
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

// DhcpSet - update existing scope
// Parameters
//	scopeIds - ids of scopes to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other fields must be filled and they are written to all scopes specified by scopeIds.
// Return
//	errors - list of errors
func (s *ServerConnection) DhcpSet(scopeIds StringList, details DhcpScope) (ErrorList, error) {
	params := struct {
		ScopeIds StringList `json:"scopeIds"`
		Details  DhcpScope  `json:"details"`
	}{scopeIds, details}
	data, err := s.CallRaw("Dhcp.set", params)
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

// DhcpRemove - remove scope
// Parameters
//	scopeIds - ids of scopes that should be removed
// Return
//	errors - list of errors
func (s *ServerConnection) DhcpRemove(scopeIds StringList) (ErrorList, error) {
	params := struct {
		ScopeIds StringList `json:"scopeIds"`
	}{scopeIds}
	data, err := s.CallRaw("Dhcp.remove", params)
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

// DhcpGetInterfaceTemplate - get generated scope information for interface
// Parameters
//	ifaceId - id of interface, for which the template will be created
// Return
//	details - configuration for given ifaceId - can be passed to create method
func (s *ServerConnection) DhcpGetInterfaceTemplate(ifaceId KId) (*DhcpScope, error) {
	params := struct {
		IfaceId KId `json:"ifaceId"`
	}{ifaceId}
	data, err := s.CallRaw("Dhcp.getInterfaceTemplate", params)
	if err != nil {
		return nil, err
	}
	details := struct {
		Result struct {
			Details DhcpScope `json:"details"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &details)
	return &details.Result.Details, err
}

// DhcpGetLeases - get the list of leases
// Parameters
//	query - conditions and limits
//	scopeIds - list of scope Ids, where leases should be included (empty for all scopes)
// Return
//	list - list of leases/reservations and it's details
//	totalItems - count of all leases/reservations on server (before the start/limit applied)
func (s *ServerConnection) DhcpGetLeases(query SearchQuery, scopeIds KIdList) (DhcpLeaseList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query    SearchQuery `json:"query"`
		ScopeIds KIdList     `json:"scopeIds"`
	}{query, scopeIds}
	data, err := s.CallRaw("Dhcp.getLeases", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       DhcpLeaseList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DhcpCreateLeases - add new reservation
// Parameters
//	leases - details for new reservations. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors
//	result - list of IDs assigned to each item
func (s *ServerConnection) DhcpCreateLeases(leases DhcpLeaseList) (ErrorList, CreateResultList, error) {
	params := struct {
		Leases DhcpLeaseList `json:"leases"`
	}{leases}
	data, err := s.CallRaw("Dhcp.createLeases", params)
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

// DhcpSetLeases - update existing leases
// Parameters
//	leaseIds - ids of reservations to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other fields must be filled and they are written to all scopes specified by scopeIds.
// Return
//	errors - list of errors
func (s *ServerConnection) DhcpSetLeases(leaseIds StringList, details DhcpLease) (ErrorList, error) {
	params := struct {
		LeaseIds StringList `json:"leaseIds"`
		Details  DhcpLease  `json:"details"`
	}{leaseIds, details}
	data, err := s.CallRaw("Dhcp.setLeases", params)
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

// DhcpRemoveLeases - remove leases
// Parameters
//	leaseIds - ids of leases/reservations that should be removed
// Return
//	errors - list of errors
func (s *ServerConnection) DhcpRemoveLeases(leaseIds StringList) (ErrorList, error) {
	params := struct {
		LeaseIds StringList `json:"leaseIds"`
	}{leaseIds}
	data, err := s.CallRaw("Dhcp.removeLeases", params)
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

// DhcpGetMode - read Dhcp mode (not part of persistent manager, returns actual state)
// Return
//	mode - result
func (s *ServerConnection) DhcpGetMode() (*DhcpMode, error) {
	data, err := s.CallRaw("Dhcp.getMode", nil)
	if err != nil {
		return nil, err
	}
	mode := struct {
		Result struct {
			Mode DhcpMode `json:"mode"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &mode)
	return &mode.Result.Mode, err
}

// DhcpSetMode - stores Dhcp mode (not part of persistent manager, changes mode immediately)
// Parameters
//	mode - new value
func (s *ServerConnection) DhcpSetMode(mode DhcpMode) error {
	params := struct {
		Mode DhcpMode `json:"mode"`
	}{mode}
	_, err := s.CallRaw("Dhcp.setMode", params)
	return err
}

// DhcpGetConfig - read Dhcp configuration (not part of persistent manager, returns actual state)
// Return
//	config - configuration values
func (s *ServerConnection) DhcpGetConfig() (*DhcpConfig, error) {
	data, err := s.CallRaw("Dhcp.getConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config DhcpConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// DhcpSetConfig - stores Dhcp configuration (not part of persistent manager, changes settings immediately)
// Parameters
//	config - configuration values
func (s *ServerConnection) DhcpSetConfig(config DhcpConfig) error {
	params := struct {
		Config DhcpConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("Dhcp.setConfig", params)
	return err
}

// DhcpGetOptionList - get list of all options
// Return
//	options - list of all options
func (s *ServerConnection) DhcpGetOptionList() (DhcpOptionList, error) {
	data, err := s.CallRaw("Dhcp.getOptionList", nil)
	if err != nil {
		return nil, err
	}
	options := struct {
		Result struct {
			Options DhcpOptionList `json:"options"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &options)
	return options.Result.Options, err
}

// DhcpGetDeclinedLeases - Returns count of declined leases in all scopes, defined by scopeIDs param
// Parameters
//	scopeIds - list of scope IDs or empty for all scopes
// Return
//	count - count of declined leases
func (s *ServerConnection) DhcpGetDeclinedLeases(scopeIds KIdList) (int, error) {
	params := struct {
		ScopeIds KIdList `json:"scopeIds"`
	}{scopeIds}
	data, err := s.CallRaw("Dhcp.getDeclinedLeases", params)
	if err != nil {
		return 0, err
	}
	count := struct {
		Result struct {
			Count int `json:"count"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &count)
	return count.Result.Count, err
}

// DhcpRemoveDeclinedLeases - Removes declined leases defined by scopeIDs param from engine
// Parameters
//	scopeIds - list of scope IDs or empty for all scopes
func (s *ServerConnection) DhcpRemoveDeclinedLeases(scopeIds KIdList) error {
	params := struct {
		ScopeIds KIdList `json:"scopeIds"`
	}{scopeIds}
	_, err := s.CallRaw("Dhcp.removeDeclinedLeases", params)
	return err
}

// DhcpApply - Write changes cached in manager to configuration
// Return
//	errors - list of errors
func (s *ServerConnection) DhcpApply() (ErrorList, error) {
	data, err := s.CallRaw("Dhcp.apply", nil)
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

// DhcpReset - discard changes cached in manager
func (s *ServerConnection) DhcpReset() error {
	_, err := s.CallRaw("Dhcp.reset", nil)
	return err
}
