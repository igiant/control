package control

type TrafficStatisticsType string

const (
	TrafficStatisticsInterface     TrafficStatisticsType = "TrafficStatisticsInterface"
	TrafficStatisticsTrafficRule   TrafficStatisticsType = "TrafficStatisticsTrafficRule"
	TrafficStatisticsBandwidthRule TrafficStatisticsType = "TrafficStatisticsBandwidthRule"
)

type TrafficStatistic struct {
	Id            KId                   `json:"id"`
	Name          string                `json:"name"`
	Type          TrafficStatisticsType `json:"type"`
	ComponentId   KId                   `json:"componentId"`
	InterfaceType InterfaceType         `json:"interfaceType"`
	Data          DataStatistic         `json:"data"`
}

type TrafficStatisticList []TrafficStatistic

// TrafficStatisticsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	query - filter/sort query
//	refresh - true in case, that data snapshot have to be refreshed
// Return
//	list - output data
//	totalItems - all data count
func (s *ServerConnection) TrafficStatisticsGet(query SearchQuery, refresh bool) (TrafficStatisticList, int, error) {
	params := struct {
		Query   SearchQuery `json:"query"`
		Refresh bool        `json:"refresh"`
	}{query, refresh}
	data, err := s.CallRaw("TrafficStatistics.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       TrafficStatisticList `json:"list"`
			TotalItems int                  `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// TrafficStatisticsRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	ids - Ids of statistics items to be removed.
func (s *ServerConnection) TrafficStatisticsRemove(ids KIdList) error {
	params := struct {
		Ids KIdList `json:"ids"`
	}{ids}
	_, err := s.CallRaw("TrafficStatistics.remove", params)
	return err
}

// TrafficStatisticsGetHistogram - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	id - rule or interface id returned in item.id member by get
// Return
//	hist - output data
func (s *ServerConnection) TrafficStatisticsGetHistogram(histogramType HistogramType, id KId) (*Histogram, error) {
	params := struct {
		HistogramType HistogramType `json:"histogramType"`
		Id            KId           `json:"id"`
	}{histogramType, id}
	data, err := s.CallRaw("TrafficStatistics.getHistogram", params)
	if err != nil {
		return nil, err
	}
	hist := struct {
		Result struct {
			Hist Histogram `json:"hist"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hist)
	return &hist.Result.Hist, err
}

// TrafficStatisticsGetHistogramInc - 1004  Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	startSampleTime - Specifies starting time for returned d
//	id - rule or interface id returned in item.id member by get
// Return
//	hist - output data
//	sampleTime - Returns first time, that is not included in returned data (pass it to the same method again as a lastSampleTime to obtain only new data since last request)
func (s *ServerConnection) TrafficStatisticsGetHistogramInc(histogramIntervalType HistogramIntervalType, id KId, startSampleTime DateTimeStamp) (*Histogram, *DateTimeStamp, error) {
	params := struct {
		HistogramIntervalType HistogramIntervalType `json:"histogramIntervalType"`
		Id                    KId                   `json:"id"`
		StartSampleTime       DateTimeStamp         `json:"startSampleTime"`
	}{histogramIntervalType, id, startSampleTime}
	data, err := s.CallRaw("TrafficStatistics.getHistogramInc", params)
	if err != nil {
		return nil, nil, err
	}
	hist := struct {
		Result struct {
			Hist       Histogram     `json:"hist"`
			SampleTime DateTimeStamp `json:"sampleTime"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &hist)
	return &hist.Result.Hist, &hist.Result.SampleTime, err
}
