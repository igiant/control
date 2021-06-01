package control

// SystemTasksGetSsh - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	running - SSH is running
func (s *ServerConnection) SystemTasksGetSsh() (bool, error) {
	data, err := s.CallRaw("SystemTasks.getSsh", nil)
	if err != nil {
		return false, err
	}
	running := struct {
		Result struct {
			Running bool `json:"running"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &running)
	return running.Result.Running, err
}

// SystemTasksSetSsh - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	enable - true - enable SSH, false - disable SSH
func (s *ServerConnection) SystemTasksSetSsh(enable bool) error {
	params := struct {
		Enable bool `json:"enable"`
	}{enable}
	_, err := s.CallRaw("SystemTasks.setSsh", params)
	return err
}

// SystemTasksReboot - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SystemTasksReboot() error {
	_, err := s.CallRaw("SystemTasks.reboot", nil)
	return err
}

// SystemTasksShutdown - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) SystemTasksShutdown() error {
	_, err := s.CallRaw("SystemTasks.shutdown", nil)
	return err
}
