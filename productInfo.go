package control

import "encoding/json"

type WarningType string

const (
	WarnBetaVersion           WarningType = "WarnBetaVersion"
	WarnUpdateFailed          WarningType = "WarnUpdateFailed"
	WarnConfigurationReverted WarningType = "WarnConfigurationReverted"
)

type WarningTypeList []WarningType

type WarningInfo struct {
	Type         WarningType `json:"type"`
	Suppressable bool        `json:"suppressable"`
	Property     string      `json:"property"`
}

type WarningInfoList []WarningInfo

const tinyBox1 string = "HWANG100"

const smallBox1 string = "HWA1000"

const smallBox2 string = "HWA1120"

const smallBox3 string = "HWANG300"

const bigBox1 string = "HWA3000"

const bigBox2 string = "HWA3120"

const bigBox3 string = "HWA3130"

const bigBox4 string = "HWANG500"

type ProductInformation struct {
	VersionString        string `json:"versionString"`
	OsDescription        string `json:"osDescription"`
	FinalVersion         bool   `json:"finalVersion"`
	BoxEdition           string `json:"boxEdition"`
	BoxName              string `json:"boxName"`
	WifiAvailable        bool   `json:"wifiAvailable"`
	Ip6Available         bool   `json:"ip6Available"`
	LicenseSet           bool   `json:"licenseSet"`
	PasswordSet          bool   `json:"passwordSet"`
	ClientStatisticsSet  bool   `json:"clientStatisticsSet"`
	CentralManagementSet bool   `json:"centralManagementSet"`
}

type StatisticsData struct {
	ScreenWidth  int    `json:"screenWidth"`
	ScreenHeight int    `json:"screenHeight"`
	Localization string `json:"localization"`
	LoadTime     int    `json:"loadTime"`
	InitTime     int    `json:"initTime"`
}

type ApiEntity string

const (
	PolicyWizard ApiEntity = "PolicyWizard"
	AlertList    ApiEntity = "AlertList"
)

// ProductInfoGetAcknowledgmentsUrl - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	url - requested url
func (s *ServerConnection) ProductInfoGetAcknowledgmentsUrl() (string, error) {
	data, err := s.CallRaw("ProductInfo.getAcknowledgmentsUrl", nil)
	if err != nil {
		return "", err
	}
	url := struct {
		Result struct {
			Url string `json:"url"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &url)
	return url.Result.Url, err
}

// ProductInfoGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGet() (*ProductInformation, error) {
	data, err := s.CallRaw("ProductInfo.get", nil)
	if err != nil {
		return nil, err
	}
	productInfo := struct {
		Result struct {
			ProductInfo ProductInformation `json:"productInfo"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &productInfo)
	return &productInfo.Result.ProductInfo, err
}

// ProductInfoGetWarnings - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	warnings - list of warnings
func (s *ServerConnection) ProductInfoGetWarnings() (WarningInfoList, error) {
	data, err := s.CallRaw("ProductInfo.getWarnings", nil)
	if err != nil {
		return nil, err
	}
	warnings := struct {
		Result struct {
			Warnings WarningInfoList `json:"warnings"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &warnings)
	return warnings.Result.Warnings, err
}

// ProductInfoDisableWarning - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoDisableWarning(warningType WarningType) error {
	params := struct {
		WarningType WarningType `json:"warningType"`
	}{warningType}
	_, err := s.CallRaw("ProductInfo.disableWarning", params)
	return err
}

// ProductInfoGetSystemHostname - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGetSystemHostname() (string, error) {
	data, err := s.CallRaw("ProductInfo.getSystemHostname", nil)
	if err != nil {
		return "", err
	}
	hostname := struct {
		Result struct {
			Hostname string `json:"hostname"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hostname)
	return hostname.Result.Hostname, err
}

// ProductInfoConfigUpdate - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoConfigUpdate() error {
	_, err := s.CallRaw("ProductInfo.configUpdate", nil)
	return err
}

// ProductInfoUploadLicense - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoUploadLicense(fileId string) error {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	_, err := s.CallRaw("ProductInfo.uploadLicense", params)
	return err
}

// ProductInfoAcceptUnregisteredTrial - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoAcceptUnregisteredTrial() error {
	_, err := s.CallRaw("ProductInfo.acceptUnregisteredTrial", nil)
	return err
}

// ProductInfoGetSupportInfo - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGetSupportInfo() (*Download, error) {
	data, err := s.CallRaw("ProductInfo.getSupportInfo", nil)
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

// ProductInfoGetClientStatistics - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGetClientStatistics() (bool, error) {
	data, err := s.CallRaw("ProductInfo.getClientStatistics", nil)
	if err != nil {
		return false, err
	}
	setting := struct {
		Result struct {
			Setting bool `json:"setting"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &setting)
	return setting.Result.Setting, err
}

// ProductInfoSetClientStatistics - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoSetClientStatistics(setting bool) error {
	params := struct {
		Setting bool `json:"setting"`
	}{setting}
	_, err := s.CallRaw("ProductInfo.setClientStatistics", params)
	return err
}

// ProductInfoSetStatisticsData - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoSetStatisticsData(data StatisticsData) error {
	params := struct {
		Data StatisticsData `json:"data"`
	}{data}
	_, err := s.CallRaw("ProductInfo.setStatisticsData", params)
	return err
}

// ProductInfoGetUptime - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGetUptime() (int, error) {
	data, err := s.CallRaw("ProductInfo.getUptime", nil)
	if err != nil {
		return 0, err
	}
	uptime := struct {
		Result struct {
			Uptime int `json:"uptime"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &uptime)
	return uptime.Result.Uptime, err
}

// ProductInfoGetUsedDevicesCount - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ProductInfoGetUsedDevicesCount() (int, int, error) {
	data, err := s.CallRaw("ProductInfo.getUsedDevicesCount", nil)
	if err != nil {
		return 0, 0, err
	}
	devices := struct {
		Result struct {
			Devices  int `json:"devices"`
			Accounts int `json:"accounts"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &devices)
	return devices.Result.Devices, devices.Result.Accounts, err
}

// ProductInfoAccountUsage - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	apiEntity - which entity was used
func (s *ServerConnection) ProductInfoAccountUsage(apiEntity ApiEntity) error {
	params := struct {
		ApiEntity ApiEntity `json:"apiEntity"`
	}{apiEntity}
	_, err := s.CallRaw("ProductInfo.accountUsage", params)
	return err
}
