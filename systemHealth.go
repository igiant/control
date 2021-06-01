package control

import "encoding/json"

type HistogramType string

const (
	HistogramOneDay   HistogramType = "HistogramOneDay"   // Data interval: 5min,  Max samples: 288
	HistogramTwoHours HistogramType = "HistogramTwoHours" // Data interval: 20sec, Max samples: 360
	HistogramOneWeek  HistogramType = "HistogramOneWeek"  // Data interval: 30min, Max samples: 336
	HistogramOneMonth HistogramType = "HistogramOneMonth" // Data interval: 2h,   Max samples: 372
)

type HistogramIntervalType string

const (
	HistogramInterval5m  HistogramIntervalType = "HistogramInterval5m"  // Data interval: 5min,  Max samples: 288, Length 1day
	HistogramInterval20s HistogramIntervalType = "HistogramInterval20s" // Data interval: 20sec, Max samples: 360, Length 2Hours
	HistogramInterval30m HistogramIntervalType = "HistogramInterval30m" // Data interval: 30min, Max samples: 336, Length 1Week
	HistogramInterval2h  HistogramIntervalType = "HistogramInterval2h"  // Data interval: 2h,   Max samples: 372, Length 1Month
)

// PercentHistogram - 0-100%, sample count and rate depens on requested type and is the same as in ActiveHost histogram
type PercentHistogram []float64

type SystemHealthData struct {
	Cpu         PercentHistogram `json:"cpu"`
	Memory      PercentHistogram `json:"memory"`
	MemoryTotal float64          `json:"memoryTotal"` // memory histogram has to have fixed maximum to this value
	DiskTotal   float64          `json:"diskTotal"`   // total number of bytes on data partition (install dir on windows)
	DiskFree    float64          `json:"diskFree"`
}

// SystemHealthGet -
func (s *ServerConnection) SystemHealthGet(histogramType HistogramType) (*SystemHealthData, error) {
	params := struct {
		HistogramType HistogramType `json:"histogramType"`
	}{histogramType}
	data, err := s.CallRaw("SystemHealth.get", params)
	if err != nil {
		return nil, err
	}
	systemHealthData := struct {
		Result struct {
			Data SystemHealthData `json:"data"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &systemHealthData)
	return &systemHealthData.Result.Data, err
}

// SystemHealthGetInc -
func (s *ServerConnection) SystemHealthGetInc(histogramIntervalType HistogramIntervalType, startSampleTime DateTimeStamp) (*SystemHealthData, *DateTimeStamp, error) {
	params := struct {
		HistogramIntervalType HistogramIntervalType `json:"histogramIntervalType"`
		StartSampleTime       DateTimeStamp         `json:"startSampleTime"`
	}{histogramIntervalType, startSampleTime}
	data, err := s.CallRaw("SystemHealth.getInc", params)
	if err != nil {
		return nil, nil, err
	}
	systemHealthData := struct {
		Result struct {
			Data       SystemHealthData `json:"data"`
			SampleTime DateTimeStamp    `json:"sampleTime"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &systemHealthData)
	return &systemHealthData.Result.Data, &systemHealthData.Result.SampleTime, err
}
