package control

import "encoding/json"

// ValidType - Certificate Time properties info
type ValidType string

const (
	Valid       ValidType = "Valid"
	NotValidYet ValidType = "NotValidYet"
	ExpireSoon  ValidType = "ExpireSoon"
	Expired     ValidType = "Expired"
)

// ValidPeriod - Certificate Time properties
type ValidPeriod struct {
	ValidFromDate Date      `json:"validFromDate"` // @see SharedStructures.idl shared in lib
	ValidFromTime Time      `json:"validFromTime"` // @see SharedStructures.idl shared in lib
	ValidToDate   Date      `json:"validToDate"`   // @see SharedStructures.idl shared in lib
	ValidToTime   Time      `json:"validToTime"`   // @see SharedStructures.idl shared in lib
	ValidType     ValidType `json:"validType"`
}

type CertificateType string

const (
	ActiveCertificate   CertificateType = "ActiveCertificate"
	InactiveCertificate CertificateType = "InactiveCertificate"
	CertificateRequest  CertificateType = "CertificateRequest"
	Authority           CertificateType = "Authority"
	LocalAuthority      CertificateType = "LocalAuthority"
	BuiltInAuthority    CertificateType = "BuiltInAuthority"
	ServerCertificate   CertificateType = "ServerCertificate"
)

// CertificatesGenerateEx - Generate certificate.
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

// Certificate properties
// issuer & subject valid names:
//  hostname;        // max 127 bytes
//  organizationName;    // max 127 bytes
//  organizationalUnitName; // max 127 bytes
//  city;          // max 127 bytes
//  state;          // max 127 bytes
//  country;         // ISO 3166 code
// Certificate -  emailAddress;      // max 255 bytes
type Certificate struct {
	Id                         KId                 `json:"id"`
	Status                     StoreStatus         `json:"status"`
	Name                       string              `json:"name"`
	Issuer                     NamedValueList      `json:"issuer"`
	Subject                    NamedValueList      `json:"subject"`
	SubjectAlternativeNameList NamedMultiValueList `json:"subjectAlternativeNameList"`
	Fingerprint                string              `json:"fingerprint"`       // 128-bit MD5, i.e. 16 hexa values separated by colons
	FingerprintSha1            string              `json:"fingerprintSha1"`   // 160-bit SHA1, i.e. 20 hexa values separated by colons
	FingerprintSha256          string              `json:"fingerprintSha256"` // 512-bit SHA256, i.e. 64 hexa values separated by colons
	ValidPeriod                ValidPeriod         `json:"validPeriod"`
	Valid                      bool                `json:"valid"` // exists and valid content
	Type                       CertificateType     `json:"type"`
	IsUntrusted                bool                `json:"isUntrusted"`
	VerificationMessage        string              `json:"verificationMessage"`
	ChainInfo                  StringList          `json:"chainInfo"`
	IsSelfSigned               bool                `json:"isSelfSigned"`
}

type CertificateList []Certificate

// CertificatesDetect - Detect certificate of given VPN host.
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

// CertificatesApply - write changes cached in manager to configuration
// Return
//	errors - list of errors
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

// CertificatesReset - discard changes cached in manager
func (s *ServerConnection) CertificatesReset() error {
	_, err := s.CallRaw("Certificates.reset", nil)
	return err
}

// CertificatesImportCertificateP12 - Import certificate in PKCS #12 format
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

// CertificatesExportCertificateP12 - Export certificate in PKCS #12 format
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

// CertificatesImportCertificateUrl - Import certificate from url
//	url - url, where will be certificate downloaded from
func (s *ServerConnection) CertificatesImportCertificateUrl(url string) error {
	params := struct {
		Url string `json:"url"`
	}{url}
	_, err := s.CallRaw("Certificates.importCertificateUrl", params)
	return err
}

// CertificatesSetDistrusted - Distrust list of certificate records
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
