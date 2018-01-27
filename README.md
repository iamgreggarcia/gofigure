# Gofigure
A minimal Go RESTful web application using Bootstrap and [mux](https://github.com/gorilla/mux).

[![Build Status](https://travis-ci.org/iamgreggarcia/gofigure.svg?branch=master)](https://travis-ci.org/iamgreggarcia/gofigure)

## Getting Started

1. Be sure [Go](https://golang.org/doc/install) is properly installed, along with a configured [$GOPATH](https://github.com/golang/go/wiki/SettingGOPATH).  

2. Set the environment variables `GITHUB_USERNAME`and `CONFIG_CF`. For development, set `CONFIG_CF=config.development` and for production set `CONFIG_CF=config.production`. The corresponding JSON config files can be found in `/._config_files`/.
```shell
$ export GITHUB_USERNAME=<your_github_username>
$ export CONFIG_CF=config.development
```


3. Set environment variables `username`, `password`, and `database` for [Postgresql](https://www.postgresql.org/), and be sure to create them in postgresql as well:
```shell
$ export GOFIGURE_USERNAME=<your_postgres_username_for_database>
$ export GOFIGURE_PASSWORD=<your_password>
$ export GOFIGURE_DATABASE=<database_name>
```

## Running

To run the application:
```shell
$ sh clean.sh
$ sh run.sh
```
Navigate to http://localhost:8080 and you should see the following page:
![alt tag](https://github.com/iamgreggarcia/gofigure/static/img/index.png)


## Project Structure

The main entry of for the application is is `main.go`, where the configuration
file is loaded, a database connection is created, and the application is started.

The server-side code is located in `api/`:
```shell
$ tree api/
api/
├── app
│   ├── app.go
│   └── app_test.go
├── config
│   ├── configuration.go
│   ├── configuration_test.go
│   ├── database.go
│   ├── directory.go
│   ├── log.go
│   └── server.go
├── handler
│   ├── handlers.go
│   └── handlers_test.go
└── model
    ├── database.go
    ├── database_test.go
    ├── message.go
    └── model.go


```


There is built in support for creating api endpoints:

```go
// ./api/app/app.go
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
```

All client-side code is location in the `static/` directory:
```shell
$ tree static
static
├── css
│   └── main.css
├── img
│   └── index.png
├── index.html
├── js
└── vendor
    └── bootstrap-4.0.0
```

