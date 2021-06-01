package control

import "encoding/json"

type RouterAdvertisement struct {
	Id          KId            `json:"id"`
	Enabled     bool           `json:"enabled"`
	InterfaceId IdReference    `json:"interfaceId"`
	Prefix      Ip6AddressMask `json:"prefix"`
}

type RouterAdvertisementList []RouterAdvertisement

type RouterAdvertisementsConfig struct {
	Enabled bool `json:"enabled"`
}

type RouterAdvertisementsModeType string

const (
	RouterAdvertisementsAutomatic RouterAdvertisementsModeType = "RouterAdvertisementsAutomatic"
	RouterAdvertisementsManual    RouterAdvertisementsModeType = "RouterAdvertisementsManual"
)

// RouterAdvertisementsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	list - list of entries
func (s *ServerConnection) RouterAdvertisementsGet() (RouterAdvertisementList, error) {
	data, err := s.CallRaw("RouterAdvertisements.get", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List RouterAdvertisementList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// RouterAdvertisementsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	advertisements - list of advertisment configurations (prefixes) to be stored and advertised
// Return
//	errors - list of errors \n
func (s *ServerConnection) RouterAdvertisementsSet(advertisements RouterAdvertisementList) (ErrorList, error) {
	params := struct {
		Advertisements RouterAdvertisementList `json:"advertisements"`
	}{advertisements}
	data, err := s.CallRaw("RouterAdvertisements.set", params)
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

// RouterAdvertisementsGetConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	config - configuration values
func (s *ServerConnection) RouterAdvertisementsGetConfig() (*RouterAdvertisementsConfig, error) {
	data, err := s.CallRaw("RouterAdvertisements.getConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config RouterAdvertisementsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// RouterAdvertisementsSetConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - configuration values
func (s *ServerConnection) RouterAdvertisementsSetConfig(config RouterAdvertisementsConfig) error {
	params := struct {
		Config RouterAdvertisementsConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("RouterAdvertisements.setConfig", params)
	return err
}

// RouterAdvertisementsGetMode - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	mode - result
func (s *ServerConnection) RouterAdvertisementsGetMode() (*RouterAdvertisementsModeType, error) {
	data, err := s.CallRaw("RouterAdvertisements.getMode", nil)
	if err != nil {
		return nil, err
	}
	mode := struct {
		Result struct {
			Mode RouterAdvertisementsModeType `json:"mode"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &mode)
	return &mode.Result.Mode, err
}

// RouterAdvertisementsSetMode - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	mode - new value
func (s *ServerConnection) RouterAdvertisementsSetMode(mode RouterAdvertisementsModeType) error {
	params := struct {
		Mode RouterAdvertisementsModeType `json:"mode"`
	}{mode}
	_, err := s.CallRaw("RouterAdvertisements.setMode", params)
	return err
}
