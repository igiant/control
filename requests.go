package control

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type ServerConnection struct {
	Config *Config
	Token  *string
	client *http.Client
}

func (c *Config) NewConnection() (*ServerConnection, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	connection := ServerConnection{
		Config: c,
		client: client,
	}
	return &connection, nil
}

func newClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &http.Client{Jar: jar}, nil
}

func (s *ServerConnection) CallRaw(method string, params interface{}) ([]byte, error) {
	buffer, err := marshal(s.Config.getID(), method, s.Token, params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", s.Config.url, bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "ApiApplication/json-rpc")
	if s.Token != nil {
		req.Header.Add("X-Token", *s.Token)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = checkError(data); err != nil {
		return nil, err
	}
	return data, nil
}

func NewSearchQuery(fields []string, conditions SubConditionList, combining LogicalOperator, start, limit int, orderBy SortOrderList) SearchQuery {
	if fields == nil {
		fields = make([]string, 0)
	}
	if conditions == nil {
		conditions = make(SubConditionList, 0)
	}
	if limit == 0 {
		limit = -1
	}
	if orderBy == nil {
		orderBy = make(SortOrderList, 0)
	}
	return SearchQuery{
		Fields:     fields,
		Conditions: conditions,
		Combining:  combining,
		Start:      start,
		Limit:      limit,
		OrderBy:    orderBy,
	}
}
