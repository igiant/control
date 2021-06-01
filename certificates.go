package control

// CertificatesGenerateEx - Invalid params. - "Unable to generate certificate, properties are invalid."
// Parameters
//	subject - properties specified by user
//	name - name of the new certificate
//	period - time properties specified by user, not relevant for CertificateRequest
//	subjectAlternativeNameList - Lists of subject alternative names in certificate. Key is similar to openSSL subj. alt. name type (see http://www.openssl.org/docs/apps/x509v3_config.html)
// Return
//	id - ID of generated certificate
func (s *ServerConnection) CertificatesGenerateEx(subject NamedValueList, name string, certificateType CertificateType, period ValidPeriod, subjectAlternativeNameList NamedMultiValueList) (*KId, error) {
	params := struct {
		Subject                    NamedValueList      `json:"subject"`
		Name                       string              `json:"name"`
		CertificateType            CertificateType     `json:"certificateType"`
		Period                     ValidPeriod         `json:"period"`
		SubjectAlternativeNameList NamedMultiValueList `json:"subjectAlternativeNameList"`
	}{subject, name, certificateType, period, subjectAlternativeNameList}
	data, err := s.CallRaw("Certificates.generateEx", params)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// CertificatesDetect - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	host - the host certificate of which will be detected
// Return
//	certificate - detected properties
func (s *ServerConnection) CertificatesDetect(host string) (*Certificate, error) {
	params := struct {
		Host string `json:"host"`
	}{host}
	data, err := s.CallRaw("Certificates.detect", params)
	if err != nil {
		return nil, err
	}
	certificate := struct {
		Result struct {
			Certificate Certificate `json:"certificate"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &certificate)
	return &certificate.Result.Certificate, err
}

// CertificatesApply - 8002 Database error. - "Unable to delete certificate."
// Return
//	errors - list of errors \n
func (s *ServerConnection) CertificatesApply() (ErrorList, error) {
	data, err := s.CallRaw("Certificates.apply", nil)
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

// CertificatesReset - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) CertificatesReset() error {
	_, err := s.CallRaw("Certificates.reset", nil)
	return err
}

// CertificatesImportCertificateP12 - Invalid params. - "Unable to import certificate, the content is invalid!"
// Parameters
//	fileId - id of uploaded file
//	name - name of the new certificate
//	password - password needed to decode certificate
// Return
//	id - ID of generated certificate
func (s *ServerConnection) CertificatesImportCertificateP12(fileId string, name string, certificateType CertificateType, password string) (*KId, error) {
	params := struct {
		FileId          string          `json:"fileId"`
		Name            string          `json:"name"`
		CertificateType CertificateType `json:"certificateType"`
		Password        string          `json:"password"`
	}{fileId, name, certificateType, password}
	data, err := s.CallRaw("Certificates.importCertificateP12", params)
	if err != nil {
		return nil, err
	}
	id := struct {
		Result struct {
			Id KId `json:"id"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &id)
	return &id.Result.Id, err
}

// CertificatesExportCertificateP12 - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	id - ID of the certificate or certificate request
//	password - password, which will be used to encrypt output certificate
//	includeCa - if true, engine will include whole certificate chain up to highest CA (only if all parents are present)
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) CertificatesExportCertificateP12(id KId, password string, includeCa bool) (*Download, error) {
	params := struct {
		Id        KId    `json:"id"`
		Password  string `json:"password"`
		IncludeCa bool   `json:"includeCa"`
	}{id, password, includeCa}
	data, err := s.CallRaw("Certificates.exportCertificateP12", params)
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

// CertificatesImportCertificateUrl - Invalid params. - "Unable to import certificate, the content is invalid!"
// Parameters
//	url - url, where will be certificate downloaded from
func (s *ServerConnection) CertificatesImportCertificateUrl(url string) error {
	params := struct {
		Url string `json:"url"`
	}{url}
	_, err := s.CallRaw("Certificates.importCertificateUrl", params)
	return err
}

// CertificatesSetDistrusted - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	ids - list of identifiers of deleted user templates
// Return
//	errors - error message list
func (s *ServerConnection) CertificatesSetDistrusted(ids KIdList) (ErrorList, error) {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("Certificates.setDistrusted", params)
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
