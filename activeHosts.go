package control

import "encoding/json"

type ActiveHostType string

const (
	ActiveHostFirevall  ActiveHostType = "ActiveHostFirevall"
	ActiveHostVpnClient ActiveHostType = "ActiveHostVpnClient"
	ActiveHostHost      ActiveHostType = "ActiveHostHost"
	ActiveHostGuest     ActiveHostType = "ActiveHostGuest"
)

type AuthMethodType string

const (
	AuthMethodWeb       AuthMethodType = "AuthMethodWeb"       /* Plaintext */
	AuthMethodSslWeb    AuthMethodType = "AuthMethodSslWeb"    /* SSL */
	AuthMethodNtlm      AuthMethodType = "AuthMethodNtlm"      /* NTLM */
	AuthMethodProxy     AuthMethodType = "AuthMethodProxy"     /* Proxy */
	AuthMethodAutomatic AuthMethodType = "AuthMethodAutomatic" /* Automatic */
	AuthMethodVpnClient AuthMethodType = "AuthMethodVpnClient" /* VPN Client */
	AuthMethodSso       AuthMethodType = "AuthMethodSso"       /* "Kerio Unity Sign-On" */
	AuthMethodApi       AuthMethodType = "AuthMethodApi"       /* webadmin API */
	AuthMethodRadius    AuthMethodType = "AuthMethodRadius"    /* webadmin API */
	AuthMethodNone      AuthMethodType = "AuthMethodNone"      /* "" */
)

type ActiveHost struct {
	Id                KId            `json:"id"`
	Type              ActiveHostType `json:"type"`
	Ip                IpAddress      `json:"ip"`
	Ip6Addresses      Ip6AddressList `json:"ip6Addresses"`
	Hostname          string         `json:"hostname"`
	MacAddress        string         `json:"macAddress"`
	IpAddressFromDHCP bool           `json:"ipAddressFromDHCP"`
	StartTime         DateTimeConfig `json:"startTime"`
	InactivityTime    int            `json:"inactivityTime"`
	User              UserReference  `json:"user"`
	LoginTime         DateTimeConfig `json:"loginTime"`
	LoginDuration     int            `json:"loginDuration"`
	AuthMethod        AuthMethodType `json:"authMethod"`
	Connections       int            `json:"connections"`
	TotalDownload     float64        `json:"totalDownload"`
	TotalUpload       float64        `json:"totalUpload"`
	CurrentDownload   float64        `json:"currentDownload"`
	CurrentUpload     float64        `json:"currentUpload"`
}

type ActiveHostList []ActiveHost

type ActivityType string

const (
	ActivityTypeWeb              ActivityType = "ActivityTypeWeb" /* not used in current implementation */
	ActivityTypeWebSearch        ActivityType = "ActivityTypeWebSearch"
	ActivityTypeMail             ActivityType = "ActivityTypeMail"
	ActivityTypeDownload         ActivityType = "ActivityTypeDownload"
	ActivityTypeUpload           ActivityType = "ActivityTypeUpload"
	ActivityTypeMultimedia       ActivityType = "ActivityTypeMultimedia"
	ActivityTypeP2p              ActivityType = "ActivityTypeP2p"
	ActivityTypeRemoteAccess     ActivityType = "ActivityTypeRemoteAccess"
	ActivityTypeVpn              ActivityType = "ActivityTypeVpn"
	ActivityTypeInstantMessaging ActivityType = "ActivityTypeInstantMessaging"
	ActivityTypeHugeConnection   ActivityType = "ActivityTypeHugeConnection"
	ActivityTypeMailConnection   ActivityType = "ActivityTypeMailConnection"
	ActivityTypeP2pAttempt       ActivityType = "ActivityTypeP2pAttempt"
	ActivityTypeWebConnection    ActivityType = "ActivityTypeWebConnection"
	ActivityTypeHTTPConnection   ActivityType = "ActivityTypeHTTPConnection"
	ActivityTypeWebMultimedia    ActivityType = "ActivityTypeWebMultimedia"
	ActivityTypeSip              ActivityType = "ActivityTypeSip"
	ActivityTypeSocialNetwork    ActivityType = "ActivityTypeSocialNetwork"
)

type Activity struct {
	Id          KId                `json:"id"`
	Time        TimeHMS            `json:"time"`
	Type        ActivityType       `json:"type"`
	Description LocalizableMessage `json:"description"`
	Url         string             `json:"url"`
}

type ActivityList []Activity

// ActiveHostsGet - Returns Active Hosts data
// Parameters
//	query - filter/sort query
//	refresh - true in case, that data snapshot have to be refreshed
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) ActiveHostsGet(query SearchQuery, refresh bool) (ActiveHostList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query   SearchQuery `json:"query"`
		Refresh bool        `json:"refresh"`
	}{query, refresh}
	data, err := s.CallRaw("ActiveHosts.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ActiveHostList `json:"list"`
			TotalItems int            `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ActiveHostsGetActivityList - Returns Activities for specified host
// Parameters
//	id - Active Host id returned by get
// Return
//	list - output data
func (s *ServerConnection) ActiveHostsGetActivityList(id KId) (ActivityList, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("ActiveHosts.getActivityList", params)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List ActivityList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// ActiveHostsGetHistogram - Returns throughput Histogram for specified host
// Parameters
//	id - Active Host id returned by get
// Return
//	hist - samples of traffic rate for given host
func (s *ServerConnection) ActiveHostsGetHistogram(histogramType HistogramType, id KId) (*Histogram, error) {
	params := struct {
		HistogramType HistogramType `json:"histogramType"`
		Id            KId           `json:"id"`
	}{histogramType, id}
	data, err := s.CallRaw("ActiveHosts.getHistogram", params)
	if err != nil {
		return nil, err
	}
	hist := struct {
		Result struct {
			Hist Histogram `json:"hist"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hist)
	return &hist.Result.Hist, err
}

// ActiveHostsGetHistogramInc - Returns throughput Histogram for specified host
// Parameters
//	startSampleTime - Specifies starting time for returned data
//	id - rule or interface id returned in item.id member by get
// Return
//	hist - output data
//	sampleTime - Returns first time, that is not included in returned data (pass it to the same method again as a lastSampleTime to obtain only new data since last request)
func (s *ServerConnection) ActiveHostsGetHistogramInc(histogramIntervalType HistogramIntervalType, id KId, startSampleTime DateTimeStamp) (*Histogram, *DateTimeStamp, error) {
	params := struct {
		HistogramIntervalType HistogramIntervalType `json:"histogramIntervalType"`
		Id                    KId                   `json:"id"`
		StartSampleTime       DateTimeStamp         `json:"startSampleTime"`
	}{histogramIntervalType, id, startSampleTime}
	data, err := s.CallRaw("ActiveHosts.getHistogramInc", params)
	if err != nil {
		return nil, nil, err
	}
	hist := struct {
		Result struct {
			Hist       Histogram     `json:"hist"`
			SampleTime DateTimeStamp `json:"sampleTime"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hist)
	return &hist.Result.Hist, &hist.Result.SampleTime, err
}

// ActiveHostsLogout - Logout users from specified hosts / empty for all hosts
// Parameters
//	ids - Active Host ids returned by get or empty for logout all hosts
func (s *ServerConnection) ActiveHostsLogout(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("ActiveHosts.logout", params)
	return err
}

// ActiveHostsLogin - Logs in user for specified host
// Parameters
//	hostId - internal identifier of a host computer in network
//	userName - Name of a user to be loged from given host specified by hostId (including domain if needed)
func (s *ServerConnection) ActiveHostsLogin(hostId KId, userName string) error {
	params := struct {
		HostId   KId    `json:"hostId"`
		UserName string `json:"userName"`
	}{hostId, userName}
	_, err := s.CallRaw("ActiveHosts.login", params)
	return err
}
