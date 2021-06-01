package control

type AntiHammeringConfig struct {
	Enabled          bool        `json:"enabled"`
	WhitelistEnabled bool        `json:"whitelistEnabled"`
	Whitelist        IdReference `json:"whitelist"`
	LockedTime       int         `json:"lockedTime"`
}

// AntiHammeringSet -
func (s *ServerConnection) AntiHammeringSet(config AntiHammeringConfig) error {
	params := struct {
		Config AntiHammeringConfig `json:"config"`
	}{config}
	_, err := s.CallRaw("AntiHammering.set", params)
	return err
}

// AntiHammeringGet -
func (s *ServerConnection) AntiHammeringGet() (*AntiHammeringConfig, error) {
	data, err := s.CallRaw("AntiHammering.get", nil)
	if err != nil {
		return nil, err
	}
	antihammeringCfg := struct {
		Result struct {
			AntihammeringCfg AntiHammeringConfig `json:"antihammeringCfg"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &antihammeringCfg)
	return &antihammeringCfg.Result.AntihammeringCfg, err
}

// AntiHammeringGetBlockedIpCount -
func (s *ServerConnection) AntiHammeringGetBlockedIpCount() (int, error) {
	data, err := s.CallRaw("AntiHammering.getBlockedIpCount", nil)
	if err != nil {
		return 0, err
	}
	count := struct {
		Result struct {
			Count int `json:"count"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &count)
	return count.Result.Count, err
}

// AntiHammeringUnblockAll -
func (s *ServerConnection) AntiHammeringUnblockAll() error {
	_, err := s.CallRaw("AntiHammering.unblockAll", nil)
	return err
}
