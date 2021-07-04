package control

import "encoding/json"

// UserInfo - A contact to user
type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProductInfo - Product identification
type ProductInfo struct {
	ProductVersion  string `json:"productVersion"`
	ProductName     string `json:"productName"`
	OperatingSystem string `json:"operatingSystem"`
	LicenseNumber   string `json:"licenseNumber"` // first 11 chars only
}

// AdditionalFiles - List of files attached to the ticket
type AdditionalFiles []string

type SystemInfo struct {
	Files       AdditionalFiles `json:"files"`
	Description string          `json:"description"`
}

// Interface for support incidents

// TechnicalSupportGetInfo - Get information from running product
// Return
//	userInfo (out UserInfo) User information
//	productInfo (out ProductInfo) Product information
//	systemInfo (out SystemInfo) System information
//	isUploadServerAvailable (out boolean) Is possible to upload attachment ?
func (s *ServerConnection) TechnicalSupportGetInfo() (*UserInfo, *ProductInfo, *SystemInfo, bool, error) {
	data, err := s.CallRaw("TechnicalSupport.getInfo", nil)
	if err != nil {
		return nil, nil, nil, false, err
	}
	userInfo := struct {
		Result struct {
			UserInfo                UserInfo    `json:"userInfo"`
			ProductInfo             ProductInfo `json:"productInfo"`
			SystemInfo              SystemInfo  `json:"systemInfo"`
			IsUploadServerAvailable bool        `json:"isUploadServerAvailable"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &userInfo)
	return &userInfo.Result.UserInfo, &userInfo.Result.ProductInfo, &userInfo.Result.SystemInfo, userInfo.Result.IsUploadServerAvailable, err
}

// TechnicalSupportAddSystemInfoToTicket - Add system info to the ticket
//	ticketId (in string) tickedId of target ticket
//	email (in string) email of the customer
func (s *ServerConnection) TechnicalSupportAddSystemInfoToTicket(ticketId string, email string) error {
	params := struct {
		TicketId string `json:"ticketId"`
		Email    string `json:"email"`
	}{ticketId, email}
	_, err := s.CallRaw("TechnicalSupport.addSystemInfoToTicket", params)
	return err
}
