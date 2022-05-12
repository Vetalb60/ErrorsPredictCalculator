// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"CourseWork/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name PredictionAlgorithmsServer --spec ../swagger-api/swagger.yaml --principal interface{}

func configureFlags(api *operations.PredictionAlgorithmsServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PredictionAlgorithmsServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.BinProducer = runtime.ByteStreamProducer()
	api.CsvProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("csv producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PostUploadMaxParseMemory = 32 << 20

	if api.DeleteDeleteIDHandler == nil {
		api.DeleteDeleteIDHandler = operations.DeleteDeleteIDHandlerFunc(func(params operations.DeleteDeleteIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteDeleteID has not yet been implemented")
		})
	}
	if api.GetHandler == nil {
		api.GetHandler = operations.GetHandlerFunc(func(params operations.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.Get has not yet been implemented")
		})
	}
	if api.GetHealthzHandler == nil {
		api.GetHealthzHandler = operations.GetHealthzHandlerFunc(func(params operations.GetHealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHealthz has not yet been implemented")
		})
	}
	if api.PostMethodIDHandler == nil {
		api.PostMethodIDHandler = operations.PostMethodIDHandlerFunc(func(params operations.PostMethodIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostMethodID has not yet been implemented")
		})
	}
	if api.PostUploadHandler == nil {
		api.PostUploadHandler = operations.PostUploadHandlerFunc(func(params operations.PostUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUpload has not yet been implemented")
		})
	}
	if api.GetListHandler == nil {
		api.GetListHandler = operations.GetListHandlerFunc(func(params operations.GetListParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetList has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
