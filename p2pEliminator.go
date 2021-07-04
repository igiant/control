package control

import "encoding/json"

// P2pEliminatorConfig - must be included after BandwidthLimiter
type P2pEliminatorConfig struct {
	Ports                string  `json:"ports"`
	ConnectionCount      int     `json:"connectionCount"`
	TrustedServiceIdList KIdList `json:"trustedServiceIdList"`
}

// P2pEliminatorGet - Returns the configuration of P2P eliminator.
// Return
//	config - structure with configuration of P2P eliminator.
func (s *ServerConnection) P2pEliminatorGet() (*P2pEliminatorConfig, error) {
	data, err := s.CallRaw("P2pEliminator.get", nil)
	if err != nil {
		return nil, err
	}
	config := struct {
		Result struct {
			Config P2pEliminatorConfig `json:"config"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &config)
	return &config.Result.Config, err
}

// P2pEliminatorSet - Sets the configuration of P2P eliminator.
// Parameters
//	config - structure with configuration of P2P eliminator.
// Return
//	errors - list of errors \n
func (s *ServerConnection) P2pEliminatorSet(config P2pEliminatorConfig) (ErrorList, error) {
	params := struct {
		Config P2pEliminatorConfig `json:"config"`
	}{config}
	data, err := s.CallRaw("P2pEliminator.set", params)
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
