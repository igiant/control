package control

import "encoding/json"

type UrlEntryType string

const (
	Url           UrlEntryType = "Url"
	UrlChildGroup UrlEntryType = "UrlChildGroup"
)

type UrlGroup struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

type UrlEntry struct {
	Id          KId          `json:"id"`
	GroupId     KId          `json:"groupId"`
	SharedId    KId          `json:"sharedId"` // read-only; filled when the item is shared in MyKerio
	GroupName   string       `json:"groupName"`
	Description string       `json:"description"`
	Type        UrlEntryType `json:"type"`
	Enabled     bool         `json:"enabled"`
	Status      StoreStatus  `json:"status"`
	/*@{ url */
	Url     string `json:"url"`
	IsRegex bool   `json:"isRegex"`
	/*@}*/
	/*@{ group */
	ChildGroupId   KId    `json:"childGroupId"`
	ChildGroupName string `json:"childGroupName"`
	/*@}*/
}

type UrlEntryList []UrlEntry

type UrlGroupList []UrlGroup

// UrlGroupsGet - Get the list of Url groups
// Parameters
//	query - conditions and limits. Included from weblib. Kerio Control engine implementation notes: \n
// Return
//	list - list of groups and it's details
//	totalItems - count of all groups on server (before the start/limit applied)
func (s *ServerConnection) UrlGroupsGet(query SearchQuery) (UrlEntryList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("UrlGroups.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       UrlEntryList `json:"list"`
			TotalItems int          `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// UrlGroupsCreate - Add new groups
// Parameters
//	groups - details for new groups. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) UrlGroupsCreate(groups UrlEntryList) (ErrorList, CreateResultList, error) {
	params := struct {
		Groups UrlEntryList `json:"groups"`
	}{groups}
	data, err := s.CallRaw("UrlGroups.create", params)
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

// UrlGroupsSet - Update existing groups
// Parameters
//	groupIds - ids of groups to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other fields must be filled and they are written to all groups specified by groupIds.
// Return
//	errors - list of errors \n
func (s *ServerConnection) UrlGroupsSet(groupIds StringList, details UrlEntry) (ErrorList, error) {
	params := struct {
		GroupIds StringList `json:"groupIds"`
		Details  UrlEntry   `json:"details"`
	}{groupIds, details}
	data, err := s.CallRaw("UrlGroups.set", params)
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

// UrlGroupsRemove - Remove groups
// Parameters
//	groupIds - ids of groups that should be removed
// Return
//	errors - list of errors TODO write particular errors
func (s *ServerConnection) UrlGroupsRemove(groupIds StringList) (ErrorList, error) {
	params := struct {
		GroupIds StringList `json:"groupIds"`
	}{groupIds}
	data, err := s.CallRaw("UrlGroups.remove", params)
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

// UrlGroupsApply - Write changes cached in manager to configuration
// Return
//	errors - list of errors \n
func (s *ServerConnection) UrlGroupsApply() (ErrorList, error) {
	data, err := s.CallRaw("UrlGroups.apply", nil)
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

// UrlGroupsReset - Discard changes cached in manager
func (s *ServerConnection) UrlGroupsReset() error {
	_, err := s.CallRaw("UrlGroups.reset", nil)
	return err
}

// UrlGroupsGetGroupList - Get the list of groups, sorted in asc order
func (s *ServerConnection) UrlGroupsGetGroupList() (UrlGroupList, error) {
	data, err := s.CallRaw("UrlGroups.getGroupList", nil)
	if err != nil {
		return nil, err
	}
	groups := struct {
		Result struct {
			Groups UrlGroupList `json:"groups"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &groups)
	return groups.Result.Groups, err
}
