package control

import "encoding/json"

const UpdateTimeNever int = 0

type IntrusionPreventionAction string

const (
	IntrusionPreventionActionDropAndLog IntrusionPreventionAction = "IntrusionPreventionActionDropAndLog"
	IntrusionPreventionActionLog        IntrusionPreventionAction = "IntrusionPreventionActionLog"
	IntrusionPreventionActionNothing    IntrusionPreventionAction = "IntrusionPreventionActionNothing"
)

type BlackList struct {
	Id     KId                       `json:"id"`
	Name   string                    `json:"name"`
	Url    string                    `json:"url"`
	Action IntrusionPreventionAction `json:"action"`
}

type RuleReference struct {
	Id          string `json:"id"` // gid:sid format ("1:123456") - this is intentionally string, not kerio::web::KId
	Description string `json:"description"`
}

type IntrusionPreventionUpdatePhases string

const (
	IntrusionPreventionUpdateOk       IntrusionPreventionUpdatePhases = "IntrusionPreventionUpdateOk"
	IntrusionPreventionUpdateError    IntrusionPreventionUpdatePhases = "IntrusionPreventionUpdateError"
	IntrusionPreventionUpdateProgress IntrusionPreventionUpdatePhases = "IntrusionPreventionUpdateProgress"
)

type IntrusionPreventionInfo struct {
	LastUpdateCheck TimeSpan                        `json:"lastUpdateCheck"`
	DatabaseVersion string                          `json:"databaseVersion"`
	Phase           IntrusionPreventionUpdatePhases `json:"phase"`
	ErrorMessage    string                          `json:"errorMessage"`
}

type BlackListList []BlackList

type RuleReferenceList []RuleReference

type IntrusionPreventionConfig struct {
	Enabled             bool                      `json:"enabled"`
	High                IntrusionPreventionAction `json:"high"`
	Medium              IntrusionPreventionAction `json:"medium"`
	Low                 IntrusionPreventionAction `json:"low"`
	UpdateCheckInterval OptionalLong              `json:"updateCheckInterval"`
	BlackLists          BlackListList             `json:"blackLists"`
	Ports               NamedValueList            `json:"ports"`
}

// IntrusionPreventionGet - Returns configuration
// Return
//	config - complete configuration of Intrusion Prevention system
func (s *ServerConnection) IntrusionPreventionGet() (*IntrusionPreventionConfig, error) {
	data, err := s.CallRaw("IntrusionPrevention.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config IntrusionPreventionConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// IntrusionPreventionSet - Stores configuration
// Parameters
//	config - complete configuration of Intrusion Prevention system
// Return
//	errors - list of errors \n
func (s *ServerConnection) IntrusionPreventionSet(config IntrusionPreventionConfig) (ErrorList, error) {
	params := struct {
		Config IntrusionPreventionConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("IntrusionPrevention.set", params)
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

// IntrusionPreventionGetSignatureDescription - Returns signature description
func (s *ServerConnection) IntrusionPreventionGetSignatureDescription(id string) (string, error) {
	params := struct {
		Id string `json:"id"`
	}{id}
	data, err := s.CallRaw("IntrusionPrevention.getSignatureDescription", params)
	if err != nil {
		return "", err
	}
	description := struct {
		Result struct {
			Description string `json:"description"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &description)
	return description.Result.Description, err
}

// IntrusionPreventionGetIgnoredRules - Returns list of ignored signatures
// Return
//	ignored - List of rules that are excluded from usage in IPS
func (s *ServerConnection) IntrusionPreventionGetIgnoredRules() (RuleReferenceList, error) {
	data, err := s.CallRaw("IntrusionPrevention.getIgnoredRules", nil)
	if err != nil {
		return nil, err
	}
	ignored := struct {
		Result struct {
			Ignored RuleReferenceList `json:"ignored"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &ignored)
	return ignored.Result.Ignored, err
}

// IntrusionPreventionSetIgnoredRules - Sets list of ignored signatures
// Parameters
//	ignored - List of rules that are excluded from usage in IPS
// Return
//	errors - list of errors
func (s *ServerConnection) IntrusionPreventionSetIgnoredRules(ignored RuleReferenceList) (ErrorList, error) {
	params := struct {
		Ignored RuleReferenceList `json:"ignored"`
	}{ignored}
	data, err := s.CallRaw("IntrusionPrevention.setIgnoredRules", params)
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

// IntrusionPreventionUpdate - Checks new version of database and updates it
// Parameters
//  force - checks new version force if true
func (s *ServerConnection) IntrusionPreventionUpdate(force bool) error {
	params := struct {
		Force bool `json:"force"`
	}{force}
	_, err := s.CallRaw("IntrusionPrevention.update", params)
	return err
}

// IntrusionPreventionGetUpdateStatus - Returns actual state of Update
// Return
//  status - actual state of Update
func (s *ServerConnection) IntrusionPreventionGetUpdateStatus() (*IntrusionPreventionInfo, error) {
	data, err := s.CallRaw("IntrusionPrevention.getUpdateStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status IntrusionPreventionInfo `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}
