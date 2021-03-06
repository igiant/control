package control

import "encoding/json"

type Target string

const (
	TargetSamepageIo Target = "TargetSamepageIo"
	TargetFtpServer  Target = "TargetFtpServer"
	TargetMyKerio    Target = "TargetMyKerio"
)

type ConfigurationBackupConfig struct {
	Enabled  bool   `json:"enabled"`
	Target   Target `json:"target"`
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

type ConfigurationBackupPhase string

const (
	ConfigurationBackupOk         ConfigurationBackupPhase = "ConfigurationBackupOk"
	ConfigurationBackupInProgress ConfigurationBackupPhase = "ConfigurationBackupInProgress"
	ConfigurationBackupFailed     ConfigurationBackupPhase = "ConfigurationBackupFailed"
)

type ConfigurationBackupStatus struct {
	Phase        ConfigurationBackupPhase `json:"phase"`
	LastBackup   TimeSpan                 `json:"lastBackup"`
	Url          string                   `json:"url"`
	ErrorMessage string                   `json:"errorMessage"`
}

// ConfigurationBackupGet - Returns configuration
// Return
//	config - Contains Structure with Configuration backup settings.
func (s *ServerConnection) ConfigurationBackupGet() (*ConfigurationBackupConfig, error) {
	data, err := s.CallRaw("ConfigurationBackup.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ConfigurationBackupConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ConfigurationBackupSet - Stores configuration
//	config - Contains Structure with Configuration backup settings.
func (s *ServerConnection) ConfigurationBackupSet(config ConfigurationBackupConfig) error {
	params := struct {
		Config ConfigurationBackupConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ConfigurationBackup.set", params)
	return err
}

// ConfigurationBackupBackupNow - Runs backup
func (s *ServerConnection) ConfigurationBackupBackupNow() error {
	_, err := s.CallRaw("ConfigurationBackup.backupNow", nil)
	return err
}

// ConfigurationBackupGetStatus - Returns actual state of Configuration backup status
// Return
//	status - a phase of update process.
func (s *ServerConnection) ConfigurationBackupGetStatus() (*ConfigurationBackupStatus, error) {
	data, err := s.CallRaw("ConfigurationBackup.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status ConfigurationBackupStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}
