package control

import "encoding/json"

type Day string

const (
	Monday    Day = "Monday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
	Thursday  Day = "Thursday"
	Friday    Day = "Friday"
	Saturday  Day = "Saturday"
	Sunday    Day = "Sunday"
)

type DayList []Day

// TimeRangesApply - Write changes cached in manager to configuration
// Return
//	errors - list of errors
func (s *ServerConnection) TimeRangesApply() (ErrorList, error) {
	data, err := s.CallRaw("TimeRanges.apply", nil)
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

// TimeRangesReset - Discard changes cached in manager
func (s *ServerConnection) TimeRangesReset() error {
	_, err := s.CallRaw("TimeRanges.reset", nil)
	return err
}
