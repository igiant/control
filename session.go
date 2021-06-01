package control

import "encoding/json"

type LoginType string

const (
	LoginRegular      LoginType = "LoginRegular"
	LoginAutomatic    LoginType = "LoginAutomatic"
	LoginReactivation LoginType = "LoginReactivation"
)

type ClientTimestamp struct {
	Name      string `json:"name"`
	Timestamp int    `json:"timestamp"`
}

type ClientTimestampList []ClientTimestamp

// May be created only if user is authenticated (request contains valid cookie)

// SessionGetCsrfToken - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SessionGetCsrfToken() (string, error) {
	data, err := s.CallRaw("Session.getCsrfToken", nil)
	if err != nil {
		return "", err
	}
	token := struct {
		Result struct {
			Token string `json:"token"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &token)
	return token.Result.Token, err
}

// SessionGetUserName - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SessionGetUserName() (string, error) {
	data, err := s.CallRaw("Session.getUserName", nil)
	if err != nil {
		return "", err
	}
	name := struct {
		Result struct {
			Name string `json:"name"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &name)
	return name.Result.Name, err
}

// Login - [KLoginMethod]
// Parameters
//	userName - login name + domain name (can be omitted if primary/local) of the user to be logged in, e.g. "jdoe" or "jdoe@company.com"
//	password - password of the user to be logged in
//	application - client application description
func (s *ServerConnection) Login(userName string, password string, application *ApiApplication) error {
	if application == nil {
		application = NewApplication("", "", "")
	}

	params := struct {
		UserName    string         `json:"userName"`
		Password    string         `json:"password"`
		Application ApiApplication `json:"application"`
	}{userName, password, *application}
	data, err := s.CallRaw("Session.login", params)
	if err != nil {
		return err
	}
	token := struct {
		Result struct {
			Token string `json:"token"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return err
	}
	s.Token = &token.Result.Token
	return nil
}

// Logout - [KLogoutMethod]
func (s *ServerConnection) Logout() error {
	_, err := s.CallRaw("Session.logout", nil)
	return err
}

// SessionGetSessionVariable - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SessionGetSessionVariable(name string) (string, error) {
	params := struct {
		Name string `json:"name"`
	}{name}
	data, err := s.CallRaw("Session.getSessionVariable", params)
	if err != nil {
		return "", err
	}
	value := struct {
		Result struct {
			Value string `json:"value"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &value)
	return value.Result.Value, err
}

// SessionSetSessionVariable - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SessionSetSessionVariable(name string, value string) error {
	params := struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{name, value}
	_, err := s.CallRaw("Session.setSessionVariable", params)
	return err
}

// SessionReset - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SessionReset() error {
	_, err := s.CallRaw("Session.reset", nil)
	return err
}

// SessionGetConfigTimestamp - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	clientTimestampList - is empty in case, that cut-off prevention is not active
func (s *ServerConnection) SessionGetConfigTimestamp() (ClientTimestampList, error) {
	data, err := s.CallRaw("Session.getConfigTimestamp", nil)
	if err != nil {
		return nil, err
	}
	clientTimestampList := struct {
		Result struct {
			ClientTimestampList ClientTimestampList `json:"clientTimestampList"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &clientTimestampList)
	return clientTimestampList.Result.ClientTimestampList, err
}

// SessionConfirmConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	clientTimestampList - values obtained by getConfigTimestamp
// Return
//	confirmed - true in case, that cut-off prevention was active and timestamp matched last provided timestamp
func (s *ServerConnection) SessionConfirmConfig(clientTimestampList ClientTimestampList) (bool, error) {
	params := struct {
		ClientTimestampList ClientTimestampList `json:"clientTimestampList"`
	}{clientTimestampList}
	data, err := s.CallRaw("Session.confirmConfig", params)
	if err != nil {
		return false, err
	}
	confirmed := struct {
		Result struct {
			Confirmed bool `json:"confirmed"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &confirmed)
	return confirmed.Result.Confirmed, err
}

// SessionGetConnectedInterface - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	id - id of interface or empty in case of localhost
func (s *ServerConnection) SessionGetConnectedInterface() (*KId, error) {
	data, err := s.CallRaw("Session.getConnectedInterface", nil)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// SessionGetLoginType - [KAuthentication(AuthenticationMode.NO_AUTHENTICATION_REQUIRED)]
func (s *ServerConnection) SessionGetLoginType() (*LoginType, error) {
	data, err := s.CallRaw("Session.getLoginType", nil)
	if err != nil {
		return nil, err
	}
	loginType := struct {
		Result struct {
			LoginType LoginType `json:"loginType"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &loginType)
	return &loginType.Result.LoginType, err
}
