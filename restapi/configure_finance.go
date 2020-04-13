// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gofinance/internal"
	"github.com/gofinance/models"
	"github.com/gofinance/restapi/operations"
	"github.com/gofinance/restapi/operations/financeapi"
)

//go:generate swagger generate server --target ../../go-finance-api --name Finance --spec ../swagger.yml

func configureFlags(api *operations.FinanceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FinanceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set

	api.ApikeyAuth = func(token string) (interface{}, error) {
		if token == "lol" {
			prin := models.Principal(token)
			return &prin, nil
		}
		return nil, errors.New(401, "incorrect api key auth")
	}

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
		if user == "root" && pass == "root" {
			prin := models.Principal(user)
			return &prin, nil
		}
		return nil, errors.New(401, "incorrect base auth")
	}

	api.FinanceapiGetPutPriceHandler = financeapi.GetPutPriceHandlerFunc(func(params financeapi.GetPutPriceParams, principal interface{}) middleware.Responder {

		put_price, err := internal.Get_put_price(*params.PutPrice.TimeToMaturity, *params.PutPrice.SpotPrice, *params.PutPrice.StrikePrice, *params.PutPrice.RiskFreeRate, *params.PutPrice.Sigma)
		if err != nil {
			resp_err := financeapi.NewGetCallPriceDefault(500)
			err := err.Error()
			erro := models.Error{Code: 500, Message: &err}
			resp_err.SetPayload(&erro)
			return resp_err
		}
		resp_ok := financeapi.NewGetCallPriceOK()
		resp_ok.SetPayload(put_price)
		return resp_ok
	})

	//actuall handiling of get call price api call
	api.FinanceapiGetCallPriceHandler = financeapi.GetCallPriceHandlerFunc(func(params financeapi.GetCallPriceParams, principal interface{}) middleware.Responder {
		call_price, err := internal.Get_call_price(*params.CallPrice.TimeToMaturity, *params.CallPrice.SpotPrice, *params.CallPrice.StrikePrice, *params.CallPrice.RiskFreeRate, *params.CallPrice.Sigma)
		if err != nil {
			resp_err := financeapi.NewGetCallPriceDefault(500)
			err := err.Error()
			erro := models.Error{Code: 500, Message: &err}
			resp_err.SetPayload(&erro)
			return resp_err
		}
		resp_ok := financeapi.NewGetCallPriceOK()
		resp_ok.SetPayload(call_price)
		return resp_ok
	})

	api.FinanceapiMovingaverageHandler = financeapi.MovingaverageHandlerFunc(func(params financeapi.MovingaverageParams, principal interface{}) middleware.Responder {
		mov, err := internal.Get_movingaverage(params.Movingaverage.TimeData, int(*params.Movingaverage.PointToAvg))
		if err != nil {
			resp_err := financeapi.NewMovingaverageDefault(500)
			err := err.Error()
			erro := models.Error{Code: 500, Message: &err}
			resp_err.SetPayload(&erro)
			return resp_err
		}
		resp_ok := financeapi.NewMovingaverageOK()
		resp_ok.SetPayload(mov)
		return resp_ok
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
