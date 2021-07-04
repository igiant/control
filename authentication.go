package control

import "encoding/json"

type HttpProxyAuth struct {
	Enabled      bool           `json:"enabled"`
	AddressGroup OptionalEntity `json:"addressGroup"`
}

type Radius struct {
	Enabled     bool        `json:"enabled"`
	Password    Password    `json:"password"`
	Certificate IdReference `json:"certificate"`
}

type GuestConfig struct {
	Password OptionalString `json:"password"`
	Message  string         `json:"message"`
}

type AuthenticationConfig struct {
	AuthenticationRequired bool          `json:"authenticationRequired"`
	NtlmEnabled            bool          `json:"ntlmEnabled"`
	HttpProxyAuth          HttpProxyAuth `json:"httpProxyAuth"`
	Radius                 Radius        `json:"radius"`
	Guest                  GuestConfig   `json:"guest"`
	InactivityTimeout      OptionalLong  `json:"inactivityTimeout"` // in minutes
}

type TotpConfig struct {
	Required     bool `json:"required"`
	RemoteConfig bool `json:"remoteConfig"`
}

type JoinStatus string

const (
	JoinStatusConnected    JoinStatus = "JoinStatusConnected"
	JoinStatusDisconnected JoinStatus = "JoinStatusDisconnected"
	JoinStatusError        JoinStatus = "JoinStatusError"
)

// AuthenticationGet - Returns Authentication option settings
// Return
//	config - configuration values
func (s *ServerConnection) AuthenticationGet() (*AuthenticationConfig, error) {
	data, err := s.CallRaw("Authentication.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config AuthenticationConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// AuthenticationSet - Stores Authentication option settings
//	config - configuration values
func (s *ServerConnection) AuthenticationSet(config AuthenticationConfig) error {
	params := struct {
		Config AuthenticationConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("Authentication.set", params)
	return err
}

// AuthenticationJoin - joins computer to domain
//	hostName - name of the computer that will be set to computer and in domain controller
//	domainName - domain name (e.g. example.com)
//	credentials - domain account with rights to join the domain
//	server - server name (e.g. server.example.com) - filled only in case, that joinNeedServer returns true
// Return
//	message - text related to join result
//	status - current status
func (s *ServerConnection) AuthenticationJoin(hostName string, domainName string, credentials CredentialsConfig, server string) (*LocalizableMessage, *JoinStatus, error) {
	params := struct {
		HostName    string            `json:"hostName"`
		DomainName  string            `json:"domainName"`
		Credentials CredentialsConfig `json:"credentials"`
		Server      string            `json:"server"`
	}{hostName, domainName, credentials, server}
	data, err := s.CallRaw("Authentication.join", params)
	if err != nil {
		return nil, nil, err
	}
	message := struct {
		Result struct {
			Message LocalizableMessage `json:"message"`
			Status  JoinStatus         `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &message)
	return &message.Result.Message, &message.Result.Status, err
}

// AuthenticationIsJoinServerNeeded - tests, if join's param server need to be filled
//	domainName - domain name (e.g. example.com)
// Return
//	needServer - true - join must have param server.enabled on true and server.value filled
func (s *ServerConnection) AuthenticationIsJoinServerNeeded(domainName string) (bool, error) {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	data, err := s.CallRaw("Authentication.isJoinServerNeeded", params)
	if err != nil {
		return false, err
	}
	needServer := struct {
		Result struct {
			NeedServer bool `json:"needServer"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &needServer)
	return needServer.Result.NeedServer, err
}

// AuthenticationLeave - disconnects computer from domain
//	credentials - domain account with rights to leave from domain
// Return
//	message - text related to leave result
//	status - current status
func (s *ServerConnection) AuthenticationLeave(credentials CredentialsConfig) (*LocalizableMessage, *JoinStatus, error) {
	params := struct {
		Credentials CredentialsConfig `json:"credentials"`
	}{credentials}
	data, err := s.CallRaw("Authentication.leave", params)
	if err != nil {
		return nil, nil, err
	}
	message := struct {
		Result struct {
			Message LocalizableMessage `json:"message"`
			Status  JoinStatus         `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &message)
	return &message.Result.Message, &message.Result.Status, err
}

// AuthenticationGetJoinStatus - tests if computer is joined to domain
// Return
//	status - current status
//	domainName - a string representation of joined domain.
func (s *ServerConnection) AuthenticationGetJoinStatus() (*JoinStatus, string, error) {
	data, err := s.CallRaw("Authentication.getJoinStatus", nil)
	if err != nil {
		return nil, "", err
	}
	status := struct {
		Result struct {
			Status     JoinStatus `json:"status"`
			DomainName string     `json:"domainName"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, status.Result.DomainName, err
}

// AuthenticationGetTotpConfig - Returns TotpConfig
// Return
//	config - configuration values
func (s *ServerConnection) AuthenticationGetTotpConfig() (*TotpConfig, error) {
	data, err := s.CallRaw("Authentication.getTotpConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config TotpConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// AuthenticationSetTotpConfig - Stores TotpConfig
//	config - configuration values
func (s *ServerConnection) AuthenticationSetTotpConfig(config TotpConfig) error {
	params := struct {
		Config TotpConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("Authentication.setTotpConfig", params)
	return err
}
