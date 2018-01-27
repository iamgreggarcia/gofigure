package config

// ServerConfiguration struct
type ServerConfiguration struct {
	Port string `json:"port"`
}

// GetPort returns the port from
// ServerConfiguration
func (s *ServerConfiguration) GetPort() string {
	return s.Port
}
