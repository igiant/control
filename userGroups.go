package control

import "encoding/json"

type UserGroup struct {
	Id          KId               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Rights      UserRights        `json:"rights"`
	Members     UserReferenceList `json:"members"`
}

type UserGroupList []UserGroup

// UserGroupsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	query - conditions and limits
//	domainId - id of domain - only groups from this domain will be listed
// Return
//	list - list of groups and it's details
//	totalItems - count of all groups on server (before the start/limit applied)
func (s *ServerConnection) UserGroupsGet(query SearchQuery, domainId KId) (UserGroupList, int, error) {
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("UserGroups.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       UserGroupList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// UserGroupsCreate - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	groups - details for new groups. field id is assigned by the manager to temporary value until apply() or reset().
//	domainId - id of domain - specifies domain, where groups will be created (only local is supported)
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) UserGroupsCreate(groups UserGroupList, domainId KId) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups   UserGroupList `json:"groups"`
		DomainId KId           `json:"domainId"`
	}{groups, domainId}
	data, err := s.CallRaw("UserGroups.create", params)
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

// UserGroupsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	groupIds - ids of groups to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other values have to be present
//	domainId - id of domain - groups from this domain will be updated
// Return
//	errors - list of errors \n
func (s *ServerConnection) UserGroupsSet(groupIds StringList, details UserGroup, domainId KId) (ErrorList, error) {
	params := struct {
		GroupIds StringList `json:"groupIds"`
		Details  UserGroup  `json:"details"`
		DomainId KId        `json:"domainId"`
	}{groupIds, details, domainId}
	data, err := s.CallRaw("UserGroups.set", params)
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

// UserGroupsRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	groupIds - ids of groups that should be removed
//	domainId - id of domain - specifies domain, where groups will be removed (only local is supported)
// Return
//	errors - list of errors \n
func (s *ServerConnection) UserGroupsRemove(groupIds StringList, domainId KId) (ErrorList, error) {
	params := struct {
		GroupIds StringList `json:"groupIds"`
		DomainId KId        `json:"domainId"`
	}{groupIds, domainId}
	data, err := s.CallRaw("UserGroups.remove", params)
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
