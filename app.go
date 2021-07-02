package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rpsingh21/checklist-api/config"
	"github.com/rpsingh21/checklist-api/database"
	"github.com/rpsingh21/checklist-api/handler"
	"github.com/rpsingh21/checklist-api/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.uber.org/zap"
)

// App object entry point of application
// It will connet logger, router and DB object
type App struct {
	logger *zap.SugaredLogger
	server *http.Server
	Router *mux.Router
	DB     *mongo.Database
}

// ConfigApp setup app
func ConfigApp(logger *zap.SugaredLogger, config *config.Config) *App {
	app := &App{logger: logger}
	app.Initialize(config)
	return app
}

// Initialize app
func (app *App) Initialize(config *config.Config) {
	app.DB = database.NewDBConnection(config.DatabaseName, config.MongoConnectionURI)
	app.createIndex()

	app.Router = mux.NewRouter()
	app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.setRouter()

	app.server = &http.Server{
		Addr:    config.ServerHost, // configure the bind address
		Handler: app.Router,        // set the default handler
		// ErrorLog:     (*log.Logger)(logger.Sugar()), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
}

// UseMiddleware will add global middleware in router
func (app *App) UseMiddleware(middleware mux.MiddlewareFunc) {
	app.Router.Use(middleware)
}

// Get will register Get method for an endpoint
func (app *App) Get(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("GET").Queries(queries...)
}

// Post will register Post method for an endpoint
func (app *App) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

// Put will register Put method for an endpoint
func (app *App) Put(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PUT").Queries(queries...)
}

// Patch will register Patch method for an endpoint
func (app *App) Patch(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PATCH").Queries(queries...)
}

// Delete will register Delete method for an endpoint
func (app *App) Delete(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("DELETE").Queries(queries...)
}

// Run will start the http server on host.
func (app *App) Run() {
	app.logger.Infof("Starting server on %s .....", app.server.Addr)
	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error(err)
	}
}

// Shutdown application
func (app *App) Shutdown(ctx context.Context) {
	app.logger.Info("Shutdown application .....")
	app.server.Shutdown(ctx)
	app.DB.Client().Disconnect(ctx)
}

func (app *App) setRouter() {
	userRopo := repository.NewUserRepository(app.DB)
	ah := handler.NewAuthHandler(app.logger, userRopo)
	app.Router.HandleFunc("/auth", ah.Get).Methods(http.MethodGet)
	app.Router.HandleFunc("/auth", ah.Create).Methods(http.MethodPost)
}

func (app *App) createIndex() {
	keys := bsonx.Doc{
		{Key: "username", Value: bsonx.Int32(1)},
	}
	user := app.DB.Collection("user")
	database.SetIndexes(user, keys)
}

// RequestHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
