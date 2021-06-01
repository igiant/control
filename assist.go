package control

// Importance - Certificate Time properties info
type Importance string

const (
	MainProcess       Importance = "MainProcess"
	TestingImportance Importance = "TestingImportance" // Only for testing purposes
	OtherProcess      Importance = "OtherProcess"
)
