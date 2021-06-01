package control

// Global limits:
// 1. maximum of returned lines at once is 50000
// 2. maximal line length is 1024 (it does not include date, time, ID in the beginning of log line)
// Type of the log
// valid Connect names: config, debug, error, mail, security, spam, warning
// LogType - valid Control names: alert, config, connection, debug, dial, error, filter, http, security, sslvpn, warning, web
type LogType string

// LogItem - 1 log
type LogItem struct {
	LogName     LogType `json:"logName"`     // name of the log
	HasMessages bool    `json:"hasMessages"` // has the log messages?
}

// LogSet - List of valid logs
type LogSet []LogItem

// RotationPeriod - Period of rotation
type RotationPeriod string

const (
	RotateNever   RotationPeriod = "RotateNever"   // don't rotate
	RotateHourly  RotationPeriod = "RotateHourly"  // rotate hourly
	RotateDaily   RotationPeriod = "RotateDaily"   // rotate daily
	RotateWeekly  RotationPeriod = "RotateWeekly"  // rotate weekly
	RotateMonthly RotationPeriod = "RotateMonthly" // rotate monthly
)

// FacilityUnit - Available types of syslog facility according RFC 3164
type FacilityUnit string

const (
	FacilityKernel        FacilityUnit = "FacilityKernel"        //  0 = kernel messages
	FacilityUserLevel     FacilityUnit = "FacilityUserLevel"     //  1 = user-level messages
	FacilityMailSystem    FacilityUnit = "FacilityMailSystem"    //  2 = mail system
	FacilitySystemDaemons FacilityUnit = "FacilitySystemDaemons" //  3 = system daemons
	FacilitySecurity1     FacilityUnit = "FacilitySecurity1"     //  4 = security/authorization messages
	FacilityInternal      FacilityUnit = "FacilityInternal"      //  5 = messages generated internally by syslogd
	FacilityLinePrinter   FacilityUnit = "FacilityLinePrinter"   //  6 = line printer subsystem
	FacilityNetworkNews   FacilityUnit = "FacilityNetworkNews"   //  7 = network news subsystem
	FacilityUucpSubsystem FacilityUnit = "FacilityUucpSubsystem" //  8 = UUCP subsystem
	FacilityClockDaemon1  FacilityUnit = "FacilityClockDaemon1"  //  9 = clock daemon
	FacilitySecurity2     FacilityUnit = "FacilitySecurity2"     // 10 = security/authorization messages
	FacilityFtpDaemon     FacilityUnit = "FacilityFtpDaemon"     // 11 = FTP daemon
	FacilityNtpSubsystem  FacilityUnit = "FacilityNtpSubsystem"  // 12 = NTP subsystem
	FacilityLogAudit      FacilityUnit = "FacilityLogAudit"      // 13 = log audit
	FacilityLogAlert      FacilityUnit = "FacilityLogAlert"      // 14 = log alert
	FacilityClockDaemon2  FacilityUnit = "FacilityClockDaemon2"  // 15 = clock daemon
	FacilityLocal0        FacilityUnit = "FacilityLocal0"        // 16 = local use 0
	FacilityLocal1        FacilityUnit = "FacilityLocal1"        // 17 = local use 1
	FacilityLocal2        FacilityUnit = "FacilityLocal2"        // 18 = local use 2
	FacilityLocal3        FacilityUnit = "FacilityLocal3"        // 19 = local use 3
	FacilityLocal4        FacilityUnit = "FacilityLocal4"        // 20 = local use 4
	FacilityLocal5        FacilityUnit = "FacilityLocal5"        // 21 = local use 5
	FacilityLocal6        FacilityUnit = "FacilityLocal6"        // 22 = local use 6
	FacilityLocal7        FacilityUnit = "FacilityLocal7"        // 23 = local use 7
)

// SeverityUnit - Available types of severity
type SeverityUnit string

const (
	SeverityEmergency     SeverityUnit = "SeverityEmergency"
	SeverityAlert         SeverityUnit = "SeverityAlert"
	SeverityCritical      SeverityUnit = "SeverityCritical"
	SeverityError         SeverityUnit = "SeverityError"
	SeverityWarning       SeverityUnit = "SeverityWarning"
	SeverityNotice        SeverityUnit = "SeverityNotice"
	SeverityInformational SeverityUnit = "SeverityInformational"
	SeverityDebug         SeverityUnit = "SeverityDebug"
)

// LogFileSettings - general log settings
type LogFileSettings struct {
	Enabled  bool   `json:"enabled"`  // Is logging to file enabled
	FileName string `json:"fileName"` // log file name
}

// LogRotationSettings - log rotation settings
type LogRotationSettings struct {
	Period      RotationPeriod `json:"period"`      // How often does log rotate?
	MaxLogSize  int            `json:"maxLogSize"`  // Maximum log file size [MegaBytes]; Unlimited CAN be used
	RotateCount int            `json:"rotateCount"` // How many rotated files can be kept at most?; Unlimited CANNOT be used
}

// SyslogSettings - syslog settings
type SyslogSettings struct {
	Enabled     bool         `json:"enabled"`     // Syslog is [dis|en]abled
	ServerUrl   string       `json:"serverUrl"`   // Path to syslog server
	Facility    FacilityUnit `json:"facility"`    // which facility is message sent from
	Severity    SeverityUnit `json:"severity"`    // read-only; severity level of message
	Application string       `json:"application"` // user defined application name; it is 1*48PRINTUSASCII where PRINTUSASCII = %d33-126.
}

// LogSettings - Log file and output settings for 1 log
type LogSettings struct {
	General  LogFileSettings     `json:"general"`  // general log settings
	Rotation LogRotationSettings `json:"rotation"` // log rotation settings
	Syslog   SyslogSettings      `json:"syslog"`   // syslog settings
}

// HighlightColor - Highlight color definition in format RRGGBB
type HighlightColor string

// HighlightItem - Log highlighting item
type HighlightItem struct {
	Id             KId            `json:"id"`             // global identification
	Enabled        bool           `json:"enabled"`        // Rule is [dis|en]abled
	Description    string         `json:"description"`    // Text description
	Condition      string         `json:"condition"`      // Match condition
	IsRegex        bool           `json:"isRegex"`        // Is condition held as regular expression? (server does NOT check if regex is valid)
	Color          HighlightColor `json:"color"`          // Highlight matching log lines by this color
	IsOrderChanged bool           `json:"isOrderChanged"` // True if item order was changed by user
}

// HighlightRules - List of highlight items to be applied on all logs (global settings)
type HighlightRules []HighlightItem

// TreeLeaf - Leaf item of the tree
type TreeLeaf struct {
	Id          int    `json:"id"`          // leaf identification
	ParentName  string `json:"parentName"`  // name of the group
	Description string `json:"description"` // text after checkbox
	Enabled     bool   `json:"enabled"`     // leaf is [not] enabled
}

// TreeLeafList - sequence of leaves
type TreeLeafList []TreeLeaf

// ExportFormat - File type for log export
type ExportFormat string

const (
	PlainText ExportFormat = "PlainText" // export in plain text
	Html      ExportFormat = "Html"      // export in html
)

// LogRow - row of the log
type LogRow struct {
	Content   string         `json:"content"`   // 1 data row
	Highlight HighlightColor `json:"highlight"` // appropriate highlight color
}

type LogRowList []LogRow

// SearchStatus - Status of the Search
type SearchStatus string

const (
	ResultFound    SearchStatus = "ResultFound"    // the seach is finished and the match has been found
	Searching      SearchStatus = "Searching"      // the search still continues, the result is not available so far
	Cancelled      SearchStatus = "Cancelled"      // the search was cancelled by client
	ResultNotFound SearchStatus = "ResultNotFound" // the seach is finished but nothing was found
)

// Log object

// LogsCancelSearch - Cancel search on server (useful for large logs).
// Parameters
//	searchId - identifier from search()
func (s *ServerConnection) LogsCancelSearch(searchId string) error {
	params := struct {
		SearchId string `json:"searchId"`
	}{searchId}
	_, err := s.CallRaw("Logs.cancelSearch", params)
	return err
}

// LogsClear - Delete all log lines.
// Parameters
//	logName - unique name of the log
func (s *ServerConnection) LogsClear(logName LogType) error {
	params := struct {
		LogName LogType `json:"logName"`
	}{logName}
	_, err := s.CallRaw("Logs.clear", params)
	return err
}

// LogsExportLog - Exporting a given log.
// Parameters
//	logName - unique name of the log
//	fromLine - number of the line to start the search from;
//	countLines - number of lines to transfer; Unlimited - symbolic name for end of log
// Return
//	fileDownload - file download structure
func (s *ServerConnection) LogsExportLog(logName LogType, fromLine int, countLines int, exportFormat ExportFormat) (*Download, error) {
	params := struct {
		LogName      LogType      `json:"logName"`
		FromLine     int          `json:"fromLine"`
		CountLines   int          `json:"countLines"`
		ExportFormat ExportFormat `json:"exportFormat"`
	}{logName, fromLine, countLines, exportFormat}
	data, err := s.CallRaw("Logs.exportLog", params)
	if err != nil {
		return nil, err
	}
	fileDownload := struct {
		Result struct {
			FileDownload Download `json:"fileDownload"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileDownload)
	return &fileDownload.Result.FileDownload, err
}

// LogsExportLogRelative - Exporting a given log with relative download path.
// Parameters
//	logName - unique name of the log
//	fromLine - number of the line to start the search from;
//	countLines - number of lines to transfer; Unlimited - symbolic name for end of log
// Return
//	fileDownload - file download structure
func (s *ServerConnection) LogsExportLogRelative(logName LogType, fromLine int, countLines int, exportFormat ExportFormat) (*Download, error) {
	params := struct {
		LogName      LogType      `json:"logName"`
		FromLine     int          `json:"fromLine"`
		CountLines   int          `json:"countLines"`
		ExportFormat ExportFormat `json:"exportFormat"`
	}{logName, fromLine, countLines, exportFormat}
	data, err := s.CallRaw("Logs.exportLogRelative", params)
	if err != nil {
		return nil, err
	}
	fileDownload := struct {
		Result struct {
			FileDownload Download `json:"fileDownload"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileDownload)
	return &fileDownload.Result.FileDownload, err
}

// LogsGet - Obtain log data without linebreaks.
// Parameters
//	logName - unique name of the log
//	fromLine - number of the line to start from; if (fromLine == Unlimited) then fromline is end of log minus countLines
//	countLines - number of lines to transfer
// Return
//	viewport - list of log lines; count of lines = min(count, NUMBER_OF_CURRENT LINES - from)
//	totalItems - current count of all log lines
func (s *ServerConnection) LogsGet(logName LogType, fromLine int, countLines int) (LogRowList, int, error) {
	params := struct {
		LogName    LogType `json:"logName"`
		FromLine   int     `json:"fromLine"`
		CountLines int     `json:"countLines"`
	}{logName, fromLine, countLines}
	data, err := s.CallRaw("Logs.get", params)
	if err != nil {
		return nil, 0, err
	}
	viewport := struct {
		Result struct {
			Viewport   LogRowList `json:"viewport"`
			TotalItems int        `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &viewport)
	return viewport.Result.Viewport, viewport.Result.TotalItems, err
}

// LogsGetHighlightRules - Obtain a list of sorted highlighting rules.
// Return
//	rules - highlight rules
func (s *ServerConnection) LogsGetHighlightRules() (*HighlightRules, error) {
	data, err := s.CallRaw("Logs.getHighlightRules", nil)
	if err != nil {
		return nil, err
	}
	rules := struct {
		Result struct {
			Rules HighlightRules `json:"rules"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &rules)
	return &rules.Result.Rules, err
}

// LogsGetLogSet - Retrieve set of valid logs.
// Return
//	logSet - list of valid logs
func (s *ServerConnection) LogsGetLogSet() (*LogSet, error) {
	data, err := s.CallRaw("Logs.getLogSet", nil)
	if err != nil {
		return nil, err
	}
	logSet := struct {
		Result struct {
			LogSet LogSet `json:"logSet"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &logSet)
	return &logSet.Result.LogSet, err
}

// LogsGetMessages - Obtain log message settings; make sense only if LogItem.hasMessages == true.
// Return
//	messages - tree of log messages
func (s *ServerConnection) LogsGetMessages() (TreeLeafList, error) {
	data, err := s.CallRaw("Logs.getMessages", nil)
	if err != nil {
		return nil, err
	}
	messages := struct {
		Result struct {
			Messages TreeLeafList `json:"messages"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &messages)
	return messages.Result.Messages, err
}

// LogsGetSearchProgress - Clears timeout for search() and obtains status of the search.
func (s *ServerConnection) LogsGetSearchProgress() error {
	_, err := s.CallRaw("Logs.getSearchProgress", nil)
	return err
}

// LogsGetSettings - Obtain log settings.
// Parameters
//	logName - unique name of the log
// Return
//	currentSettings - current valid settings (or undefined data on failure)
func (s *ServerConnection) LogsGetSettings(logName LogType) (*LogSettings, error) {
	params := struct {
		LogName LogType `json:"logName"`
	}{logName}
	data, err := s.CallRaw("Logs.getSettings", params)
	if err != nil {
		return nil, err
	}
	currentSettings := struct {
		Result struct {
			CurrentSettings LogSettings `json:"currentSettings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &currentSettings)
	return &currentSettings.Result.CurrentSettings, err
}

// LogsSearch - Start searching for a string in a given log; The search exists 1 minute unless prolonged by getSearchProgress.
// Parameters
//	logName - unique name of the log
//	what - searched string
//	fromLine - line to start searching from; fromLine>toLine means search up; Unlimited - symbolic name for end of log
//	toLine - line to start searching from; fromLine<toLine means search down
//	forward - direction of the search; true = forward, false = backward
// Return
//	searchId - identifier that can be used for cancelSearch and getSearchProgress
func (s *ServerConnection) LogsSearch(logName LogType, what string, fromLine int, toLine int, forward bool) (string, error) {
	params := struct {
		LogName  LogType `json:"logName"`
		What     string  `json:"what"`
		FromLine int     `json:"fromLine"`
		ToLine   int     `json:"toLine"`
		Forward  bool    `json:"forward"`
	}{logName, what, fromLine, toLine, forward}
	data, err := s.CallRaw("Logs.search", params)
	if err != nil {
		return "", err
	}
	searchId := struct {
		Result struct {
			SearchId string `json:"searchId"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &searchId)
	return searchId.Result.SearchId, err
}

// LogsSetHighlightRules - Set highlighting rules, rules have to be sorted purposely, the only way to change a rule is to change the whole ruleset.
// Parameters
//	rules - highlight rules (ordered by priority)
func (s *ServerConnection) LogsSetHighlightRules(rules HighlightRules) error {
	params := struct {
		Rules HighlightRules `json:"rules"`
	}{rules}
	_, err := s.CallRaw("Logs.setHighlightRules", params)
	return err
}

// LogsSetMessages - Change log message settings; makes sense only if LogItem.hasMessages == true.
// Parameters
//	messages - tree of log messages
func (s *ServerConnection) LogsSetMessages(messages TreeLeafList) error {
	params := struct {
		Messages TreeLeafList `json:"messages"`
	}{messages}
	_, err := s.CallRaw("Logs.setMessages", params)
	return err
}

// LogsSetSettings - Change log settings.
// Parameters
//	logName - unique name of the log
//	newSettings
func (s *ServerConnection) LogsSetSettings(logName LogType, newSettings LogSettings) error {
	params := struct {
		LogName     LogType     `json:"logName"`
		NewSettings LogSettings `json:"newSettings"`
	}{logName, newSettings}
	_, err := s.CallRaw("Logs.setSettings", params)
	return err
}
