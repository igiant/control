package control

import "encoding/json"

// SystemTasksGetSsh - Returns SSH status
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

// SystemTasksSetSsh - Sets SSH setting
//	enable - true - enable SSH, false - disable SSH
func (s *ServerConnection) SystemTasksSetSsh(enable bool) error {
	params := struct {
		Enable bool `json:"enable"`
	}{enable}
	_, err := s.CallRaw("SystemTasks.setSsh", params)
	return err
}

// SystemTasksReboot - Performs system reboot
func (s *ServerConnection) SystemTasksReboot() error {
	_, err := s.CallRaw("SystemTasks.reboot", nil)
	return err
}

// SystemTasksShutdown - Performs system shutdown
func (s *ServerConnection) SystemTasksShutdown() error {
	_, err := s.CallRaw("SystemTasks.shutdown", nil)
	return err
}
