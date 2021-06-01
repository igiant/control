package control

type PacketDumpStatus struct {
	SizeKb  int  `json:"sizeKb"`
	Running bool `json:"running"`
	Exists  bool `json:"exists"` // dump file exists on disk
}

// PacketDumpGetExpression - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpGetExpression() (string, error) {
	data, err := s.CallRaw("PacketDump.getExpression", nil)
	if err != nil {
		return "", err
	}
	expression := struct {
		Result struct {
			Expression string `json:"expression"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &expression)
	return expression.Result.Expression, err
}

// PacketDumpSetExpression - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpSetExpression(expression string) error {
	params := struct {
		Expression string `json:"expression"`
	}{expression}
	_, err := s.CallRaw("PacketDump.setExpression", params)
	return err
}

// PacketDumpStart - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpStart() error {
	_, err := s.CallRaw("PacketDump.start", nil)
	return err
}

// PacketDumpStop - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpStop() error {
	_, err := s.CallRaw("PacketDump.stop", nil)
	return err
}

// PacketDumpClear - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpClear() error {
	_, err := s.CallRaw("PacketDump.clear", nil)
	return err
}

// PacketDumpDownload - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpDownload() (*Download, error) {
	data, err := s.CallRaw("PacketDump.download", nil)
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

// PacketDumpGetStatus - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) PacketDumpGetStatus() (*PacketDumpStatus, error) {
	data, err := s.CallRaw("PacketDump.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status PacketDumpStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}
