package control

import "encoding/json"

type CoreDump struct {
	Size      ByteValueWithUnits `json:"size"`
	Timestamp DateTimeStamp      `json:"timestamp"`
}

type CoreDumpWithImportance struct {
	Size       ByteValueWithUnits `json:"size"`
	Timestamp  DateTimeStamp      `json:"timestamp"`
	Importance Importance         `json:"importance"`
}

type DumpList []CoreDump

type DumpListWithImportance []CoreDumpWithImportance

// DumpsGet - Obtain list of available crash dumps
// Return
//	dumps - list of all available crash dumps
func (s *ServerConnection) DumpsGet() (DumpList, error) {
	data, err := s.CallRaw("Dumps.get", nil)
	if err != nil {
		return nil, err
	}
	dumps := struct {
		Result struct {
			Dumps DumpList `json:"dumps"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dumps)
	return dumps.Result.Dumps, err
}

// DumpsGetWithImportance - Obtain list of available crash dumps with highest importance
// Return
//	dumps - list of all available crash dumps with importance
func (s *ServerConnection) DumpsGetWithImportance() (*DumpListWithImportance, error) {
	data, err := s.CallRaw("Dumps.getWithImportance", nil)
	if err != nil {
		return nil, err
	}
	dumps := struct {
		Result struct {
			Dumps DumpListWithImportance `json:"dumps"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dumps)
	return &dumps.Result.Dumps, err
}

// DumpsRemove - Remove all crash dumps from server disk
func (s *ServerConnection) DumpsRemove() error {
	_, err := s.CallRaw("Dumps.remove", nil)
	return err
}

// DumpsSend - Upload last available crash dump to Kerio.
// Parameters
//	description - plain text information to be sent with crash dump
//	email - contact information to be sent with crash dump
func (s *ServerConnection) DumpsSend(description string, email string) error {
	params := struct {
		Description string `json:"description"`
		Email       string `json:"email"`
	}{description, email}
	_, err := s.CallRaw("Dumps.send", params)
	return err
}
