# go-crud-rest-api

## Description
- Implemented create/read/read all/update/delete endpoints for Task object, used [Echo](https://github.com/labstack/echo) as a web framework.
- Implemented persistence with MySQL and [GORM](https://github.com/jinzhu/gorm).
- Added login endpoint with JWT token, and simple middleware that checks header for 'Authorization: Bearer %jwt_token%' in each request. Otherwise return 403 and json struct with error.
- Implemented endpoint that will use OAuth2 authorization for FB, to login and issue access_token.
- Added CORS middleware (reply with header Access-Control-Allow-Origin: *).
- Added support for OPTIONS HTTP method for each endpoint.
- Added custom logging of each request including status code with [logrus](https://github.com/sirupsen/logrus).
- Added scripts for DB migrations with [Goose](https://github.com/pressly/goose) library.
- Configured daemon over simple YAML config. File path is specified as process flag for daemon. Required params: ListenAddress, DatabaseUri.
 - Used vendoring with [dep](https://github.com/golang/dep).

### Task 
```
type Task struct {
    Id          int64
    Title       string
    Description string
    Priority    int
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
    CompletedAt *time.Time
    IsDeleted   bool
    IsCompleted bool
}
```

## Install
### Prepare MySQL database
```
CREATE SCHEMA go_crud_rest_api;
CREATE USER 'gouser'@'localhost' IDENTIFIED BY 'gopassword';
GRANT ALL ON go_crud_rest_api.* TO 'gouser'@'localhost';
FLUSH PRIVILEGES;
```
### Migrate database to the most recent version available
Run in terminal:
```
sh prepare-db.sh
```
This will install Goose and run all migrarion scripts from ./db/migrations directory

### Run server
```
sh start.sh
```
The server will start up on 4000 port, Task API available at [http://localhost:4000/api/tasks](http://localhost:4000/api/tasks)

### Update libraries (optional)
To install Glide and update libraries in ./vendor directory, run:
```
sh update.sh
```

## Notes
- Server configuration is stored in YAML file (conf/application.yaml) - we must provide `-config` flag to start a service.
- Added universal CRUD repository, which incapsulates common db operations, any other repository (including task repository) uses it for CRUD operations.
- BaseModel struct contains common fields for all models (includong Task).
- Added custom logger as a middleware in server/config package, server logs are written to ./logs/server.log
- Login controller checks creadentials as form values - for simplicity user name and password are hardcoded:
  **demouser** / **demopassword**
- JWT token is stored in session context (Echo).
- For simplicity all configuration values are put into one config struct.
- Only dev environemt is supported, but we can user Viper or similar tool to load environment specific configuration (e.g. to build service for staging/prod etc).
