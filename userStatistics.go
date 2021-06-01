package control

type UserStatisticType string

const (
	UserStatisticAll   UserStatisticType = "UserStatisticAll"
	UserStatisticUser  UserStatisticType = "UserStatisticUser"
	UserStatisticOther UserStatisticType = "UserStatisticOther"
	UserStatisticGuest UserStatisticType = "UserStatisticGuest"
)

type UserStatistic struct {
	Id       KId               `json:"id"`
	UserName string            `json:"userName"`
	Type     UserStatisticType `json:"type"`
	FullName string            `json:"fullName"`
	Quota    int               `json:"quota"`
	Data     DataStatistic     `json:"data"`
}

type UserStatisticList []UserStatistic

// UserStatisticsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	query - filter/sort query
//	refresh - true in case, that data snapshot have to be refreshed
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) UserStatisticsGet(query SearchQuery, refresh bool) (UserStatisticList, int, error) {
	params := struct {
		Query   SearchQuery `json:"query"`
		Refresh bool        `json:"refresh"`
	}{query, refresh}
	data, err := s.CallRaw("UserStatistics.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       UserStatisticList `json:"list"`
			TotalItems int               `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// UserStatisticsRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	ids - list of user ids returned in user member by get
func (s *ServerConnection) UserStatisticsRemove(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("UserStatistics.remove", params)
	return err
}
