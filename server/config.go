package server

// Configuration for storing configuration
type config struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	GUIPath  string `json:"guiPath,omitempty"`
}
