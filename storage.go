package control

import "encoding/json"

type StorageDataType string

const (
	StorageDataStar       StorageDataType = "StorageDataStar"
	StorageDataLogs       StorageDataType = "StorageDataLogs"
	StorageDataCrash      StorageDataType = "StorageDataCrash"
	StorageDataPktdump    StorageDataType = "StorageDataPktdump"
	StorageDataUpdate     StorageDataType = "StorageDataUpdate"
	StorageDataQuarantine StorageDataType = "StorageDataQuarantine"
	StorageDataHttpCache  StorageDataType = "StorageDataHttpCache"
)

type StorageData struct {
	Type StorageDataType `json:"type"`
	Size float64         `json:"size"`
}

type StorageDataList []StorageData

// StorageGet - Returns list of data present on storage.
// Return
//	data - list of data present on storage.
func (s *ServerConnection) StorageGet() (StorageDataList, error) {
	data, err := s.CallRaw("Storage.get", nil)
	if err != nil {
		return nil, err
	}
	dataList := struct {
		Result struct {
			Data StorageDataList `json:"data"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &dataList)
	return dataList.Result.Data, err
}

// StorageRemove - Delete data specified by type from storage.
// Parameters
//  storageDataType - data specified by type from storage.
func (s *ServerConnection) StorageRemove(storageDataType StorageDataType) error {
	params := struct {
		StorageDataType StorageDataType `json:"storageDataType"`
	}{storageDataType}
	_, err := s.CallRaw("Storage.remove", params)
	return err
}
