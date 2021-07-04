package control

import "encoding/json"

type CentralManagementConfig struct {
	Enabled bool   `json:"enabled"`
	AppUrl  string `json:"appUrl"`
}

type CentralManagementStatus struct {
	Connected bool           `json:"connected"`
	Paired    bool           `json:"paired"`
	Url       OptionalString `json:"url"`
}

// CentralManagementGet - Returns configuration
// Return
//	config - Contains Structure with Central management settings.
func (s *ServerConnection) CentralManagementGet() (*CentralManagementConfig, error) {
	data, err := s.CallRaw("CentralManagement.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config CentralManagementConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// CentralManagementSet - Stores configuration
//	config - Contains Structure with Central management settings.
func (s *ServerConnection) CentralManagementSet(config CentralManagementConfig) error {
	params := struct {
		Config CentralManagementConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("CentralManagement.set", params)
	return err
}

// CentralManagementGetStatus - Returns actual state of Central management
// Return
//	status - actual state of Central management.
func (s *ServerConnection) CentralManagementGetStatus() (*CentralManagementStatus, error) {
	data, err := s.CallRaw("CentralManagement.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status CentralManagementStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// CentralManagementReset - Runs reset
func (s *ServerConnection) CentralManagementReset() error {
	_, err := s.CallRaw("CentralManagement.reset", nil)
	return err
}
