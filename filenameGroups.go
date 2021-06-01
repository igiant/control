package control

import "encoding/json"

type OptionalStringList struct {
	Enabled bool       `json:"enabled"`
	Value   StringList `json:"value"`
}

type OptionalIpAddress OptionalString

type OptionalIpAddressList OptionalStringList

type CredentialsConfig struct {
	UserName string `json:"userName"`
	/** it should never be filled by get() method and transferred to the client */
	Password string `json:"password"`
	/** if true (user has changed the password), set() method write password to configuration, otherwise let it untouched */
	PasswordChanged bool `json:"passwordChanged"`
}

type IdReference struct {
	Id      KId    `json:"id"`
	Name    string `json:"name"`
	Invalid bool   `json:"invalid"`
}

type IdReferenceList []IdReference

type OptionalIdReference struct {
	Enabled bool        `json:"enabled"`
	Value   IdReference `json:"value"`
}

type DataStatistic struct {
	Today    float64 `json:"today"`
	TodayIn  float64 `json:"todayIn"`
	TodayOut float64 `json:"todayOut"`
	Week     float64 `json:"week"`
	WeekIn   float64 `json:"weekIn"`
	WeekOut  float64 `json:"weekOut"`
	Month    float64 `json:"month"`
	MonthIn  float64 `json:"monthIn"`
	MonthOut float64 `json:"monthOut"`
	Total    float64 `json:"total"`
	TotalIn  float64 `json:"totalIn"`
	TotalOut float64 `json:"totalOut"`
}

type TimeHMS struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
	Sec  int `json:"sec"`
}

type DateTimeConfig struct {
	Date Date    `json:"date"`
	Time TimeHMS `json:"time"`
}

type TimeSpan struct {
	IsValid bool `json:"isValid"` // false - time span cannot be computed (no update yet, update is in the future)
	Days    int  `json:"days"`
	Hours   int  `json:"hours"`
	Minutes int  `json:"minutes"`
}

type RuleAction string

const (
	Allow  RuleAction = "Allow"
	Deny   RuleAction = "Deny"
	Drop   RuleAction = "Drop"
	NotSet RuleAction = "NotSet"
)

type VpnType string

const (
	VpnKerio VpnType = "VpnKerio"
	VpnIpsec VpnType = "VpnIpsec"
)

type HistogramData struct {
	Inbound  float64 `json:"inbound"`
	Outbound float64 `json:"outbound"`
}

type HistogramDataList []HistogramData

type Histogram struct {
	Units      ByteUnits         `json:"units"`
	AverageIn  float64           `json:"averageIn"`
	AverageOut float64           `json:"averageOut"`
	MaxIn      float64           `json:"maxIn"`
	MaxOut     float64           `json:"maxOut"`
	Data       HistogramDataList `json:"data"`
}

type RuleConditionType string

const (
	RuleAny              RuleConditionType = "RuleAny"
	RuleSelectedEntities RuleConditionType = "RuleSelectedEntities"
	RuleInvalidCondition RuleConditionType = "RuleInvalidCondition" // nothing
)

type Ip6Address string

type Ip6AddressList []Ip6Address

type Collision struct {
	Rule           IdReference `json:"rule"`
	OverlappedRule IdReference `json:"overlappedRule"`
}

type Password struct {
	Value string `json:"value"`
	IsSet bool   `json:"isSet"`
}

type CollisionList []Collision

type FilenameGroup struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

type FilenameGroupList []FilenameGroup

// FilenameGroupsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	groups - list of filename groups
func (s *ServerConnection) FilenameGroupsGet() (FilenameGroupList, error) {
	data, err := s.CallRaw("FilenameGroups.get", nil)
	if err != nil {
		return nil, err
	}
	groups := struct {
		Result struct {
			Groups FilenameGroupList `json:"groups"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &groups)
	return groups.Result.Groups, err
}
