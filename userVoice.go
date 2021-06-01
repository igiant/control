package control

import "encoding/json"

// UserVoiceSettings - Settings of UserVoice
type UserVoiceSettings struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserVoiceGetUrl - 8000 Internal error.  - "URL encoding of token failed."
// Return
//	url - URL to userVoice with single sign on token
func (s *ServerConnection) UserVoiceGetUrl() (string, error) {
	data, err := s.CallRaw("UserVoice.getUrl", nil)
	if err != nil {
		return "", err
	}
	url := struct {
		Result struct {
			Url string `json:"url"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &url)
	return url.Result.Url, err
}

// UserVoiceSet - 1004 Access denied.  - "Insufficient rights to perform the requested operation." \n
// Parameters
//	settings - structure with settings
func (s *ServerConnection) UserVoiceSet(settings UserVoiceSettings) error {
	params := struct {
		Settings UserVoiceSettings `json:"settings"`
	}{settings}
	_, err := s.CallRaw("UserVoice.set", params)
	return err
}
