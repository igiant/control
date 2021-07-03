package control

import "encoding/json"

type ExportOptions struct {
	Certificates bool `json:"certificates"`
	DhcpLeases   bool `json:"dhcpLeases"`
	Stats        bool `json:"stats"`
}

type ImportedInterface struct {
	ImportedId       KId                `json:"importedId"`
	Name             string             `json:"name"`
	Group            InterfaceGroupType `json:"group"`
	Ip               IpAddress          `json:"ip"`
	Type             InterfaceType      `json:"type"`
	UseForFullImport bool               `json:"useForFullImport"`
	CurrentInterface IdReference        `json:"currentInterface"` // id and system name of the corresponding interface in the currently running machine
	PortId           KId                `json:"portId"`           // only for engine purposes
}

type ImportedInterfaceList []ImportedInterface

type CurrentInterface struct {
	Id               IdReference   `json:"id"` // id and system name of the interface in the currently running machine
	Ip               IpAddress     `json:"ip"`
	MAC              string        `json:"MAC"`
	Type             InterfaceType `json:"type"`
	UseForFullImport bool          `json:"useForFullImport"`
}

type CurrentInterfaceList []CurrentInterface

// ConfigurationExportConfig - Creates backup file and returns id.
// Parameters
//	options - A set of options which configuration variables/list to store in exported file.
// Return
//	fileDownload - description of the output file
func (s *ServerConnection) ConfigurationExportConfig(options ExportOptions) (*Download, error) {
	params := struct {
		Options ExportOptions `json:"options"`
	}{options}
	data, err := s.CallRaw("Configuration.exportConfig", params)
	if err != nil {
		return nil, err
	}
	fileDownload := struct {
		Result struct {
			FileDownload Download `json:"fileDownload"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &fileDownload)
	return &fileDownload.Result.FileDownload, err
}

// ConfigurationGetImportInfo - Returns additiaonal information about imported configuration needed during it's import.
// Parameters
//	fileId - id of uploaded configuration file. (see spec. for uploader)
// Return
//	errors - list of errors \n
//	fullImportPossible - tells whether it is possible to import configuration 1:1 (with IP settings).
//	needIfaceMapping - tells whether it has to be setup iface mapping in case of fullImportPossible
//	importedInterfaces - a list of interfaces loaded from the imported configuration file.
//	currentInterfaces - a list of interfaces available in currently loaded configuration.
func (s *ServerConnection) ConfigurationGetImportInfo(fileId string) (ErrorList, bool, bool, ImportedInterfaceList, CurrentInterfaceList, error) {
	params := struct {
		FileId string `json:"fileId"`
	}{fileId}
	data, err := s.CallRaw("Configuration.getImportInfo", params)
	if err != nil {
		return nil, false, false, nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors             ErrorList             `json:"errors"`
			FullImportPossible bool                  `json:"fullImportPossible"`
			NeedIfaceMapping   bool                  `json:"needIfaceMapping"`
			ImportedInterfaces ImportedInterfaceList `json:"importedInterfaces"`
			CurrentInterfaces  CurrentInterfaceList  `json:"currentInterfaces"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.FullImportPossible, errors.Result.NeedIfaceMapping, errors.Result.ImportedInterfaces, errors.Result.CurrentInterfaces, err
}

// ConfigurationApply - Applies changes obtained from imported configuration file and users interaction.
// Parameters
//	interfaces - a list of interfaces from imported configuration with mappings to the currently present interfeces. This mapping should be ignored in fullimport.
//	id - id of uploaded configuration file. (see spec. for uploader)
//	fullImport - whether to do a full import (overvrite IP & domain setting with imported values)
// Return
//	errors - list of errors \n
func (s *ServerConnection) ConfigurationApply(interfaces ImportedInterfaceList, id string, fullImport bool) (ErrorList, error) {
	params := struct {
		Interfaces ImportedInterfaceList `json:"interfaces"`
		Id         string                `json:"id"`
		FullImport bool                  `json:"fullImport"`
	}{interfaces, id, fullImport}
	data, err := s.CallRaw("Configuration.apply", params)
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
