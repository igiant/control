package control

import "encoding/json"

// DirectoryServiceType - Common part, that can be shared between the products
type DirectoryServiceType string

const (
	WindowsActiveDirectory DirectoryServiceType = "WindowsActiveDirectory" // Windows Active Directory
	AppleDirectoryKerberos DirectoryServiceType = "AppleDirectoryKerberos" // Apple Open Directory with Kerberos authentication
	AppleDirectoryPassword DirectoryServiceType = "AppleDirectoryPassword" // Apple Open Directory with Password Server authentication
)

type DirectoryService struct {
	Enabled            bool                 `json:"enabled"`
	Type               DirectoryServiceType `json:"type"`
	DomainName         string               `json:"domainName"`
	UserName           string               `json:"userName"`
	Password           string               `json:"password"`
	UseSpecificServers bool                 `json:"useSpecificServers"` // valid for type WindowsActiveDirectory
	PrimaryServer      string               `json:"primaryServer"`
	SecondaryServer    string               `json:"secondaryServer"`
}

type DirectoryServiceAdvanced struct {
	LdapSecure       bool           `json:"ldapSecure"`
	KerberosRealm    OptionalString `json:"kerberosRealm"`    // valid for type AppleDirectoryKerberos
	LdapSearchSuffix OptionalString `json:"ldapSearchSuffix"` // valid for type AppleDirectory*
}

// DirectoryServiceConfiguration - Common part for testDomainController
type DirectoryServiceConfiguration struct {
	Id       KId                      `json:"id"`
	Service  DirectoryService         `json:"service"`
	Advanced DirectoryServiceAdvanced `json:"advanced"`
}

// Domain - Product dependent part
type Domain struct {
	Id     KId         `json:"id"`
	Status StoreStatus `json:"status"`
	/*@{ Shared part */
	Service  DirectoryService         `json:"service"`
	Advanced DirectoryServiceAdvanced `json:"advanced"`
	/*@}*/
	Description string `json:"description"`
	Primary     bool   `json:"primary"`
	/*@{ Type WindowsActiveDirectory and primary == true */
	AuthenticationOnly bool `json:"authenticationOnly"`
	NtAuthMode         bool `json:"ntAuthMode"`
	AdAutoImport       bool `json:"adAutoImport"`
	/*@}*/
	TemplateData UserData `json:"templateData"`
}

type DomainList []Domain

type TestResult struct {
	Successful   bool   `json:"successful"`
	ErrorMessage string `json:"errorMessage"`
}

// DomainsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	query - conditions and limits
// Return
//	list - list of domains and it's details
//	totalItems - count of all domains on server (before the start/limit applied)
func (s *ServerConnection) DomainsGet(query SearchQuery) (DomainList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Domains.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       DomainList `json:"list"`
			TotalItems int        `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// DomainsCreate - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domains - details for new domains. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) DomainsCreate(domains DomainList) (ErrorList, CreateResultList, error) {
	params := struct {
		Domains DomainList `json:"domains"`
	}{domains}
	data, err := s.CallRaw("Domains.create", params)
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

// DomainsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domainIds - ids of domains to be updated.
//	pattern - details for update. Field "kerio::web::KId" is ignored. All other fields except password must be filled and they are written to all domains specified by domainIds.
// Return
//	errors - list of errors \n
func (s *ServerConnection) DomainsSet(domainIds KIdList, pattern Domain) (ErrorList, error) {
	params := struct {
		DomainIds KIdList `json:"domainIds"`
		Pattern   Domain  `json:"pattern"`
	}{domainIds, pattern}
	data, err := s.CallRaw("Domains.set", params)
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

// DomainsRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	domainIds - ids of domains that should be removed
// Return
//	errors - list of errors \n
func (s *ServerConnection) DomainsRemove(domainIds KIdList) (ErrorList, error) {
	params := struct {
		DomainIds KIdList `json:"domainIds"`
	}{domainIds}
	data, err := s.CallRaw("Domains.remove", params)
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

// DomainsTestDomainController - @note: Members useSpecificServers, primaryServer and secondaryServer from DirectoryService are not used in this method
// Parameters
//	hostnames - List of servers, that will be used instead of servers in DirectoryService. Empty string means, that server should be autodetected
//	directory - domain parameters for connection test
// Return
//	errors - Various error messages related to given data, indexed by values in "hostnames" parameter, contains leading '.'
func (s *ServerConnection) DomainsTestDomainController(hostnames StringList, directory DirectoryServiceConfiguration) (ErrorList, error) {
	params := struct {
		Hostnames StringList                    `json:"hostnames"`
		Directory DirectoryServiceConfiguration `json:"directory"`
	}{hostnames, directory}
	data, err := s.CallRaw("Domains.testDomainController", params)
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

// DomainsApply - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	errors - list of errors \n
func (s *ServerConnection) DomainsApply() (ErrorList, error) {
	data, err := s.CallRaw("Domains.apply", nil)
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

// DomainsReset - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) DomainsReset() error {
	_, err := s.CallRaw("Domains.reset", nil)
	return err
}
