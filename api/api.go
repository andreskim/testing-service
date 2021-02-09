package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// API Main API data structure
type API struct {
}

// NewAPI API constructor
func NewAPI() *API {
	return &API{}
}

/*
* Miscellaneous functions
 */

// MakeSwagger Instantiates swagger instance and set server field to `nil`
func MakeSwagger() *openapi3.Swagger {
	swagger, err := GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	return swagger
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendError(ctx echo.Context, code int, message string) error {
	err := Error{
		Code:    int32(code),
		Message: message,
	}
	jsonErr := ctx.JSON(code, err)
	return jsonErr
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshall that.
func sendJSON(ctx echo.Context, code int, message interface{}) error {
	jsonErr := ctx.JSON(code, message)
	return jsonErr
}

/*
* API implementation
 */

// HealthCheck If replies, it's healthy
func (api *API) HealthCheck(ctx echo.Context) error {
	return sendError(ctx, http.StatusOK, "Yup, I'm here")
}

// ReadyCheck Readiness check. Responds 200 if can access repo and keyvaults
func (api *API) ReadyCheck(ctx echo.Context) error {
	// It's a good idea to check all dependencies are OK: database connections, API endpoints, etc
	//
	//var err error
	// ready, err := api.sops2Kv.IsReady()
	// if err != nil {
	// 	return sendError(ctx, http.StatusServiceUnavailable, fmt.Sprintf("Not Ready! %s", err.Error()))
	// }
	// if !ready {
	// 	return sendError(ctx, http.StatusServiceUnavailable, "Not Ready, but no error produced... Is it a bug?")
	// }

	return sendError(ctx, http.StatusOK, "Ready!")
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		 r.ParseForm()
		 request = append(request, "\n")
		 request = append(request, r.Form.Encode())
	}
	 // Return the request as a string
	 return strings.Join(request, "\n")
}

// GetEcho Business logic...
func (api *API) GetEcho(ctx echo.Context) error {
	return sendError(ctx, http.StatusOK, formatRequest(ctx.Request()))
}