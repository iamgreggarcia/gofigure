package model

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/iamgreggarcia/gofigure/api/config"
	_ "github.com/lib/pq"
)

var file = "../../_config_files/config.test.json"
var database config.Configuration
var conn *sql.DB
var d DB

func TestGetConnection(t *testing.T) {

	c := loadJSON(file)

	conn := GetConnection(c)
	if err := conn.Ping(); err != nil {
		t.Error(err)
	}
}

// loadJSON is a helper function to decode our
// config.test.json file, containing our test values, into
// our DatabaseConfiguration struct
// See config.test.json to adjust values accordingly
func loadJSON(file string) config.Configuration {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully loaded json file.")
	}
	// defer closing our json file
	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&database); err != nil {
		log.Fatal(err)
	}

	return database
}
