package control

// return information about serialnumber of hardware box

// HardwareInfoGetBoxSerialNumber -
func (s *ServerConnection) HardwareInfoGetBoxSerialNumber() (string, error) {
	data, err := s.CallRaw("HardwareInfo.getBoxSerialNumber", nil)
	if err != nil {
		return "", err
	}
	serialNumber := struct {
		Result struct {
			SerialNumber string `json:"serialNumber"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &serialNumber)
	return serialNumber.Result.SerialNumber, err
}
