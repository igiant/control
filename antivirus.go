package control

import "encoding/json"

// AntivirusOption - Common part, that can be shared between the products
type AntivirusOption struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	DefaultValue string `json:"defaultValue"` // read only value
}

type AntivirusOptionList []AntivirusOption

type ExternalAntivirus struct {
	Id                  string              `json:"id"`          // example: avir_avg
	Description         string              `json:"description"` // example: AVG Email Server Edition
	AreOptionsAvailable bool                `json:"areOptionsAvailable"`
	Options             AntivirusOptionList `json:"options"`
}

type ExternalAntivirusList []ExternalAntivirus

type InternalAntivirus struct {
	UpdateCheckInterval OptionalLong `json:"updateCheckInterval"` // should we periodically ask for a new version? + update checking period in hours
	Available           bool         `json:"available"`           // license is valid for internal antivirus, it is present: true - checkbox "Use integrated..." is enabled
	Expired             bool         `json:"expired"`             // license is not valid for McAfee: message "McAfee(R) antivirus subscription expired." is displayed
}

type AntivirusStatus string

const (
	AntivirusOk             AntivirusStatus = "AntivirusOk"             // no message is needed
	AntivirusNotActive      AntivirusStatus = "AntivirusNotActive"      // neither internal nor external antivirus is active
	AntivirusInternalFailed AntivirusStatus = "AntivirusInternalFailed" // problem with internal intivirus
	AntivirusExternalFailed AntivirusStatus = "AntivirusExternalFailed" // problem with external intivirus
	AntivirusBothFailed     AntivirusStatus = "AntivirusBothFailed"     // both internal and external antivirus has failed
)

type AntivirusSetting struct {
	InternalEnabled    bool                  `json:"internalEnabled"`    // integrated antivirus is used?
	ExternalEnabled    bool                  `json:"externalEnabled"`    // an external antivirus is used? note: both internal and extenal can be used together
	Status             AntivirusStatus       `json:"status"`             // status of antivirus to be used for informative message
	ExternalList       ExternalAntivirusList `json:"externalList"`       // list of available antivirus plugins
	SelectedExternalId string                `json:"selectedExternalId"` // identifier of currently selected antivirus plugin
	Internal           InternalAntivirus     `json:"internal"`           // integrated engine settings
}

type AntivirusUpdatePhases string

const (
	AntivirusUpdateStarted        AntivirusUpdatePhases = "AntivirusUpdateStarted"        // "Update process started"
	AntivirusUpdateChecking       AntivirusUpdatePhases = "AntivirusUpdateChecking"       // "Checking for new version..."
	AntivirusUpdateDownload       AntivirusUpdatePhases = "AntivirusUpdateDownload"       // "Downloading new virus definition files..."
	AntivirusUpdateDownloadEngine AntivirusUpdatePhases = "AntivirusUpdateDownloadEngine" // "Downloading new engine..."
	AntivirusUpdateOk             AntivirusUpdatePhases = "AntivirusUpdateOk"             // Update finished, update not called yet
	AntivirusUpdateFailed         AntivirusUpdatePhases = "AntivirusUpdateFailed"         // "Update failed (see error log)"
)

type InternalUpdateStatus struct {
	Phase           AntivirusUpdatePhases `json:"phase"`           // state of update process
	Percentage      int                   `json:"percentage"`      // percent of progress, valid for: AntivirusUpdateChecking, AntivirusUpdateDownload, AntivirusUpdateDownloadEngine
	DatabaseAge     TimeSpan              `json:"databaseAge"`     // how old is virus database
	LastUpdateCheck TimeSpan              `json:"lastUpdateCheck"` // how long is since last database update check
	DatabaseVersion string                `json:"databaseVersion"` // virus database version
	EngineVersion   string                `json:"engineVersion"`   // scanning engine version
}

// ScannedProtocols - Product dependent part
type ScannedProtocols struct {
	Http bool `json:"http"`
	Ftp  bool `json:"ftp"`
	Smtp bool `json:"smtp"`
	Pop3 bool `json:"pop3"`
}

type ScanRuleType string

const (
	ScanRuleUrl       ScanRuleType = "ScanRuleUrl"
	ScanRuleMime      ScanRuleType = "ScanRuleMime"
	ScanRuleFilename  ScanRuleType = "ScanRuleFilename"
	ScanRuleFileGroup ScanRuleType = "ScanRuleFileGroup"
)

type ScanRuleConfig struct {
	Id          KId          `json:"id"`
	Enabled     bool         `json:"enabled"`
	Type        ScanRuleType `json:"type"`
	Pattern     string       `json:"pattern"`
	Scan        bool         `json:"scan"`
	Description string       `json:"description"`
}

type ScanRuleConfigList []ScanRuleConfig

type HttpFtpScanningConfig struct {
	MoveToQuarantine bool               `json:"moveToQuarantine"` // not available on Linux
	AlertClient      bool               `json:"alertClient"`
	AllowNotScanned  bool               `json:"allowNotScanned"`
	ScanRuleList     ScanRuleConfigList `json:"scanRuleList"`
}

type EmailScanningConfig struct {
	MoveToQuarantine bool           `json:"moveToQuarantine"` // not available on Linux
	PrependText      OptionalString `json:"prependText"`
	AllowTls         bool           `json:"allowTls"`
	AllowNotScanned  bool           `json:"allowNotScanned"`
}

type SslVpnScanningConfig struct {
	ScanUpload      bool `json:"scanUpload"`
	ScanDownload    bool `json:"scanDownload"`
	AllowNotScanned bool `json:"allowNotScanned"`
}

type AntivirusConfig struct {
	Antivirus       AntivirusSetting      `json:"antivirus"`
	Protocols       ScannedProtocols      `json:"protocols"`
	FileSizeLimit   OptionalLong          `json:"fileSizeLimit"`
	HttpFtpScanning HttpFtpScanningConfig `json:"httpFtpScanning"`
	EmailScanning   EmailScanningConfig   `json:"emailScanning"`
	SslVpnScanning  SslVpnScanningConfig  `json:"sslVpnScanning"` // not available on Linux
}

// AntivirusGet - Get Antivirus Settings
// Return
//  config - Antivirus Settings
func (s *ServerConnection) AntivirusGet() (*AntivirusConfig, error) {
	data, err := s.CallRaw("Antivirus.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config AntivirusConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// AntivirusSet - Set Antivirus
// Parameters
//	config - structure with complete antivirus settings
// Return
//	errors - list of errors \n
func (s *ServerConnection) AntivirusSet(config AntivirusConfig) (ErrorList, error) {
	params := struct {
		Config AntivirusConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("Antivirus.set", params)
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

// AntivirusUpdate - Force update of integrated antivirus
func (s *ServerConnection) AntivirusUpdate() error {
	_, err := s.CallRaw("Antivirus.update", nil)
	return err
}

// AntivirusGetUpdateStatus - Get progress of antivirus updating
// Return
//  status - progress of antivirus updating
func (s *ServerConnection) AntivirusGetUpdateStatus() (*InternalUpdateStatus, error) {
	data, err := s.CallRaw("Antivirus.getUpdateStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status InternalUpdateStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}
