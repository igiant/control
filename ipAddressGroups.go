package control

import "encoding/json"

// IpAddressGroupsApply - Write changes cached in manager to configuration
// Return
//	errors - list of errors
func (s *ServerConnection) IpAddressGroupsApply() (ErrorList, error) {
	data, err := s.CallRaw("IpAddressGroups.apply", nil)
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

// IpAddressGroupsReset - Discard changes cached in manager
func (s *ServerConnection) IpAddressGroupsReset() error {
	_, err := s.CallRaw("IpAddressGroups.reset", nil)
	return err
}
