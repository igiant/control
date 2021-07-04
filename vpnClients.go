package control

import "encoding/json"

type VpnClientState string

const (
	VpnClientConnecting     VpnClientState = "VpnClientConnecting"
	VpnClientAuthenticating VpnClientState = "VpnClientAuthenticating"
	VpnClientAuthenticated  VpnClientState = "VpnClientAuthenticated"
	VpnClientConnected      VpnClientState = "VpnClientConnected"
	VpnClientOther          VpnClientState = "VpnClientOther"
)

type OsCodeType string

const (
	OsWindows OsCodeType = "OsWindows"
	OsLinux   OsCodeType = "OsLinux"
	OsMacos   OsCodeType = "OsMacos"
	OsUnknown OsCodeType = "OsUnknown"
)

type VpnClientInfo struct {
	Id        KId            `json:"id"`
	Type      VpnType        `json:"type"`
	UserName  string         `json:"userName"`
	HostName  string         `json:"hostName"`
	Ip        IpAddress      `json:"ip"`
	ClientIp  IpAddress      `json:"clientIp"`
	State     VpnClientState `json:"state"`
	LoginTime int            `json:"loginTime"`
	Version   string         `json:"version"`
	OsCode    OsCodeType     `json:"osCode"`
	OsName    string         `json:"osName"`
}

type VpnClientList []VpnClientInfo

// VpnClientsGet - Returns VPN Clients data
// Parameters
//	query - filter/sort query
//	refresh - true in case, that data snapshot have to be refreshed
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) VpnClientsGet(query SearchQuery, refresh bool) (VpnClientList, int, error) {
	params := struct {
		Query   SearchQuery `json:"query"`
		Refresh bool        `json:"refresh"`
	}{query, refresh}
	data, err := s.CallRaw("VpnClients.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       VpnClientList `json:"list"`
			TotalItems int           `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// VpnClientsKill - Disconnects clients specified in ids list
// Parameters
//  ids - IDs list
func (s *ServerConnection) VpnClientsKill(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("VpnClients.kill", params)
	return err
}
