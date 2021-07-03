package control

import "encoding/json"

type StatusFunction struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

type StatusFunctionList []StatusFunction

type HttpLogType string

const (
	HttpLogApache HttpLogType = "HttpLogApache"
	HttpLogSquid  HttpLogType = "HttpLogSquid"
)

// LoggerLogWrite - Write a message to given log
// Parameters
//	message - text to be written into log file
func (s *ServerConnection) LoggerLogWrite(logType LogType, message string) error {
	params := struct {
		LogType LogType `json:"logType"`
		Message string  `json:"message"`
	}{logType, message}
	_, err := s.CallRaw("Logger.logWrite", params)
	return err
}

// LoggerGetStatusFunctionList - Returns list of Functions displayed in debug log context menu 'Show status'
// Return
//  functions - list of Functions displayed in debug log context menu 'Show status'
func (s *ServerConnection) LoggerGetStatusFunctionList() (StatusFunctionList, error) {
	data, err := s.CallRaw("Logger.getStatusFunctionList", nil)
	if err != nil {
		return nil, err
	}
	functions := struct {
		Result struct {
			Functions StatusFunctionList `json:"functions"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &functions)
	return functions.Result.Functions, err
}

// LoggerCallStatusFunction - Calls function from StatusFunctionList referenced by id
// Parameters
//	id - ID function
func (s *ServerConnection) LoggerCallStatusFunction(id KId) error {
	params := struct {
		Id KId `json:"id"`
	}{id}
	_, err := s.CallRaw("Logger.callStatusFunction", params)
	return err
}

// LoggerGetHttpLogType - Returns actual Http log type
// Return
//  logType - actual Http log type
func (s *ServerConnection) LoggerGetHttpLogType() (*HttpLogType, error) {
	data, err := s.CallRaw("Logger.getHttpLogType", nil)
	if err != nil {
		return nil, err
	}
	logType := struct {
		Result struct {
			LogType HttpLogType `json:"logType"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &logType)
	return &logType.Result.LogType, err
}

// LoggerSetHttpLogType - Stores Http log type
// Parameters
//  ogType - Http log type
func (s *ServerConnection) LoggerSetHttpLogType(logType HttpLogType) error {
	params := struct {
		LogType HttpLogType `json:"logType"`
	}{logType}
	_, err := s.CallRaw("Logger.setHttpLogType", params)
	return err
}

// LoggerGetLogExpression - Returns expression for dialog from debug log context menu 'IP Traffic...'
// Return
//  expression - expression for dialog from debug log context menu 'IP Traffic...'
func (s *ServerConnection) LoggerGetLogExpression() (string, error) {
	data, err := s.CallRaw("Logger.getLogExpression", nil)
	if err != nil {
		return "", err
	}
	expression := struct {
		Result struct {
			Expression string `json:"expression"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &expression)
	return expression.Result.Expression, err
}

// LoggerSetLogExpression - Stores expression from debug log context menu dialog 'IP Traffic...'
// Parameters
//  expression - expression for dialog from debug log context menu 'IP Traffic...'
func (s *ServerConnection) LoggerSetLogExpression(expression string) error {
	params := struct {
		Expression string `json:"expression"`
	}{expression}
	_, err := s.CallRaw("Logger.setLogExpression", params)
	return err
}

// LoggerGetPacketLogFormat - Returns format for dialog from debug log context menu 'Packet Log format...'
// Return
//  format - format for dialog from debug log context menu 'Packet Log format...'
func (s *ServerConnection) LoggerGetPacketLogFormat() (string, error) {
	data, err := s.CallRaw("Logger.getPacketLogFormat", nil)
	if err != nil {
		return "", err
	}
	format := struct {
		Result struct {
			Format string `json:"format"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &format)
	return format.Result.Format, err
}

// LoggerSetPacketLogFormat - Stores format from debug log context menu dialog 'Packet Log format...'
// Parameters
//  format - format for dialog from debug log context menu 'Packet Log format...'
func (s *ServerConnection) LoggerSetPacketLogFormat(format string) error {
	params := struct {
		Format string `json:"format"`
	}{format}
	_, err := s.CallRaw("Logger.setPacketLogFormat", params)
	return err
}
