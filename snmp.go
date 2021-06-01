package control

type SnmpVersion string

const (
	SnmpV2c SnmpVersion = "SnmpV2c"
	SnmpV3  SnmpVersion = "SnmpV3"
)

type SnmpSettings struct {
	Enabled   bool        `json:"enabled"`
	Location  string      `json:"location"`
	Contact   string      `json:"contact"`
	Version   SnmpVersion `json:"version"`
	Community Password    `json:"community"`
	Username  string      `json:"username"`
	Password  Password    `json:"password"`
}

// SnmpGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	settings - A structure containing all the settings of SNMP that are stored.
func (s *ServerConnection) SnmpGet() (*SnmpSettings, error) {
	data, err := s.CallRaw("Snmp.get", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings SnmpSettings `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// SnmpSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	settings - A structure containing all the settings of SNMP that sould be stored.
func (s *ServerConnection) SnmpSet(settings SnmpSettings) error {
	params := struct {
		Settings SnmpSettings `json:"settings"`
	}{settings}
	_, err := s.CallRaw("Snmp.set", params)
	return err
}
