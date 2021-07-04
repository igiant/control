package control

import "encoding/json"

// SharedDefinitionType - @brief enumerate types of definitions shared between MyKerio and your appliance
type SharedDefinitionType string

const (
	SharedDefinitionIpAddrGroup SharedDefinitionType = "SharedDefinitionIpAddrGroup"
	SharedDefinitionUrlGroup    SharedDefinitionType = "SharedDefinitionUrlGroup"
	SharedDefinitionTimeRange   SharedDefinitionType = "SharedDefinitionTimeRange"
)

// SharedDefinitionInfo - @brief bind shared definition type with version
type SharedDefinitionInfo struct {
	Type    SharedDefinitionType `json:"type"`
	Version int                  `json:"version"`
}

type SharedDefinitionInfoList []SharedDefinitionInfo

// @brief MyKerio-appliance interface to handle shared definitions

// SharedDefinitionsGetVersions - Gets appliance report to MyKerio what shared definitions it supports and what versions of shared definitions are stored in appliance
// Return
// list - list of appliance report
func (s *ServerConnection) SharedDefinitionsGetVersions() (SharedDefinitionInfoList, error) {
	data, err := s.CallRaw("SharedDefinitions.getVersions", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List SharedDefinitionInfoList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// SharedDefinitionsSetVersions - Sets appliance report to MyKerio what shared definitions it supports and what versions of shared definitions are stored in appliance
//	list - list of appliance report
func (s *ServerConnection) SharedDefinitionsSetVersions(list SharedDefinitionInfoList) error {
	params := struct {
		List SharedDefinitionInfoList `json:"list"`
	}{list}
	_, err := s.CallRaw("SharedDefinitions.setVersions", params)
	return err
}
