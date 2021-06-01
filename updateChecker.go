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

// UpdateCheckerGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
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

// UpdateCheckerSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - Contains Structure with update checker's settings to be stored &a pplied.
func (s *ServerConnection) UpdateCheckerSet(config UpdateCheckerConfig) error {
	params := struct {
		Config UpdateCheckerConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("UpdateChecker.set", params)
	return err
}

// UpdateCheckerCheck - 8000 Internal error.  - "Internal error."
func (s *ServerConnection) UpdateCheckerCheck(checkVersionType CheckVersionType) error {
	params := struct {
		CheckVersionType CheckVersionType `json:"checkVersionType"`
	}{checkVersionType}
	_, err := s.CallRaw("UpdateChecker.check", params)
	return err
}

// UpdateCheckerGetStatus - 8000 Internal error. - "Internal error."
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

// UpdateCheckerGetProgressStatus - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	percentage - Returns percentage progress for Status Downloading \n
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

// UpdateCheckerDownload - 8000 Internal error. - "Internal error."
func (s *ServerConnection) UpdateCheckerDownload(checkVersionType CheckVersionType) error {
	params := struct {
		CheckVersionType CheckVersionType `json:"checkVersionType"`
	}{checkVersionType}
	_, err := s.CallRaw("UpdateChecker.download", params)
	return err
}

// UpdateCheckerUploadImage - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
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

// UpdateCheckerPerformCustomUpgrade - 8000 Internal error. - "Internal error."
func (s *ServerConnection) UpdateCheckerPerformCustomUpgrade(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("UpdateChecker.performCustomUpgrade", params)
	return err
}

// UpdateCheckerPerformUpgrade - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) UpdateCheckerPerformUpgrade() error {
	_, err := s.CallRaw("UpdateChecker.performUpgrade", nil)
	return err
}

// UpdateCheckerCancelDownload - 8000 Internal error. - "Internal error."
func (s *ServerConnection) UpdateCheckerCancelDownload() error {
	_, err := s.CallRaw("UpdateChecker.cancelDownload", nil)
	return err
}
