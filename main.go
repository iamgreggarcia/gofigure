package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/iamgreggarcia/gofigure/api/app"
	"github.com/iamgreggarcia/gofigure/api/config"
	"github.com/iamgreggarcia/gofigure/api/model"
	"github.com/spf13/viper"
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
	envConfig = "CONFIG_CF"
)

var dbConn model.DB
var configuration config.Configuration

func main() {
	// Get config filename for env variable
	cf := os.Getenv(envConfig)

	// Set the config filename for viper
	viper.SetConfigName(cf)
	// Add path in which viper will look for config
	viper.AddConfigPath(configPath)

	// Handle err
	// TODO: configure defaults if config file not found
	if err := viper.ReadInConfig(); err != nil {
		color.Set(color.FgHiRed)
		color.Set(color.BgBlack)
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Get environment variables for DB access
	// Set prefix for env vars, and bind the key names used in the config json file
	// along with the environment variable name itself (second param in BindEnv())
	viper.SetEnvPrefix("gostarter")
	// Database username, password, and database name
	viper.BindEnv("database.username", "GOSTARTER_USERNAME")
	viper.BindEnv("database.password", "GOSTARTER_PASSWORD")
	viper.BindEnv("database.database", "GOSTARTER_DATABASE")

	// Unmarshal viper config values into configuration struct
	err := viper.Unmarshal(&configuration)
	if err != nil {
		color.Set(color.FgHiRed)
		color.Set(color.BgBlack)
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	configuration.Directory.StaticFilesPath = staticFilesPath
	// Setup logging configuration.
	masterLogPath := logPath + configuration.Logging.OutputFile
	logLocalPath := logPath + "test/" + configuration.Logging.LocalFile

	// Check if log directories exist
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, 0700)
	}

	if _, err := os.Stat(logPath + "test/"); os.IsNotExist(err) {
		os.Mkdir(logPath+"test/", 0700)
	}

	logFile, err := os.OpenFile(masterLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		color.Set(color.FgHiRed)
		color.Set(color.BgBlack)
		panic(fmt.Errorf("unable to retrieve master log file and path, %v", err))
	}

	logLocalFile, err := os.OpenFile(logLocalPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		color.Set(color.FgHiRed)
		color.Set(color.BgBlack)
		panic(fmt.Errorf("unable to retrieve local session log file and path, %v", err))
	}

	// Use multiwriter to write to both console (os.Stdout), the master
	// log file (logfile, and the local session log file (localLogFile).
	mw := io.MultiWriter(os.Stdout, logFile, logLocalFile)
	log.SetOutput(mw)
	// Set "green" color output for session start
	color.Set(color.FgHiGreen)
	// Defer close for when we are done with the files.
	defer logFile.Close()
	defer logLocalFile.Close()
	log.Println("Log session started.")
	color.Unset()
	log.Printf("View logs at:\n")
	log.Printf("Master log: ./logs/%s\n", configuration.Logging.OutputFile)
	log.Printf("Session logs: ./logs/test/%s\n", configuration.Logging.LocalFile)

	// Connect to database, and make it accesible for rest of app
	dbConn.DB = model.GetConnection(configuration)

	a := app.App{}
	a.Initialize(&configuration)
	a.Run()
}
