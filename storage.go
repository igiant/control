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

// StorageGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	data - Returns list of data present on storage.
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

// StorageRemove - 1000 Operation failed. - "Some files cannot be deleted, they may be currently in use."
func (s *ServerConnection) StorageRemove(storageDataType StorageDataType) error {
	params := struct {
		StorageDataType StorageDataType `json:"storageDataType"`
	}{storageDataType}
	_, err := s.CallRaw("Storage.remove", params)
	return err
}
