package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	cfg "github.com/iamgreggarcia/gofigure/api/config"
	"github.com/iamgreggarcia/gofigure/api/handler"
)

var config cfg.Configuration

// App is the entry point from main.go
// to start the application
type App struct {
	Router *mux.Router
}

type server struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

// newServer creates a new server object
func newServer(a *App) *server {
	port := ":" + config.Server.Port
	// create server
	s := &server{
		Server: http.Server{
			Addr:         port,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}
	s.Handler = a.Router

	return s
}

// Initialize mux router
func (a *App) Initialize(appConfig *cfg.Configuration) {
	config = *appConfig

	a.Router = mux.NewRouter()
	a.initializeRoutes(&config)
}

// Run starts the server that will listen
// on parameter "port"
func (a *App) Run() {
	// Create new spinner
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)

	// Set colors for logging
	color.Set(color.FgHiGreen)
	log.Println("Starting server...")
	s.Prefix = "Starting: "
	s.FinalMSG = "Server Running!\n"
	s.Start()
	time.Sleep(2 * time.Second)

	// Get port address from configuration
	port := config.Server.Port

	// Create server
	server := newServer(a)
	s.Stop()
	color.Set(color.FgHiGreen)
	log.Println("Server started on port: " + port)
	log.Println("Navigate to http://localhost:" + port)
	color.Unset()

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	//log.Fatal(http.ListenAndServe(":"+port, a.Router))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan // wait for SIGINT
	color.Set(color.FgHiYellow)
	s.Prefix = "Draining server connections..: "
	s.FinalMSG = "Completed.\n"
	log.Println("\nShutting down server...")
	color.Unset()
	s.Start()
	time.Sleep(2 * time.Second)
	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Stop()
	color.Set(color.FgHiGreen)
	log.Println("Server gracefully stopped.")
	server.Shutdown(ctx)
}

// initializeRoutes creates API endpoints and intializes any static files to be served
func (a *App) initializeRoutes(routeConfig *cfg.Configuration) {
	// Initialize config for handler
	handler.Initialize(routeConfig)
	// API Endpoints
	// Register your API endpoints here. E.g.:
	// api := a.Router.PathPrefix("/api/v1").Subrouter() <-- create appropriate
	// path prefix for Subrouter
	// api.Methods("GET").Path("/<dir>").HandlerFunc(<your_handler>)
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.Methods("GET").Path("/greetings/hello/{name}").HandlerFunc(handler.GetHelloHandler)

	// Static files
	var static string
	dir := routeConfig.Directory.StaticFilesPath
	flag.StringVar(&static, "static", dir, "The directory form which to serve static files.")
	flag.Parse()
	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(static))))
}
