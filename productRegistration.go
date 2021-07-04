package control

import "encoding/json"

// @brief ineger value used in KISS as "UNLIMITED" in license, see Registration::subscribers and RegistrationFullStatus::users below
const unlimitedUsers int = -2

// RegDate - @brief a date structure
type RegDate struct {
	Year  int `json:"year"`  // year
	Month int `json:"month"` // 1-12
	Day   int `json:"day"`   // 1-31 max day is limited by month
}

// RegistrationNumber - @brief Registration Number info
type RegistrationNumber struct {
	Key         string `json:"key"`         // Registration number
	Type        string `json:"type"`        // A type of the reg. number (base id, addon...)
	Description string `json:"description"` // A description of the reg. number.
}

// RegStringList - @brief A list of strings
type RegStringList []string

// SurveyAnswer - @brief An answer for survey questions
type SurveyAnswer struct {
	QuestionID string `json:"questionID"` // ID of the question
	Answer     string `json:"answer"`     // answer to the question
}

// SurveyAnswerList - @brief a list of answers for survey questions
type SurveyAnswerList []SurveyAnswer

// RegistrationNumberList - @brief A list of registration numbers related to a registration.
type RegistrationNumberList []RegistrationNumber

// Extension - @brief Extension information
type Extension struct {
	Name string `json:"name"` // extension name
}

// ExtensionList - @brief A list of extensions related to a registration.
type ExtensionList []Extension

// LicenseDetail - @brief Details about the license's owner
type LicenseDetail struct {
	Organization string `json:"organization"` // compulsory
	Person       string `json:"person"`       // compulsory
	Email        string `json:"email"`        // compulsory
	Phone        string `json:"phone"`        // compulsory
	Web          string `json:"web"`
	Country      string `json:"country"` // compulsory
	State        string `json:"state"`   // compulsory for countries such as USA, Canada, Australia etc.
	City         string `json:"city"`    // compulsory
	Street       string `json:"street"`  // compulsory
	Zip          string `json:"zip"`     // compulsory
	Comment      string `json:"comment"`
}

// @brief The data related to a registration. Content of the structure is obtained
//  from our registration server by method getRegistrationInfo, modified by client and sent back
//  to the server by method finishRegistration

// Registration - @see getRegistrationInfo, finishRegistration
type Registration struct {
	Details          LicenseDetail          `json:"details"`          // Information about user
	ExpirationDate   RegDate                `json:"expirationDate"`   // Expiration date
	Subscribers      int                    `json:"subscribers"`      // A count of Subscribers (typically users) of the product
	ShowQuestions    bool                   `json:"showQuestions"`    // Have to show questions?
	RegistrationType int                    `json:"registrationType"` // is it edu/gov registration?
	EduInfo          string                 `json:"eduInfo"`          // information special for EDUcational type of organization
	RegNumbers       RegistrationNumberList `json:"regNumbers"`       // All registration numbers included in registration
	SurveyAnswers    SurveyAnswerList       `json:"surveyAnswers"`    // Survey answers - Answers should be sent in the same order as the questions are displayed.
	Extensions       ExtensionList          `json:"extensions"`       // list of extensions
}

// RegistrationFinishType - Type of registration finish
type RegistrationFinishType string

const (
	rfCreate   RegistrationFinishType = "rfCreate"   // Create a new registration
	rfModify   RegistrationFinishType = "rfModify"   // Modify existing Registration
	rfDownload RegistrationFinishType = "rfDownload" // Download license key without any modification
	rfStore    RegistrationFinishType = "rfStore"    // Just store in product without downloading key and modifying reg. (trial)
)

// RegistrationType - A type of the current registration of the product
type RegistrationType string

const (
	rsNoRegistration    RegistrationType = "rsNoRegistration"    // The product has not been registered yet
	rsTrialRegistered   RegistrationType = "rsTrialRegistered"   // The product has a valid trial registration
	rsTrialExpired      RegistrationType = "rsTrialExpired"      // The product has a trial registration but it has expired!
	rsProductRegistered RegistrationType = "rsProductRegistered" // The product has been registered.
)

// RegistrationStatus - A registration Status of the current product.
type RegistrationStatus struct {
	RegType RegistrationType `json:"regType"` // The registration type of the current prooduct
	Id      string           `json:"Id"`      // Base or trial ID used for registration
}

// ExpireType - Type of expiration
type ExpireType string

const (
	License      ExpireType = "License"      // License
	Subscription ExpireType = "Subscription" // Subscription
)

// ExpireInfo - Expire date information
type ExpireInfo struct {
	Type          ExpireType `json:"type"`          // type of expiration
	IsUnlimited   bool       `json:"isUnlimited"`   // is it a special license with expiration == never ?
	RemainingDays int        `json:"remainingDays"` // days remaining to subscription expiration
	Date          int        `json:"date"`          // last date of subscription
}

// LicenseExpireInfo - All expiration dates information in license
type LicenseExpireInfo []ExpireInfo

// RegistrationFullStatus - Full status information
type RegistrationFullStatus struct {
	RegType     RegistrationType  `json:"regType"`     // The registration type of the current prooduct
	Id          string            `json:"Id"`          // Base or trial ID used for registration
	Company     string            `json:"company"`     // Company name
	Users       int               `json:"users"`       // Users count
	Expirations LicenseExpireInfo `json:"expirations"` // sequence of expire information
}

// @brief Class provides an interface for both standard and trial registration
// process of Kerio products
// (Trial) registration begins with calling method startRegistration().
// The method returns a token (for identifying registration session) and
// a "security" picture.
// A code from the picture and the token are sent back to server as argument
// of method get() together with a registration number.
// If the reg. number already exists in the system the server sends back
// License details which could be displayied by the wizard.
// As step 3 the wizard offers a user interface for adding new registration
// numbers. A new number can be verified by method verifyRegistrationNumber().
// The registration can be finished by method finish().

// ProductRegistrationFinish - The Method finishes registration and installs the valid licenseKey.
//	token - ID of wizard's session
//	baseId - Base ID of registration
//	registrationInfo - Registration data retrieved from server by getRegistrationInfo() and modified by user.
//	finishType - how to finish the registration? Create a new one, modyfy an existing or just download an existing license?
func (s *ServerConnection) ProductRegistrationFinish(token string, baseId string, registrationInfo Registration, finishType RegistrationFinishType) error {
	params := struct {
		Token            string                 `json:"token"`
		BaseId           string                 `json:"baseId"`
		RegistrationInfo Registration           `json:"registrationInfo"`
		FinishType       RegistrationFinishType `json:"finishType"`
	}{token, baseId, registrationInfo, finishType}
	_, err := s.CallRaw("ProductRegistration.finish", params)
	return err
}

// ProductRegistrationGet - Retrieves existing registration data from the server.
//	token - ID of wizard's session
//	securityCode - a code number from the security immage
//	baseId - license ID
// Return
//	registrationInfo - the registration data related to the license ID
//	newRegistration - flag indicates whether the registration has already existed.
//	trial - trial ID registered on web, do not display registrationInfo and finish immediatelly
func (s *ServerConnection) ProductRegistrationGet(token string, securityCode string, baseId string) (*Registration, bool, bool, error) {
	params := struct {
		Token        string `json:"token"`
		SecurityCode string `json:"securityCode"`
		BaseId       string `json:"baseId"`
	}{token, securityCode, baseId}
	data, err := s.CallRaw("ProductRegistration.get", params)
	if err != nil {
		return nil, false, false, err
	}
	registrationInfo := struct {
		Result struct {
			RegistrationInfo Registration `json:"registrationInfo"`
			NewRegistration  bool         `json:"newRegistration"`
			Trial            bool         `json:"trial"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &registrationInfo)
	return &registrationInfo.Result.RegistrationInfo, registrationInfo.Result.NewRegistration, registrationInfo.Result.Trial, err
}

// ProductRegistrationGetFullStatus - @see RegistrationFullStatus
// Return
//	status - A current registration status of the product.
func (s *ServerConnection) ProductRegistrationGetFullStatus() (*RegistrationFullStatus, error) {
	data, err := s.CallRaw("ProductRegistration.getFullStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status RegistrationFullStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// ProductRegistrationGetStatus - @see RegistrationStatus
// Return
//	status - Current registration status of the product.
func (s *ServerConnection) ProductRegistrationGetStatus() (*RegistrationStatus, error) {
	data, err := s.CallRaw("ProductRegistration.getStatus", nil)
	if err != nil {
		return nil, err
	}
	status := struct {
		Result struct {
			Status RegistrationStatus `json:"status"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &status)
	return &status.Result.Status, err
}

// ProductRegistrationStart - Starts registration process. Methods connect to a server and obtain an identification token and a security image.
//	langId - language id
// Return
//	token - ID of wizard's session
//	image - URL of the image with the security code
//	showImage - show captcha image in wizard if true
func (s *ServerConnection) ProductRegistrationStart(langId string) (string, string, bool, error) {
	params := struct {
		LangId string `json:"langId"`
	}{langId}
	data, err := s.CallRaw("ProductRegistration.start", params)
	if err != nil {
		return "", "", false, err
	}
	token := struct {
		Result struct {
			Token     string `json:"token"`
			Image     string `json:"image"`
			ShowImage bool   `json:"showImage"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &token)
	return token.Result.Token, token.Result.Image, token.Result.ShowImage, err
}

// ProductRegistrationVerifyNumber - an uncomplete registration.
//	token - ID of wizard's session
//	baseId - Registration's baseId
//	regNumbersToVerify - a list of numbers to be verified
// Return
//	errors - description of an error in case of failure
//	regNumberInfo - information related to given registration key(s)
//	allowFinish - if false, the number is OK, but the registration cannot be finished without adding some other numbers.
//	users - the count of users connected to the license
//	expirationDate - licence expiration date
func (s *ServerConnection) ProductRegistrationVerifyNumber(token string, baseId string, regNumbersToVerify RegStringList) (ErrorList, RegistrationNumberList, bool, int, *RegDate, error) {
	params := struct {
		Token              string        `json:"token"`
		BaseId             string        `json:"baseId"`
		RegNumbersToVerify RegStringList `json:"regNumbersToVerify"`
	}{token, baseId, regNumbersToVerify}
	data, err := s.CallRaw("ProductRegistration.verifyNumber", params)
	if err != nil {
		return nil, nil, false, 0, nil, err
	}
	errors := struct {
		Result struct {
			Errors         ErrorList              `json:"errors"`
			RegNumberInfo  RegistrationNumberList `json:"regNumberInfo"`
			AllowFinish    bool                   `json:"allowFinish"`
			Users          int                    `json:"users"`
			ExpirationDate RegDate                `json:"expirationDate"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.RegNumberInfo, errors.Result.AllowFinish, errors.Result.Users, &errors.Result.ExpirationDate, err
}
