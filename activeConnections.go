package control

import "encoding/json"

type ConnectionDirection string

const (
	ConnectionDirectionInbound  ConnectionDirection = "ConnectionDirectionInbound"
	ConnectionDirectionOutbound ConnectionDirection = "ConnectionDirectionOutbound"
	ConnectionDirectionLocal    ConnectionDirection = "ConnectionDirectionLocal"
)

type ConnectionPoint struct {
	Host string    `json:"host"`
	Ip   IpAddress `json:"ip"`
	Port int       `json:"port"`
}

type ActiveConnection struct {
	Id                KId                 `json:"id"`
	Src               ConnectionPoint     `json:"src"`
	Dst               ConnectionPoint     `json:"dst"`
	Protocol          string              `json:"protocol"`
	Timeout           int                 `json:"timeout"`
	Age               int                 `json:"age"`
	Rx                string              `json:"rx"`
	Tx                string              `json:"tx"`
	RxNum             float64             `json:"rxNum"`
	TxNum             float64             `json:"txNum"`
	Info              string              `json:"info"`
	Active            bool                `json:"active"`
	Direction         ConnectionDirection `json:"direction"`
	TrafficRule       string              `json:"trafficRule"`
	Service           string              `json:"service"`
	InternetLink      string              `json:"internetLink"`
	BandwidthRuleName string              `json:"bandwidthRuleName"`
}

type ActiveConnectionList []ActiveConnection

// ActiveConnectionsGet - Returns Active Connections data
//	query - filter/sort query
//	refresh - true in case, that data snapshot have to be refreshed
//	hostId - return data only for this host id
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) ActiveConnectionsGet(query SearchQuery, refresh bool, hostId KId) (ActiveConnectionList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query   SearchQuery `json:"query"`
		Refresh bool        `json:"refresh"`
		HostId  KId         `json:"hostId"`
	}{query, refresh, hostId}
	data, err := s.CallRaw("ActiveConnections.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ActiveConnectionList `json:"list"`
			TotalItems int                  `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ActiveConnectionsKill - Kills connections specified in ids list
//  ids - list of connections id
func (s *ServerConnection) ActiveConnectionsKill(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("ActiveConnections.kill", params)
	return err
}
