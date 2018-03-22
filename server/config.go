package server

// Configuration for storing configuration
type config struct {
	Server   string             `json:"server"`
	Database string             `json:"database"`
	Host     string             `json:"host"`
	Port     int                `json:"port"`
	GUIPath  string             `json:"guiPath,omitempty"`
	CORS     *CORSConfiguration `json:"cors,omitempty"`
}

// CORSConfiguration for storing CORS Configura
type CORSConfiguration struct {
	Enable         bool     `json:"enable,omitempty"`
	AllowedHeaders []string `json:"allowedHeaders,omitempty"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	AllowedMethods []string `json:"allowedMethods,omitempty"`
}
