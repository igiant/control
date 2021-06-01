package control

type ReportConfig struct {
	DailyEnabled   bool `json:"dailyEnabled"`
	WeeklyEnabled  bool `json:"weeklyEnabled"`
	MonthlyEnabled bool `json:"monthlyEnabled"`
	OnlineAccess   bool `json:"onlineAccess"`
}

type StarReport struct {
	Id           KId               `json:"id"`
	Enabled      bool              `json:"enabled"`
	Addressee    Addressee         `json:"addressee"`
	Users        UserReferenceList `json:"users"`
	AllData      bool              `json:"allData"`
	ReportConfig ReportConfig      `json:"reportConfig"`
}

type StarReportList []StarReport

// StarReportsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	list - list of entries
//	allUsers - Structure telling whether to send individual reports to each person reularly and wheter to allow them to see their reposrt online.
func (s *ServerConnection) StarReportsGet() (StarReportList, *ReportConfig, error) {
	data, err := s.CallRaw("StarReports.get", nil)
	if err != nil {
		return nil, nil, err
	}
	list := struct {
		Result struct {
			List     StarReportList `json:"list"`
			AllUsers ReportConfig   `json:"allUsers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, &list.Result.AllUsers, err
}

// StarReportsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	reports - list of report configurations to be send regularly.
//	allUsers - Structure telling whether to send individual reports to each person reularly and wheter to allow them to see their reposrt online.
// Return
//	errors - list of errors \n
func (s *ServerConnection) StarReportsSet(reports StarReportList, allUsers ReportConfig) (ErrorList, error) {
	params := struct {
		Reports  StarReportList `json:"reports"`
		AllUsers ReportConfig   `json:"allUsers"`
	}{reports, allUsers}
	data, err := s.CallRaw("StarReports.set", params)
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

// StarReportsSend - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	reports - list of addressees and types of the report
//	language - language of the report
// Return
//	errors - list of errors \n
func (s *ServerConnection) StarReportsSend(reports StarReportList, language string) (ErrorList, error) {
	params := struct {
		Reports  StarReportList `json:"reports"`
		Language string         `json:"language"`
	}{reports, language}
	data, err := s.CallRaw("StarReports.send", params)
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
