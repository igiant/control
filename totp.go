package control

import "encoding/json"

type TotpState string

const (
	TotpDone          TotpState = "TotpDone"
	TotpConfigure     TotpState = "TotpConfigure"
	TotpNotConfigured TotpState = "TotpNotConfigured"
	TotpVerify        TotpState = "TotpVerify"
)

// TotpTotpVerify - Performs 2 step verification
// Parameters
//  code -
//  remember -
func (s *ServerConnection) TotpTotpVerify(code int, remember bool) error {
	params := struct {
		Code     int  `json:"code"`
		Remember bool `json:"remember"`
	}{code, remember}
	_, err := s.CallRaw("Totp.totpVerify", params)
	return err
}

// TotpTotpState - checks 2 step verification state
// Return
//  state -
func (s *ServerConnection) TotpTotpState() (*TotpState, error) {
	data, err := s.CallRaw("Totp.totpState", nil)
	if err != nil {
		return nil, err
	}
	state := struct {
		Result struct {
			State TotpState `json:"state"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &state)
	return &state.Result.State, err
}
