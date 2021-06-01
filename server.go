package control

type ServerOs string

const (
	Windows ServerOs = "Windows"
	Linux   ServerOs = "Linux"
)

// Available entities, entity prefix due to name collision
//enum Entity {
//EntityUser,      // User Entity
//EntityAlias,     // Alias Entity
//EntityGroup,     // Group Entity
//EntityMailingList,    // Mailing List Entity
//EntityResource,     // Resource Scheduling Entity
//EntityTimeRange,   // Time Range Entity
//EntityTimeRangeGroup,  // Time Range Group Entity
//EntityIpAddress,   // Ip Address Entity
//EntityIpAddressGroup,  // Ip Address Group Entity
//EntityService,    // Service Entity
//EntityDomain
//};
// RestrictionTuple - Restriction Items
type RestrictionTuple struct {
	Name   string          `json:"name"` // was of type kerio::web::ItemName
	Kind   RestrictionKind `json:"kind"`
	Values StringList      `json:"values"`
}

// RestrictionTupleList - Restriction tuple for 1 entity
type RestrictionTupleList []RestrictionTuple

// Restriction - Entity name restriction definition
type Restriction struct {
	EntityName string               `json:"entityName"` // was of type Entity
	Tuples     RestrictionTupleList `json:"tuples"`     // Restriction tuples
}

// RestrictionList - List of restrictions
type RestrictionList []Restriction

// ServerGetOs -
// Return
//	os - engine OS. I would like to enumerate where client depends on engine OS but such list will become obsolete soon
func (s *ServerConnection) ServerGetOs() (*ServerOs, error) {
	data, err := s.CallRaw("Server.getOs", nil)
	if err != nil {
		return nil, err
	}
	os := struct {
		Result struct {
			Os ServerOs `json:"os"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &os)
	return &os.Result.Os, err
}

// ServerGetRestrictionList -
func (s *ServerConnection) ServerGetRestrictionList() (RestrictionList, error) {
	data, err := s.CallRaw("Server.getRestrictionList", nil)
	if err != nil {
		return nil, err
	}
	restrictions := struct {
		Result struct {
			Restrictions RestrictionList `json:"restrictions"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &restrictions)
	return restrictions.Result.Restrictions, err
}
