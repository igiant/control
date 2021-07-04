package control

import "encoding/json"

type DnsForwarder struct {
	Enabled    bool   `json:"enabled"`
	Domain     string `json:"domain"`     // network/mask for Reverse DNS query, other for Name DNS query
	Forwarders string `json:"forwarders"` // empty for Do not forward
}

type DnsForwarderList []DnsForwarder

type DnsConfig struct {
	ForwarderEnabled        bool             `json:"forwarderEnabled"`
	CacheEnabled            bool             `json:"cacheEnabled"`
	CustomForwardingEnabled bool             `json:"customForwardingEnabled"`
	CustomForwarders        DnsForwarderList `json:"customForwarders"`
	UseDomainControler      OptionalString   `json:"useDomainControler"` // read-only, enabled - true in case, that domain controller will be used
	//       value - dns queries, that matches this string, will be passed to domain controller
	HostsEnabled      bool   `json:"hostsEnabled"`
	DhcpLookupEnabled bool   `json:"dhcpLookupEnabled"`
	DomainName        string `json:"domainName"`
}

type DnsHost struct {
	Enabled     bool      `json:"enabled"`
	Id          KId       `json:"id"`
	Ip          IpAddress `json:"ip"`
	Hosts       string    `json:"hosts"`
	Description string    `json:"description"`
}

type DnsHostList []DnsHost

// DnsGet - Returns DNS configuration
func (s *ServerConnection) DnsGet() (*DnsConfig, error) {
	data, err := s.CallRaw("Dns.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config DnsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// DnsGetHosts - Returns DNS ip/hosts mapping
func (s *ServerConnection) DnsGetHosts() (DnsHostList, error) {
	data, err := s.CallRaw("Dns.getHosts", nil)
	if err != nil {
		return nil, err
	}
	hosts := struct {
		Result struct {
			Hosts DnsHostList `json:"hosts"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hosts)
	return hosts.Result.Hosts, err
}

// DnsSet - Stores DNS configuration
//	config - A structure containing all the settings of DND that sould be stored.
// Return
//	errors - list of errors
func (s *ServerConnection) DnsSet(config DnsConfig) (ErrorList, error) {
	params := struct {
		Config DnsConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("Dns.set", params)
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

// DnsSetHosts - Stores DNS ip/hosts mapping
//	hosts - list of hosts file entries to be stored
// Return
//	errors - list of errors
func (s *ServerConnection) DnsSetHosts(hosts DnsHostList) (ErrorList, error) {
	params := struct {
		Hosts DnsHostList `json:"hosts"`
	}{hosts}
	data, err := s.CallRaw("Dns.setHosts", params)
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

// DnsImportHosts - Imports DNS hosts records from file (hosts format)
//	fileId - id of uploaded file
//	clean - true, if current hosts records should be removed, false, if new records should be appended to current config
func (s *ServerConnection) DnsImportHosts(fileId string, clean bool) error {
	params := struct {
		FileId string `json:"fileId"`
		Clean  bool   `json:"clean"`
	}{fileId, clean}
	_, err := s.CallRaw("Dns.importHosts", params)
	return err
}

// DnsClearCache - Flushes DNS cache
func (s *ServerConnection) DnsClearCache() error {
	_, err := s.CallRaw("Dns.clearCache", nil)
	return err
}
