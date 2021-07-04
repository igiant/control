package control

import "encoding/json"

type UpdateCheckerConfig struct {
	Enabled             bool                `json:"enabled"`
	BetaVersion         bool                `json:"betaVersion"`
	Download            bool                `json:"download"`
	AutoUpdateTimeRange OptionalIdReference `json:"autoUpdateTimeRange"`
}

type CheckVersionType string

const (
	CheckVersionConfig CheckVersionType = "CheckVersionConfig"
	CheckVersionBeta   CheckVersionType = "CheckVersionBeta"
	CheckVersionFinal  CheckVersionType = "CheckVersionFinal"
)

type UpdateStatus string

const (
	UpdateStatusOk          UpdateStatus = "UpdateStatusOk"
	UpdateStatusChecking    UpdateStatus = "UpdateStatusChecking"
	UpdateStatusCheckFailed UpdateStatus = "UpdateStatusCheckFailed"
	/* States only for Linux version */
	UpdateStatusDownloadOk     UpdateStatus = "UpdateStatusDownloadOk"
	UpdateStatusDownloading    UpdateStatus = "UpdateStatusDownloading"
	UpdateStatusDownloadFailed UpdateStatus = "UpdateStatusDownloadFailed"
	UpdateStatusUpgrading      UpdateStatus = "UpdateStatusUpgrading"
	UpdateStatusUpgradeFailed  UpdateStatus = "UpdateStatusUpgradeFailed"
)

type UpdateCheckerInfo struct {
	Status             UpdateStatus `json:"status"`
	NewVersion         bool         `json:"newVersion"`
	LastCheckTime      int          `json:"lastCheckTime"`
	LastUpdateCheck    TimeSpan     `json:"lastUpdateCheck"`
	PackageCode        string       `json:"packageCode"`
	Description        string       `json:"description"`
	DownloadUrl        string       `json:"downloadUrl"`
	InfoUrl            string       `json:"infoUrl"`
	UpdateErrorDescr   string       `json:"updateErrorDescr"`
	AutoUpdatePlanned  bool         `json:"autoUpdatePlanned"`
	AutoUpdateDateTime string       `json:"autoUpdateDateTime"`
}

// UpdateCheckerGet - Returns configuration
// Return
//	config - Contains Structure with update checker's settings.
func (s *ServerConnection) UpdateCheckerGet() (*UpdateCheckerConfig, error) {
	data, err := s.CallRaw("UpdateChecker.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config UpdateCheckerConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// UpdateCheckerSet - Stores configuration
//	config - Contains Structure with update checker's settings to be stored &a pplied.
func (s *ServerConnection) UpdateCheckerSet(config UpdateCheckerConfig) error {
	params := struct {
		Config UpdateCheckerConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("UpdateChecker.set", params)
	return err
}

// UpdateCheckerCheck - Checks for a new version
func (s *ServerConnection) UpdateCheckerCheck(checkVersionType CheckVersionType) error {
	params := struct {
		CheckVersionType CheckVersionType `json:"checkVersionType"`
	}{checkVersionType}
	_, err := s.CallRaw("UpdateChecker.check", params)
	return err
}

// UpdateCheckerGetStatus - Returns actual state of Update checker
// Return
//	status - a phase of update process.
func (s *ServerConnection) UpdateCheckerGetStatus() (*UpdateCheckerInfo, error) {
	data, err := s.CallRaw("UpdateChecker.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status UpdateCheckerInfo `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// UpdateCheckerGetProgressStatus - Returns percentage progress for Status Downloading
// METHODS ONLY FOR LINUX VERSION
// Return
//	percentage - Returns percentage progress for Status Downloading
func (s *ServerConnection) UpdateCheckerGetProgressStatus() (int, error) {
	data, err := s.CallRaw("UpdateChecker.getProgressStatus", nil)
	if err != nil {
		return 0, err
	}
	percentage := struct {
		Result struct {
			Percentage int `json:"percentage"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &percentage)
	return percentage.Result.Percentage, err
}

// UpdateCheckerDownload - Starts Downloading
// METHODS ONLY FOR LINUX VERSION
//  checkVersionType -
func (s *ServerConnection) UpdateCheckerDownload(checkVersionType CheckVersionType) error {
	params := struct {
		CheckVersionType CheckVersionType `json:"checkVersionType"`
	}{checkVersionType}
	_, err := s.CallRaw("UpdateChecker.download", params)
	return err
}

// UpdateCheckerUploadImage - Converts fileId to id, that will be passed into performCustomUpgrade.
// METHODS ONLY FOR LINUX VERSION
//	fileId - according to spec 390.
// Return
//	id - an id obtained from fileId (same values);
func (s *ServerConnection) UpdateCheckerUploadImage(fileId string) (*KId, error) {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	data, err := s.CallRaw("UpdateChecker.uploadImage", params)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// UpdateCheckerPerformCustomUpgrade - Processes newly uploaded image for upgrade. Ids according to spec 390.
// In case of return true, it reboots machine (after 2s delay)
// METHODS ONLY FOR LINUX VERSION
func (s *ServerConnection) UpdateCheckerPerformCustomUpgrade(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("UpdateChecker.performCustomUpgrade", params)
	return err
}

// UpdateCheckerPerformUpgrade - Runs upgrade
// In case of return true, it reboots machine (after 2s delay)
// METHODS ONLY FOR LINUX VERSION
func (s *ServerConnection) UpdateCheckerPerformUpgrade() error {
	_, err := s.CallRaw("UpdateChecker.performUpgrade", nil)
	return err
}

// UpdateCheckerCancelDownload - Stops Downloading
// METHODS ONLY FOR LINUX VERSION
func (s *ServerConnection) UpdateCheckerCancelDownload() error {
	_, err := s.CallRaw("UpdateChecker.cancelDownload", nil)
	return err
}
