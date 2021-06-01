package control

// Certificates (Secured Web Interface and SSL-VPN) will be handled by CertificateManager
// SslVpnConfig - not applicable on Linux
type SslVpnConfig struct {
	Enabled     bool        `json:"enabled"`
	Port        int         `json:"port"`
	Certificate IdReference `json:"certificate"`
}

type CustomizedBrand struct {
	Enabled   bool   `json:"enabled"`
	PageTitle string `json:"pageTitle"`
}

type WebInterfaceConfig struct {
	UseSsl           bool            `json:"useSsl"`
	SslConfig        SslVpnConfig    `json:"sslConfig"`
	Hostname         OptionalString  `json:"hostname"`
	DetectedHostname string          `json:"detectedHostname"`
	AdminPath        string          `json:"adminPath"`
	Port             int             `json:"port"`
	SslPort          int             `json:"sslPort"`
	Certificate      IdReference     `json:"certificate"`
	CustomizedBrand  CustomizedBrand `json:"customizedBrand"`
}

// WebInterfaceGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation." \n
func (s *ServerConnection) WebInterfaceGet() (*WebInterfaceConfig, error) {
	data, err := s.CallRaw("WebInterface.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config WebInterfaceConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// WebInterfaceSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation." \n
// Parameters
//	config - structure with settings for webinterface module
//	revertTimeout - If client doesn't confirm config to this timeout, configuration is reverted (0 - revert disabled)
func (s *ServerConnection) WebInterfaceSet(config WebInterfaceConfig, revertTimeout int) error {
	params := struct {
		Config        WebInterfaceConfig `json:"config"`
		RevertTimeout int                `json:"revertTimeout"`
	}{config, revertTimeout}
	_, err := s.CallRaw("WebInterface.set", params)
	return err
}

// WebInterfaceUploadImage - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	fileId - according to spec 390.
//	isFavicon - true = the image is favicon, false = the image is product logo
func (s *ServerConnection) WebInterfaceUploadImage(fileId string, isFavicon bool) error {
	params := struct {
		FileId    string `json:"fileId"`
		IsFavicon bool   `json:"isFavicon"`
	}{fileId, isFavicon}
	_, err := s.CallRaw("WebInterface.uploadImage", params)
	return err
}

// WebInterfaceReset - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) WebInterfaceReset() error {
	_, err := s.CallRaw("WebInterface.reset", nil)
	return err
}
