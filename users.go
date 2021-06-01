package control

type UserConditionType string

const (
	AnyUser            UserConditionType = "AnyUser"
	AuthenticatedUsers UserConditionType = "AuthenticatedUsers"
	UnrecognizedUsers  UserConditionType = "UnrecognizedUsers"
	SelectedUsers      UserConditionType = "SelectedUsers"
	Nobody             UserConditionType = "Nobody"
)

// Roles in webadmin correspond with user rights. If user doesn't have even
// UserRoleType - read right, he can't create session, so there is no "none" in the enum
type UserRoleType string

const (
	Auditor   UserRoleType = "Auditor"
	FullAdmin UserRoleType = "FullAdmin"
)

type UserRoleList []UserRoleType

type AuthType string

const (
	Internal   AuthType = "Internal"
	KerberosNt AuthType = "KerberosNt"
)

type UserRights struct {
	ReadConfig        bool `json:"readConfig"`
	WriteConfig       bool `json:"writeConfig"`
	OverrideWwwFilter bool `json:"overrideWwwFilter"`
	UnlockRule        bool `json:"unlockRule"`
	DialRasConnection bool `json:"dialRasConnection"`
	ConnectVpn        bool `json:"connectVpn"`
	ConnectSslVpn     bool `json:"connectSslVpn"`
	UseP2p            bool `json:"useP2p"`
}

type QuotaType string

const (
	QuotaBoth     QuotaType = "QuotaBoth"
	QuotaDownload QuotaType = "QuotaDownload"
	QuotaUpload   QuotaType = "QuotaUpload"
)

type QuotaInterval struct {
	Enabled bool               `json:"enabled"`
	Type    QuotaType          `json:"type"`
	Limit   ByteValueWithUnits `json:"limit"`
}

type Quota struct {
	Daily        QuotaInterval `json:"daily"`
	Weekly       QuotaInterval `json:"weekly"`
	Monthly      QuotaInterval `json:"monthly"`
	BlockTraffic bool          `json:"blockTraffic"`
	NotifyUser   bool          `json:"notifyUser"`
}

type WwwFilter struct {
	JavaApplet  bool `json:"javaApplet"`
	EmbedObject bool `json:"embedObject"`
	Script      bool `json:"script"`
	Popup       bool `json:"popup"`
	Referer     bool `json:"referer"`
}

// UserData - common user data, used in domain template
type UserData struct {
	Rights    UserRights `json:"rights"`
	Quota     Quota      `json:"quota"`
	WwwFilter WwwFilter  `json:"wwwFilter"`
	Language  string     `json:"language"`
}

type AutoLogin struct {
	MacAddresses OptionalStringList    `json:"macAddresses"`
	Addresses    OptionalIpAddressList `json:"addresses"`
	AddressGroup OptionalEntity        `json:"addressGroup"`
}

// UserReference - user or group reference, used as "member of group" in user, "have member" in group and in various policies
type UserReference struct {
	Id         KId    `json:"id"`
	Name       string `json:"name"`
	IsGroup    bool   `json:"isGroup"`
	DomainName string `json:"domainName"`
}

type UserReferenceList []UserReference

type AddresseeType string

const (
	AddresseeEmail AddresseeType = "AddresseeEmail"
	AddresseeUser  AddresseeType = "AddresseeUser"
)

type Addressee struct {
	Type  AddresseeType `json:"type"`
	Email string        `json:"email"`
	/*@{ valid for type AddresseeUser */
	User UserReference `json:"user"`
	/*@}*/
}

type UserCondition struct {
	Type  UserConditionType `json:"type"`
	Users UserReferenceList `json:"users"`
}

type UserSettings struct {
	CalculatedLanguage string        `json:"calculatedLanguage"`
	Language           string        `json:"language"`
	DetectedLanguage   string        `json:"detectedLanguage"`
	Roles              UserRoleList  `json:"roles"`
	User               UserReference `json:"user"`
	FullName           string        `json:"fullName"`
	Email              string        `json:"email"`
}

type User struct {
	Id                KId               `json:"id"`
	Credentials       CredentialsConfig `json:"credentials"`
	FullName          string            `json:"fullName"`
	Description       string            `json:"description"`
	Email             string            `json:"email"`
	AuthType          AuthType          `json:"authType"`
	LocalEnabled      bool              `json:"localEnabled"`
	AdEnabled         bool              `json:"adEnabled"`
	UseTemplate       bool              `json:"useTemplate"`
	Data              UserData          `json:"data"`
	AutoLogin         AutoLogin         `json:"autoLogin"`
	VpnAddress        OptionalIpAddress `json:"vpnAddress"`
	Groups            UserReferenceList `json:"groups"`
	ConflictWithLocal bool              `json:"conflictWithLocal"`
	TotpConfigured    bool              `json:"totpConfigured"`
}

type UserList []User

// UsersGet - 1004 Access denied.    - "Insufficient rights to perform the requested operation."
// Parameters
//	query - conditions and limits
//	domainId - id of domain - only users from this domain will be listed
// Return
//	warnings - list of warnings \n
//	list - list of users and it's details
//	totalItems - count of all users on server (before the start/limit applied)
func (s *ServerConnection) UsersGet(query SearchQuery, domainId KId) (ErrorList, UserList, int, error) {
	params := struct {
		Query    SearchQuery `json:"query"`
		DomainId KId         `json:"domainId"`
	}{query, domainId}
	data, err := s.CallRaw("Users.get", params)
	if err != nil {
		return nil, nil, 0, err
	}
	warnings := struct {
		Result struct {
			Warnings   ErrorList `json:"warnings"`
			List       UserList  `json:"list"`
			TotalItems int       `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &warnings)
	return warnings.Result.Warnings, warnings.Result.List, warnings.Result.TotalItems, err
}

// UsersCreate - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	users - details for new users. field id is assigned by the manager to temporary value until apply() or reset().
//	domainId - id of domain - specifies domain, where user will be created (only local is supported)
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) UsersCreate(users UserList, domainId KId) (ErrorList, CreateResultList, error) {
	params := struct {
		Users    UserList `json:"users"`
		DomainId KId      `json:"domainId"`
	}{users, domainId}
	data, err := s.CallRaw("Users.create", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// UsersSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	userIds - ids of users to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. Only filled details will be stored in users config defined by userIds
//	domainId - id of domain - users from this domain will be updated
// Return
//	errors - list of errors \n
func (s *ServerConnection) UsersSet(userIds KIdList, details User, domainId KId) (ErrorList, error) {
	params := struct {
		UserIds  KIdList `json:"userIds"`
		Details  User    `json:"details"`
		DomainId KId     `json:"domainId"`
	}{userIds, details, domainId}
	data, err := s.CallRaw("Users.set", params)
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

// UsersRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	userIds - ids of users that should be removed
//	domainId - id of domain - specifies domain, where user will be removed (only local is supported)
// Return
//	errors - list of errors \n
func (s *ServerConnection) UsersRemove(userIds KIdList, domainId KId) (ErrorList, error) {
	params := struct {
		UserIds  KIdList `json:"userIds"`
		DomainId KId     `json:"domainId"`
	}{userIds, domainId}
	data, err := s.CallRaw("Users.remove", params)
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

// UsersConvertLocalUsers - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domainId - id of domain - specifies domain, from which users will be loaded
func (s *ServerConnection) UsersConvertLocalUsers(domainId KId) error {
	params := struct {
		DomainId KId `json:"domainId"`
	}{domainId}
	_, err := s.CallRaw("Users.convertLocalUsers", params)
	return err
}

// UsersGetAdUsers - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domainName - name of AD domain
//	server - AD server
//	credentials - username and password for user with read privilegies
//	ldapSecure - use secured connection
// Return
//	users - list of users and details
func (s *ServerConnection) UsersGetAdUsers(domainName string, server string, credentials CredentialsConfig, ldapSecure bool) (UserList, error) {
	params := struct {
		DomainName  string            `json:"domainName"`
		Server      string            `json:"server"`
		Credentials CredentialsConfig `json:"credentials"`
		LdapSecure  bool              `json:"ldapSecure"`
	}{domainName, server, credentials, ldapSecure}
	data, err := s.CallRaw("Users.getAdUsers", params)
	if err != nil {
		return nil, err
	}
	users := struct {
		Result struct {
			Users UserList `json:"users"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &users)
	return users.Result.Users, err
}

// UsersGetNtUsers - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domainName - name of NT domain
// Return
//	users - list of users and details
func (s *ServerConnection) UsersGetNtUsers(domainName string) (UserList, error) {
	params := struct {
		DomainName string `json:"domainName"`
	}{domainName}
	data, err := s.CallRaw("Users.getNtUsers", params)
	if err != nil {
		return nil, err
	}
	users := struct {
		Result struct {
			Users UserList `json:"users"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &users)
	return users.Result.Users, err
}

// UsersGetSupportedLanguages - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	languages - list of languages
func (s *ServerConnection) UsersGetSupportedLanguages() (NamedValueList, error) {
	data, err := s.CallRaw("Users.getSupportedLanguages", nil)
	if err != nil {
		return nil, err
	}
	languages := struct {
		Result struct {
			Languages NamedValueList `json:"languages"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &languages)
	return languages.Result.Languages, err
}

// UsersGetMySettings - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	settings - list of all settings
func (s *ServerConnection) UsersGetMySettings() (*UserSettings, error) {
	data, err := s.CallRaw("Users.getMySettings", nil)
	if err != nil {
		return nil, err
	}
	settings := struct {
		Result struct {
			Settings UserSettings `json:"settings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &settings)
	return &settings.Result.Settings, err
}

// UsersSetMySettings - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	settings - list of all settings
func (s *ServerConnection) UsersSetMySettings(settings UserSettings) error {
	params := struct {
		Settings UserSettings `json:"settings"`
	}{settings}
	_, err := s.CallRaw("Users.setMySettings", params)
	return err
}

// UsersCheckWarnings - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	user - user data
// Return
//	errors - list of all warnings
func (s *ServerConnection) UsersCheckWarnings(user User) (ErrorList, error) {
	params := struct {
		User User `json:"user"`
	}{user}
	data, err := s.CallRaw("Users.checkWarnings", params)
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
