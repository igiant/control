package control

import "encoding/json"

type UserFormatType string

const (
	UserFormatFL  UserFormatType = "UserFormatFL"
	UserFormatFLU UserFormatType = "UserFormatFLU"
	UserFormatFLD UserFormatType = "UserFormatFLD"
	UserFormatLF  UserFormatType = "UserFormatLF"
	UserFormatLFU UserFormatType = "UserFormatLFU"
	UserFormatLFD UserFormatType = "UserFormatLFD"
)

type AccountingConfig struct {
	Enabled            bool                `json:"enabled"`
	ActivityLogEnabled bool                `json:"activityLogEnabled"`
	MaxAge             int                 `json:"maxAge"`
	UserFormat         UserFormatType      `json:"userFormat"`
	GatheredGroups     UserReferenceList   `json:"gatheredGroups"`
	StartWeekDay       Day                 `json:"startWeekDay"`  // @see TimeRangeManager
	StartMonthDay      int                 `json:"startMonthDay"` // 1..28
	StarReportLanguage string              `json:"starReportLanguage"`
	ValidTimeRange     OptionalIdReference `json:"validTimeRange"`
	IpAddressGroup     OptionalIdReference `json:"ipAddressGroup"`
	UserExceptions     UserReferenceList   `json:"userExceptions"`
	UrlGroup           OptionalIdReference `json:"urlGroup"`
}

// AccountingGet - Returns Accounting configuration
// Return
//  config - Accounting configuration
func (s *ServerConnection) AccountingGet() (*AccountingConfig, error) {
	data, err := s.CallRaw("Accounting.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config AccountingConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// AccountingSet - Stores Accounting configuration
//  config - Accounting configuration
func (s *ServerConnection) AccountingSet(config AccountingConfig) (ErrorList, error) {
	params := struct {
		Config AccountingConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("Accounting.set", params)
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
