package control

type ActiveTool string

const (
	ActiveToolNone       ActiveTool = "ActiveToolNone"
	ActiveToolPing       ActiveTool = "ActiveToolPing"
	ActiveToolTraceRoute ActiveTool = "ActiveToolTraceRoute"
	ActiveToolDns        ActiveTool = "ActiveToolDns"
	ActiveToolWhois      ActiveTool = "ActiveToolWhois"
)

type DnsTool string

const (
	DnsToolNslookup DnsTool = "DnsToolNslookup"
	DnsToolDig      DnsTool = "DnsToolDig"
)

type DnsType string

const (
	DnsTypeAny   DnsType = "DnsTypeAny"
	DnsTypeA     DnsType = "DnsTypeA"
	DnsTypeAAAA  DnsType = "DnsTypeAAAA"
	DnsTypeCname DnsType = "DnsTypeCname"
	DnsTypeMx    DnsType = "DnsTypeMx"
	DnsTypeNs    DnsType = "DnsTypeNs"
	DnsTypePtr   DnsType = "DnsTypePtr"
	DnsTypeSoa   DnsType = "DnsTypeSoa"
	DnsTypeSpf   DnsType = "DnsTypeSpf"
	DnsTypeSrv   DnsType = "DnsTypeSrv"
	DnsTypeTxt   DnsType = "DnsTypeTxt"
)

type IpVersion string

const (
	IpVersion4   IpVersion = "IpVersion4"
	IpVersion6   IpVersion = "IpVersion6"
	IpVersionAny IpVersion = "IpVersionAny"
)

// IpToolsGetStatus - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsGetStatus() (*ActiveTool, StringList, error) {
	data, err := s.CallRaw("IpTools.getStatus", nil)
	if err != nil {
		return nil, nil, err
	}
	activeTool := struct {
		Result struct {
			ActiveTool ActiveTool `json:"activeTool"`
			Lines      StringList `json:"lines"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &activeTool)
	return &activeTool.Result.ActiveTool, activeTool.Result.Lines, err
}

// IpToolsStop - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsStop() error {
	_, err := s.CallRaw("IpTools.stop", nil)
	return err
}

// IpToolsPing - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsPing(target string, ipv IpVersion, infinite bool, packetSize int, allowFragmentation bool) error {
	params := struct {
		Target             string    `json:"target"`
		Ipv                IpVersion `json:"ipv"`
		Infinite           bool      `json:"infinite"`
		PacketSize         int       `json:"packetSize"`
		AllowFragmentation bool      `json:"allowFragmentation"`
	}{target, ipv, infinite, packetSize, allowFragmentation}
	_, err := s.CallRaw("IpTools.ping", params)
	return err
}

// IpToolsTraceRoute - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsTraceRoute(target string, ipv IpVersion, resolveHostnames bool) error {
	params := struct {
		Target           string    `json:"target"`
		Ipv              IpVersion `json:"ipv"`
		ResolveHostnames bool      `json:"resolveHostnames"`
	}{target, ipv, resolveHostnames}
	_, err := s.CallRaw("IpTools.traceRoute", params)
	return err
}

// IpToolsWhois - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsWhois(target string) error {
	params := struct {
		Target string `json:"target"`
	}{target}
	_, err := s.CallRaw("IpTools.whois", params)
	return err
}

// IpToolsDns - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsDns(name string, server string, tool DnsTool, dnsType DnsType) error {
	params := struct {
		Name    string  `json:"name"`
		Server  string  `json:"server"`
		Tool    DnsTool `json:"tool"`
		DnsType DnsType `json:"dnsType"`
	}{name, server, tool, dnsType}
	_, err := s.CallRaw("IpTools.dns", params)
	return err
}

// IpToolsGetDnsServers - 1004 Access denied  - "Insufficient rights to perform the requested operation."
func (s *ServerConnection) IpToolsGetDnsServers() (StringList, error) {
	data, err := s.CallRaw("IpTools.getDnsServers", nil)
	if err != nil {
		return nil, err
	}
	servers := struct {
		Result struct {
			Servers StringList `json:"servers"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &servers)
	return servers.Result.Servers, err
}
