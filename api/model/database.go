package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/iamgreggarcia/gofigure/api/config"
	// pq is the postgresql driver we are using
	_ "github.com/lib/pq"
)

// DB struct contains the SQL database object
type DB struct {
	DB *sql.DB
}

// GetConnection opens a new connection pool
func GetConnection(config config.Configuration) *sql.DB {
	DB := openDBConnection(config)
	return DB
}

func openDBConnection(config config.Configuration) *sql.DB {
	user := config.Database.Username
	password := config.Database.Password
	dbname := config.Database.Database
	dbType := config.Database.DBType
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user,
		password,
		dbname)

	var err error
	DB, err := sql.Open(dbType, connectionString)
	if err != nil {
		color.Set(color.FgHiRed)
		color.Set(color.BgBlack)
		log.Fatal(err)
	} else {
		s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
		s.Prefix = "Connectings to database..."
		s.Start()
		time.Sleep(2 * time.Second)
		s.Stop()
		color.Set(color.FgHiGreen)
		log.Println("Database connection successful.")
		color.Unset()
	}
	return DB
}
