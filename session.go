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

// SessionGetCsrfToken - Retrieves an unique session ID intended to be used for CSRF protection in web forms.
// This ID is different from the session cookie but also remains the same during the session lifetime.
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

// SessionGetUserName - Retrieves name os logged user
// Return
//  name - name os logged user
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

// Login - Log in given user.
// Please note that with a session to one module you cannot use another one (eg. with admin session you cannot use webmail).
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

// Logout - Destroys session
func (s *ServerConnection) Logout() error {
	_, err := s.CallRaw("Session.logout", nil)
	return err
}

// SessionGetSessionVariable - Returns clients defined variable stored in configuration for logged user
// Return
//  value - clients defined variable stored in configuration for logged user
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

// SessionSetSessionVariable - Stores clients defined variable to configuration for logged user
// Parameters
//	name - name of variable
//  value - value of variable
func (s *ServerConnection) SessionSetSessionVariable(name string, value string) error {
	params := struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{name, value}
	_, err := s.CallRaw("Session.setSessionVariable", params)
	return err
}

// SessionReset - Reset all persistent objects (managers) in session
func (s *ServerConnection) SessionReset() error {
	_, err := s.CallRaw("Session.reset", nil)
	return err
}

// SessionGetConfigTimestamp - Reloads configuration and returns timestamp of current configuration
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

// SessionConfirmConfig - Confirm the new configuration
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

// SessionGetConnectedInterface - Returns id of interface through which is client connected to server
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

// SessionGetLoginType - Returns type of login, that has to be performed
// Return
//  loginType - type of login
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
