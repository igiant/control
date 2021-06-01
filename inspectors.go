package control

import "encoding/json"

// Inspector - Protocol Inspector - instances defined at compilation time, unchangeable
type Inspector struct {
	Name       string `json:"name"`       // unique name for the inspector
	IpProtocol int    `json:"ipProtocol"` // IP protocol (only 6 for TCP or 17 for UDP)
}

type InspectorList []Inspector

// InspectorsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	list - list of inspectors and it's details
func (s *ServerConnection) InspectorsGet() (InspectorList, error) {
	data, err := s.CallRaw("Inspectors.get", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List InspectorList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}
