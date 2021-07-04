package control

import "encoding/json"

type HttpCacheStatus struct {
	Used  float64 `json:"used"` // in B
	Files int     `json:"files"`
	Hit   int     `json:"hit"`
	Miss  int     `json:"miss"`
}

type HttpCacheConfig struct {
	TransparentEnabled    bool            `json:"transparentEnabled"`
	NonTransparentEnabled bool            `json:"nonTransparentEnabled"`
	ReverseEnabled        bool            `json:"reverseEnabled"`
	CacheSize             int             `json:"cacheSize"` // in MB
	HttpTtl               int             `json:"httpTtl"`   // in days
	Status                HttpCacheStatus `json:"status"`    // read-only
}

type UrlSpecificTtl struct {
	Ttl         int    `json:"ttl"` // in hours
	Url         string `json:"url"`
	Description string `json:"description"`
}

type UrlSpecificTtlList []UrlSpecificTtl

// HttpCacheGet - Gets Http Cache configuration and status
// Return
//	config - current configuration
func (s *ServerConnection) HttpCacheGet() (*HttpCacheConfig, error) {
	data, err := s.CallRaw("HttpCache.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config HttpCacheConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// HttpCacheSet - Sets Http Cache configuration
// Parameters
//	config - new configuration
func (s *ServerConnection) HttpCacheSet(config HttpCacheConfig) error {
	params := struct {
		Config HttpCacheConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("HttpCache.set", params)
	return err
}

// HttpCacheGetUrlSpecificTtl - Gets list of URL specific TTL
// Return
//	list - URL specific TTL list
func (s *ServerConnection) HttpCacheGetUrlSpecificTtl() (UrlSpecificTtlList, error) {
	data, err := s.CallRaw("HttpCache.getUrlSpecificTtl", nil)
	if err != nil {
		return nil, err
	}
	list := struct {
		Result struct {
			List UrlSpecificTtlList `json:"list"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, err
}

// HttpCacheSetUrlSpecificTtl - Gets list of URL specific TTL
// Parameters
//	list - URL specific TTL list
// Return
//	errors - list of errors
func (s *ServerConnection) HttpCacheSetUrlSpecificTtl(list UrlSpecificTtlList) (ErrorList, error) {
	params := struct {
		List UrlSpecificTtlList `json:"list"`
	}{list}
	data, err := s.CallRaw("HttpCache.setUrlSpecificTtl", params)
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

// HttpCacheClearCache - Removes cache content
func (s *ServerConnection) HttpCacheClearCache() error {
	_, err := s.CallRaw("HttpCache.clearCache", nil)
	return err
}
