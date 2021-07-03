package control

import "encoding/json"

type AlertRow struct {
	Id      KId    `json:"id"`
	Date    string `json:"date"`
	Alert   string `json:"alert"`
	Details string `json:"details"`
}

type AlertRowList []AlertRow

type AlertType struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

type AlertTypeList []AlertType

type AlertEventRuleType string

const (
	AlertTraffic AlertEventRuleType = "AlertTraffic"
	AlertContent AlertEventRuleType = "AlertContent"
)

type AlertRuleEvent struct {
	RuleType AlertEventRuleType `json:"ruleType"`
	Rule     IdReference        `json:"rule"`
}

type AlertRuleEventList []AlertRuleEvent

type AlertLogEvent struct {
	Log       KId    `json:"log"`
	Name      string `json:"name"`
	Interval  int    `json:"interval"`
	IsRegex   bool   `json:"isRegex"`
	Condition string `json:"condition"`
}

type AlertLogEventList []AlertLogEvent

type AlertSetting struct {
	Id             KId                `json:"id"`
	Enabled        bool               `json:"enabled"`
	Addressee      Addressee          `json:"addressee"`
	AlertList      KIdList            `json:"alertList"`
	RuleEventList  AlertRuleEventList `json:"ruleEventList"`
	LogEventList   AlertLogEventList  `json:"logEventList"`
	ValidTimeRange IdReference        `json:"validTimeRange"`
}

type AlertSettingList []AlertSetting

// AlertsGet - Returns Alert Messages data
// Parameters
//	query - paging query (sorting is not possible and it is ignored)
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) AlertsGet(query SearchQuery) (AlertRowList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("Alerts.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       AlertRowList `json:"list"`
			TotalItems int          `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// AlertsGetContent - Returns content of given Alert Message as formatted HTML
// Parameters
//	id - ID of given alert
// Return
//	content - output data
func (s *ServerConnection) AlertsGetContent(id KId) (string, error) {
	params := struct {
		Id KId `json:"id"`
	}{id}
	data, err := s.CallRaw("Alerts.getContent", params)
	if err != nil {
		return "", err
	}
	content := struct {
		Result struct {
			Content string `json:"content"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &content)
	return content.Result.Content, err
}

// AlertsGetAlertTypes - Returns list of possible Alerts types
// Return
//  types - list of possible Alerts types
func (s *ServerConnection) AlertsGetAlertTypes() (AlertTypeList, error) {
	data, err := s.CallRaw("Alerts.getAlertTypes", nil)
	if err != nil {
		return nil, err
	}
	types := struct {
		Result struct {
			Types AlertTypeList `json:"types"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &types)
	return types.Result.Types, err
}

// AlertsGetSettings - Returns list of user defined alert handling
// Return
//  config - list of user defined alert handling
func (s *ServerConnection) AlertsGetSettings() (AlertSettingList, error) {
	data, err := s.CallRaw("Alerts.getSettings", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config AlertSettingList `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return config.Result.Config, err
}

// AlertsSetSettings - Stores list of user defined alert handling
// Parameters
//	config - structure with complete alerts settings
// Return
//	errors - list of items that cannot be changed. \n
func (s *ServerConnection) AlertsSetSettings(config AlertSettingList) (ErrorList, error) {
	params := struct {
		Config AlertSettingList `json:"config"`
	}{config}
	data, err := s.CallRaw("Alerts.setSettings", params)
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

// AlertsGetDefaultLanguage - Returns default language for Alert emails
// Return
//  lang - default language for Alert emails
func (s *ServerConnection) AlertsGetDefaultLanguage() (string, error) {
	data, err := s.CallRaw("Alerts.getDefaultLanguage", nil)
	if err != nil {
		return "", err
	}
	lang := struct {
		Result struct {
			Lang string `json:"lang"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &lang)
	return lang.Result.Lang, err
}

// AlertsSetDefaultLanguage - Stores default language for Alert emails
// Parameters
//  lang - default language for Alert emails
func (s *ServerConnection) AlertsSetDefaultLanguage(lang string) error {
	params := struct {
		Lang string `json:"lang"`
	}{lang}
	_, err := s.CallRaw("Alerts.setDefaultLanguage", params)
	return err
}
