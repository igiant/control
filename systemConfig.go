package control

import "encoding/json"

type TimeZoneConfig struct {
	Id            KId    `json:"id"`
	Name          string `json:"name"`
	CurrentOffset int    `json:"currentOffset"`
	WinterOffset  int    `json:"winterOffset"`
	SummerOffset  int    `json:"summerOffset"`
}

type TimeZoneConfigList []TimeZoneConfig

type SystemConfiguration struct {
	Hostname   string         `json:"hostname"`
	NtpServer  OptionalString `json:"ntpServer"`
	TimeZoneId KId            `json:"timeZoneId"`
}

type NtpUpdatePhase string

const (
	NtpUpdateDisabled NtpUpdatePhase = "NtpUpdateDisabled"
	NtpUpdateOk       NtpUpdatePhase = "NtpUpdateOk"
	NtpUpdateError    NtpUpdatePhase = "NtpUpdateError"
	NtpUpdateProgress NtpUpdatePhase = "NtpUpdateProgress"
)

type NtpUpdateStatus struct {
	Phase        NtpUpdatePhase `json:"phase"`
	ErrorMessage string         `json:"errorMessage"`
}

// SystemConfigGet - Returns actual values for System configuration in WebAdmin
// Return
//	config - actual values for System configuration in WebAdmin
func (s *ServerConnection) SystemConfigGet() (*SystemConfiguration, error) {
	data, err := s.CallRaw("SystemConfig.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config SystemConfiguration `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// SystemConfigSet - Stores System configuration
// Parameters
//	config - contains system setting to be stored.
func (s *ServerConnection) SystemConfigSet(config SystemConfiguration) error {
	params := struct {
		Config SystemConfiguration `json:"config"`
	}{config}
	_, err := s.CallRaw("SystemConfig.set", params)
	return err
}

// SystemConfigGetTimeZones - Returns the list of known timezones.
// Parameters
//	currentDate - Client actual time to serve as an input for timezone and DST detection.
// Return
//	timeZones - list of known timezones.
func (s *ServerConnection) SystemConfigGetTimeZones(currentDate Date) (TimeZoneConfigList, error) {
	params := struct {
		CurrentDate Date `json:"currentDate"`
	}{currentDate}
	data, err := s.CallRaw("SystemConfig.getTimeZones", params)
	if err != nil {
		return nil, err
	}
	timeZones := struct {
		Result struct {
			TimeZones TimeZoneConfigList `json:"timeZones"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &timeZones)
	return timeZones.Result.TimeZones, err
}

// SystemConfigGetDateTime - Returns Date and Time for System configuration.
// Return
//	config - Returns Date and Time for System configuration.
func (s *ServerConnection) SystemConfigGetDateTime() (*DateTimeConfig, error) {
	data, err := s.CallRaw("SystemConfig.getDateTime", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config DateTimeConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// SystemConfigSetDateTime - Stores Date and Time for System configuration.
// Parameters
//	config - structure of system date and time settings
func (s *ServerConnection) SystemConfigSetDateTime(config DateTimeConfig) error {
	params := struct {
		Config DateTimeConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("SystemConfig.setDateTime", params)
	return err
}

// SystemConfigSetTimeFromNtp - Starts NTP client based on configured values.
func (s *ServerConnection) SystemConfigSetTimeFromNtp() error {
	_, err := s.CallRaw("SystemConfig.setTimeFromNtp", nil)
	return err
}

// SystemConfigGetNtpStatus - Returns Status of NTP client process.
// Return
//	status - Status of NTP client process.
func (s *ServerConnection) SystemConfigGetNtpStatus() (*NtpUpdateStatus, error) {
	data, err := s.CallRaw("SystemConfig.getNtpStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status NtpUpdateStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}
