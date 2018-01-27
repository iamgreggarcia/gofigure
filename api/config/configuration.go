package config

// Configuration type encapsulates the configurations
// for the serer, database, logging, directory paths, etc.
type Configuration struct {
	Server    ServerConfiguration    `json:"server"`
	Database  DatabaseConfiguration  `json:"database"`
	Directory DirectoryConfiguration `json:"directory"`
	Logging   LoggingConfiguration   `json:"loggin"`
}
