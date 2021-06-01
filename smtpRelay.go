package control

import "encoding/json"

type SmtpRelayConfig struct {
	UseKManager             bool              `json:"useKManager"`
	Server                  string            `json:"server"`
	RequireSecureConnection bool              `json:"requireSecureConnection"`
	AuthenticationRequired  bool              `json:"authenticationRequired"`
	Credentials             CredentialsConfig `json:"credentials"`
	FromAddress             OptionalString    `json:"fromAddress"`
}

// SmtpRelayGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SmtpRelayGet() (*SmtpRelayConfig, error) {
	data, err := s.CallRaw("SmtpRelay.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config SmtpRelayConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// SmtpRelaySet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SmtpRelaySet(config SmtpRelayConfig) error {
	params := struct {
		Config SmtpRelayConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("SmtpRelay.set", params)
	return err
}

// SmtpRelayTest - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - configuration structure of smtp relay module
//	address - email address where to send testing message. Recipient.
// Return
//	errors - list of errors \n
func (s *ServerConnection) SmtpRelayTest(config SmtpRelayConfig, address string) (ErrorList, error) {
	params := struct {
		Config  SmtpRelayConfig `json:"config"`
		Address string          `json:"address"`
	}{config, address}
	data, err := s.CallRaw("SmtpRelay.test", params)
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

// SmtpRelayGetStatus - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	errors - list of errors \n
func (s *ServerConnection) SmtpRelayGetStatus() (ErrorList, error) {
	data, err := s.CallRaw("SmtpRelay.getStatus", nil)
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
