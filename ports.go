package control

import "encoding/json"

type PortAssignmentType string

const (
	PortAssignmentSwitch     PortAssignmentType = "PortAssignmentSwitch"
	PortAssignmentStandalone PortAssignmentType = "PortAssignmentStandalone"
	PortAssignmentUnassigned PortAssignmentType = "PortAssignmentUnassigned"
)

type PortType string

const (
	PortEthernet PortType = "PortEthernet"
	PortWifi     PortType = "PortWifi"
)

type SpeedDuplexType string

const (
	SpeedDuplexAuto     SpeedDuplexType = "SpeedDuplexAuto"
	SpeedDuplexHalf10   SpeedDuplexType = "SpeedDuplexHalf10"
	SpeedDuplexFull10   SpeedDuplexType = "SpeedDuplexFull10"
	SpeedDuplexHalf100  SpeedDuplexType = "SpeedDuplexHalf100"
	SpeedDuplexFull100  SpeedDuplexType = "SpeedDuplexFull100"
	SpeedDuplexFull1000 SpeedDuplexType = "SpeedDuplexFull1000"
)

type SpeedDuplexMayNotWorkList []SpeedDuplexType

type PortConfig struct {
	Id                    KId                       `json:"id"`
	Type                  PortType                  `json:"type"`
	Name                  string                    `json:"name"`
	Assignment            PortAssignmentType        `json:"assignment"`
	Vlans                 OptionalString            `json:"vlans"`
	SpeedDuplex           SpeedDuplexType           `json:"speedDuplex"`
	SpeedDuplexMayNotWork SpeedDuplexMayNotWorkList `json:"speedDuplexMayNotWork"`
}

type PortConfigList []PortConfig

// PortsGet - Returns list of system ports together with assignments
// note: System version has to support port and switch configuration
// Return
//	list - list of ports sorted by port's order of precedence (WiFi last)
func (s *ServerConnection) PortsGet() (PortConfigList, error) {
	data, err := s.CallRaw("Ports.get", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List PortConfigList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// PortsSet - Stores configuration for system ports
// If some of given items has id, that doesn't exist in current configuration, Error is returned
// Unknown items are not added, no item is deleted
// note: System version has to support port and switch configuration
// Parameters
//	ports - list of ports (order is not significant)
//	revertTimeout - If client doesn't confirm config to this timeout, configuration is reverted (0 - revert disabled)
// Return
//	errors - list of errors \n
func (s *ServerConnection) PortsSet(ports PortConfigList, revertTimeout int) (ErrorList, error) {
	params := struct {
		Ports         PortConfigList `json:"ports"`
		RevertTimeout int            `json:"revertTimeout"`
	}{ports, revertTimeout}
	data, err := s.CallRaw("Ports.set", params)
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
