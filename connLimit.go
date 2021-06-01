package control

import "encoding/json"

type ConnLimitSettings struct {
	SrcLimit         OptionalLong        `json:"srcLimit"`
	SrcRateLimit     OptionalLong        `json:"srcRateLimit"`
	DstLimit         OptionalLong        `json:"dstLimit"`
	DstPerSrcLimit   OptionalLong        `json:"dstPerSrcLimit"`
	Exclusions       OptionalIdReference `json:"exclusions"`
	ExclSrcLimit     OptionalLong        `json:"exclSrcLimit"`
	ExclSrcRateLimit OptionalLong        `json:"exclSrcRateLimit"`
}

// ConnLimitGet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ConnLimitGet() (*ConnLimitSettings, error) {
	data, err := s.CallRaw("ConnLimit.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config ConnLimitSettings `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// ConnLimitSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) ConnLimitSet(config ConnLimitSettings) error {
	params := struct {
		Config ConnLimitSettings `json:"config"`
	}{config}
	_, err := s.CallRaw("ConnLimit.set", params)
	return err
}
