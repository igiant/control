package control

import "encoding/json"

// HttpsConfig - HTTPS configuration
type HttpsConfig struct {
	Enabled         bool                `json:"enabled"`
	IsExclusionMode bool                `json:"isExclusionMode"`
	IpAddressGroup  OptionalIdReference `json:"ipAddressGroup"`
	UserExceptions  UserReferenceList   `json:"userExceptions"`
	Disclaimer      bool                `json:"disclaimer"`
}

// SafeSearchConfig - SafeSearch configuration
type SafeSearchConfig struct {
	Enabled        bool              `json:"enabled"`
	UserExceptions UserReferenceList `json:"userExceptions"`
}

// UrlWhiteListEntry - Web Filter configuration
type UrlWhiteListEntry struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type UrlWhiteList []UrlWhiteListEntry

type UrlFilterStatus string

const (
	UrlFilterNotLicensed  UrlFilterStatus = "UrlFilterNotLicensed"  /* Whole tab disabled (there is feature section in OUR license file (or Control run in trial)) */
	UrlFilterNotActivated UrlFilterStatus = "UrlFilterNotActivated" /* licensed, but ticket not accepted (activated) by filter */
	UrlFilterActivating   UrlFilterStatus = "UrlFilterActivating"   /* licensed, but ticket activation in progress */
	UrlFilterActivated    UrlFilterStatus = "UrlFilterActivated"    /* licensed and ticket accepted (activated) by filter */
)

type UrlFilterConfig struct {
	WhiteList                    UrlWhiteList    `json:"whiteList"`
	Status                       UrlFilterStatus `json:"status"`
	ActivationErrorDescr         string          `json:"activationErrorDescr"`
	Enabled                      bool            `json:"enabled"`
	StatisticsEnabled            bool            `json:"statisticsEnabled"`
	AllowMiscategorizedReporting bool            `json:"allowMiscategorizedReporting"`
	AppidEnabled                 bool            `json:"appidEnabled"`
}

// ApplicationType - List of applications
type ApplicationType string

const (
	ApplicationWebFilterCategory ApplicationType = "ApplicationWebFilterCategory"
	ApplicationProtocol          ApplicationType = "ApplicationProtocol"
)

type ApplicationTypeList []ApplicationType

type ContentApplication struct {
	Id          int                 `json:"id"`
	Heuristic   bool                `json:"heuristic"`
	Group       string              `json:"group"`
	SubGroup    string              `json:"subGroup"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Types       ApplicationTypeList `json:"types"`
}

type ContentApplicationList []ContentApplication

type ContentConditionEntityType string

const (
	ContentConditionEntityApplication ContentConditionEntityType = "ContentConditionEntityApplication"
	ContentConditionEntityFileName    ContentConditionEntityType = "ContentConditionEntityFileName"
	ContentConditionEntityFileGroup   ContentConditionEntityType = "ContentConditionEntityFileGroup"
	ContentConditionEntityUrl         ContentConditionEntityType = "ContentConditionEntityUrl"
	ContentConditionEntityUrlGroup    ContentConditionEntityType = "ContentConditionEntityUrlGroup"
)

type ContentEntityUrlType string

const (
	ContentEntityUrlWildcard ContentEntityUrlType = "ContentEntityUrlWildcard"
	ContentEntityUrlRegex    ContentEntityUrlType = "ContentEntityUrlRegex"
	ContentEntityUrlHostname ContentEntityUrlType = "ContentEntityUrlHostname"
)

type ApplicationList []int

type ContentConditionEntity struct {
	Type ContentConditionEntityType `json:"type"`
	/*@{ ContentConditionEntityApplication */
	Applications ApplicationList `json:"applications"`
	/*@}*/
	/*@{ ContentConditionEntityFileType & ContentEntityUrl */
	Value string `json:"value"`
	/*@}*/
	/*@{ ContentEntityUrl */
	UrlType      ContentEntityUrlType `json:"urlType"`
	MatchSecured bool                 `json:"matchSecured"`
	/*@}*/
	/*@{ ContentConditionEntityUrlGroup */
	UrlGroup IdReference `json:"urlGroup"`
	/*@}*/
}

type ContentConditionEntityList []ContentConditionEntity

type ContentCondition struct {
	Type     RuleConditionType          `json:"type"`
	Entities ContentConditionEntityList `json:"entities"`
}

type SourceConditonEntityType string

const (
	SourceConditonEntityAddressGroup SourceConditonEntityType = "SourceConditonEntityAddressGroup"
	SourceConditonEntityUsers        SourceConditonEntityType = "SourceConditonEntityUsers"
	SourceConditonEntityGuests       SourceConditonEntityType = "SourceConditonEntityGuests"
)

type SourceConditonEntity struct {
	Type SourceConditonEntityType `json:"type"`
	/*@{ IP address group */
	IpAddressGroup IdReference `json:"ipAddressGroup"`
	/*@}*/
	/*@{ users */
	UserType UserConditionType `json:"userType"` // @see Users.idl, used values: AuthenticatedUsers, SelectedUsers
	User     UserReference     `json:"user"`     // @see UserManager
	/*@}*/
}

type SourceConditonEntityList []SourceConditonEntity

type SourceCondition struct {
	Type     RuleConditionType        `json:"type"`
	Entities SourceConditonEntityList `json:"entities"`
}

type DenialCondition struct {
	DenialText  string         `json:"denialText"`
	RedirectUrl OptionalString `json:"redirectUrl"`
	SendEmail   bool           `json:"sendEmail"`
}

type ContentRule struct {
	Id KId `json:"id"`
	/*@{ name */
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	/*@{ action */
	Action             RuleAction      `json:"action"`
	LogEnabled         bool            `json:"logEnabled"`
	SkipAvScan         bool            `json:"skipAvScan"`
	SkipKeywords       bool            `json:"skipKeywords"`
	SkipAuthentication bool            `json:"skipAuthentication"`
	DenialCondition    DenialCondition `json:"denialCondition"`
	/*@{ Conditions */
	ContentCondition ContentCondition `json:"contentCondition"`
	SourceCondition  SourceCondition  `json:"sourceCondition"`
	ValidTimeRange   IdReference      `json:"validTimeRange"`
}

type ContentRuleList []ContentRule

// ContentFilterGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	list - list of rule and it's details
func (s *ServerConnection) ContentFilterGet() (ContentRuleList, error) {
	data, err := s.CallRaw("ContentFilter.get", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List ContentRuleList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// ContentFilterSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	rules - list of rules and it's details
// Return
//	errors - list of errors occured during method call
func (s *ServerConnection) ContentFilterSet(rules ContentRuleList) (ErrorList, error) {
	params := struct {
		Rules ContentRuleList `json:"rules"`
	}{rules}
	data, err := s.CallRaw("ContentFilter.set", params)
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

// ContentFilterGetCollisions - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ContentFilterGetCollisions() (CollisionList, error) {
	data, err := s.CallRaw("ContentFilter.getCollisions", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List CollisionList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// ContentFilterGetContentApplicationList - 8000 Internal error.  - "Internal error." - unable to get page list of webfilter/application categories and applications
// Return
//	categories - list of webfilter/application categories and applications
func (s *ServerConnection) ContentFilterGetContentApplicationList() (ContentApplicationList, error) {
	data, err := s.CallRaw("ContentFilter.getContentApplicationList", nil)
	if err != nil {
		return nil, err
	}
	categories := struct {
		Result struct {
			Categories ContentApplicationList `json:"categories"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &categories)
	return categories.Result.Categories, err
}

// ContentFilterGetFilenameGroups - @deprecated use FilenameGroups::get() instead
// Return
//	groups - list of filename groups
func (s *ServerConnection) ContentFilterGetFilenameGroups() (FilenameGroupList, error) {
	data, err := s.CallRaw("ContentFilter.getFilenameGroups", nil)
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

// ContentFilterGetUrlFilterConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	config - configuration values
func (s *ServerConnection) ContentFilterGetUrlFilterConfig() (*UrlFilterConfig, error) {
	data, err := s.CallRaw("ContentFilter.getUrlFilterConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config UrlFilterConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ContentFilterSetUrlFilterConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - configuration values
func (s *ServerConnection) ContentFilterSetUrlFilterConfig(config UrlFilterConfig) error {
	params := struct {
		Config UrlFilterConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ContentFilter.setUrlFilterConfig", params)
	return err
}

// ContentFilterReportMiscategorizedUrl - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	url - URL, that is miscategorized
//	categoryIds - up to 3 suggested categories. Can be empty, if new category is not known
func (s *ServerConnection) ContentFilterReportMiscategorizedUrl(url string, categoryIds IntegerList) error {
	params := struct {
		Url         string      `json:"url"`
		CategoryIds IntegerList `json:"categoryIds"`
	}{url, categoryIds}
	_, err := s.CallRaw("ContentFilter.reportMiscategorizedUrl", params)
	return err
}

// ContentFilterGetUrlCategories - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	url - checked URL
// Return
//	categoryIds - list of categories, to which given URL belongs
func (s *ServerConnection) ContentFilterGetUrlCategories(url string) (IntegerList, error) {
	params := struct {
		Url string `json:"url"`
	}{url}
	data, err := s.CallRaw("ContentFilter.getUrlCategories", params)
	if err != nil {
		return nil, err
	}
	categoryIds := struct {
		Result struct {
			CategoryIds IntegerList `json:"categoryIds"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &categoryIds)
	return categoryIds.Result.CategoryIds, err
}

// ContentFilterGetHttpsConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	config - configuration values
func (s *ServerConnection) ContentFilterGetHttpsConfig() (*HttpsConfig, error) {
	data, err := s.CallRaw("ContentFilter.getHttpsConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config HttpsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ContentFilterSetHttpsConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - configuration values
func (s *ServerConnection) ContentFilterSetHttpsConfig(config HttpsConfig) error {
	params := struct {
		Config HttpsConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ContentFilter.setHttpsConfig", params)
	return err
}

// ContentFilterGetSafeSearchConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	config - configuration values
func (s *ServerConnection) ContentFilterGetSafeSearchConfig() (*SafeSearchConfig, error) {
	data, err := s.CallRaw("ContentFilter.getSafeSearchConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config SafeSearchConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ContentFilterSetSafeSearchConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - configuration values
func (s *ServerConnection) ContentFilterSetSafeSearchConfig(config SafeSearchConfig) error {
	params := struct {
		Config SafeSearchConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ContentFilter.setSafeSearchConfig", params)
	return err
}

// ContentFilterClearHttpsCertCache - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ContentFilterClearHttpsCertCache() error {
	_, err := s.CallRaw("ContentFilter.clearHttpsCertCache", nil)
	return err
}
