package config_test

import (
	"os"
	"testing"

	"github.com/iamgreggarcia/gofigure/api/config"
	"github.com/stretchr/testify/assert"
)

var (
	appName         = "/gofigure/"
	srcGithub       = "src/github.com/"
	gopath          = os.Getenv("GOPATH")
	username        = os.Getenv("GITHUB_USERNAME")
	configPath      = gopath + srcGithub + username + appName + "_config_files/"
	logPath         = gopath + srcGithub + username + appName + "logs/"
	staticFilesPath = gopath + srcGithub + username + appName + "static"
)

const (
	port        = "8080"
	dbType      = "postgres"
	logFilename = "master_dev_logs.log"

	testUsername = "test name"
	testPassword = "test password"
	testDatabase = "test database"
)

func TestConfiguration(t *testing.T) {
	var configuration config.Configuration

	// expected values
	expectedPort := port
	expectedDBConnString := dbType
	expectedUsername := testUsername
	expectedPassword := testPassword
	expectedDatabase := testDatabase
	expectedStaticFilesPath := staticFilesPath
	expectedLogFilename := logFilename

	// actual values
	configuration.Server.Port = port
	configuration.Database.DBType = dbType
	configuration.Database.Username = testUsername
	configuration.Database.Password = testPassword
	configuration.Database.Database = testDatabase
	configuration.Directory.StaticFilesPath = staticFilesPath
	configuration.Logging.OutputFile = logFilename

	// Test expected values against actual; uses assert from testify lib
	assert.Equal(t, expectedPort, configuration.Server.Port)
	assert.Equal(t, expectedDBConnString, configuration.Database.DBType)
	assert.Equal(t, expectedUsername, configuration.Database.Username)
	assert.Equal(t, expectedPassword, configuration.Database.Password)
	assert.Equal(t, expectedDatabase, configuration.Database.Database)
	assert.Equal(t, expectedStaticFilesPath, configuration.Directory.StaticFilesPath)
	assert.Equal(t, expectedLogFilename, configuration.Logging.OutputFile)
}
