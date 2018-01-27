package config

// DatabaseConfiguration is a struct for database
// username, password, and database set by environment
// variables and imported via viper
type DatabaseConfiguration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	DBType   string `json:"dbtype"`
}
