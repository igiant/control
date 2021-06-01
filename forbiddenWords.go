package control

import "encoding/json"

// ForbiddenWord - Forbidden words
type ForbiddenWord struct {
	Status      StoreStatus `json:"status"`
	Id          KId         `json:"id"`
	GroupName   string      `json:"groupName"` // don't use id here, it's not referenced entry
	GroupId     KId         `json:"groupId"`
	Enabled     bool        `json:"enabled"`
	Weight      int         `json:"weight"`
	Keyword     string      `json:"keyword"`
	Description string      `json:"description"`
}

type ForbiddenWordGroup struct {
	Name string `json:"name"`
	Id   KId    `json:"id"`
}

type ForbiddenWordList []ForbiddenWord

type ForbiddenWordGroupList []ForbiddenWordGroup

type ForbiddenWordsConfig struct {
	Limit   int  `json:"limit"`
	Enabled bool `json:"enabled"`
}

// ForbiddenWordsGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	query - conditions and limits. Included from weblib.
// Return
//	list - list of words and it's details
//	totalItems - count of all words on server (before the start/limit applied)
func (s *ServerConnection) ForbiddenWordsGet(query SearchQuery) (ForbiddenWordList, int, error) {
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("ForbiddenWords.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       ForbiddenWordList `json:"list"`
			TotalItems int               `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// ForbiddenWordsCreate - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	items - details for new words. field id is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - list of errors \n
//	result - list of IDs assigned to each item
func (s *ServerConnection) ForbiddenWordsCreate(items ForbiddenWordList) (ErrorList, CreateResultList, error) {
	params := struct {
		Items ForbiddenWordList `json:"items"`
	}{items}
	data, err := s.CallRaw("ForbiddenWords.create", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// ForbiddenWordsSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	ids - ids of words to be updated.
//	details - details for update. Field "kerio::web::KId" is ignored. All other fields must be filled and they are written to all words specified by ids.
// Return
//	errors - list of errors \n
func (s *ServerConnection) ForbiddenWordsSet(ids StringList, details ForbiddenWord) (ErrorList, error) {
	params := struct {
		Ids     StringList    `json:"ids"`
		Details ForbiddenWord `json:"details"`
	}{ids, details}
	data, err := s.CallRaw("ForbiddenWords.set", params)
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

// ForbiddenWordsRemove - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	ids - ids of words that should be removed
// Return
//	errors - list of errors
func (s *ServerConnection) ForbiddenWordsRemove(ids StringList) (ErrorList, error) {
	params := struct {
		Ids StringList `json:"ids"`
	}{ids}
	data, err := s.CallRaw("ForbiddenWords.remove", params)
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

// ForbiddenWordsApply - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	errors - list of errors \n
func (s *ServerConnection) ForbiddenWordsApply() (ErrorList, error) {
	data, err := s.CallRaw("ForbiddenWords.apply", nil)
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

// ForbiddenWordsReset - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ForbiddenWordsReset() error {
	_, err := s.CallRaw("ForbiddenWords.reset", nil)
	return err
}

// ForbiddenWordsGetConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Return
//	config - Complete configuration of Forbidden words module
func (s *ServerConnection) ForbiddenWordsGetConfig() (*ForbiddenWordsConfig, error) {
	data, err := s.CallRaw("ForbiddenWords.getConfig", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ForbiddenWordsConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ForbiddenWordsSetConfig - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
// Parameters
//	config - Complete configuration of Forbidden words module
func (s *ServerConnection) ForbiddenWordsSetConfig(config ForbiddenWordsConfig) error {
	params := struct {
		Config ForbiddenWordsConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("ForbiddenWords.setConfig", params)
	return err
}
