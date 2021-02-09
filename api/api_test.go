package api

import (
	"flag"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	echo "github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	config "github.com/resideo/testing-service/internal/config"
	"github.com/stretchr/testify/assert"
)

// Glocal variables used in tests
var (
	ge       *echo.Echo
	testpath string
)

func TestGetSwagger(t *testing.T) {
	_, err := GetSwagger()
	assert.EqualValuesf(t, err, nil, "Get swagger - want no error, got one")
}

// Send test request - harness for echo
func testSendRequest(method string, target string, body io.Reader, handler func(echo.Context) error) ([]byte, int, error) {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := ge.NewContext(req, w)
	err := handler(ctx)
	if err != nil {
		return nil, 0, err
	}
	resp := w.Result()
	b, err := ioutil.ReadAll(resp.Body)
	c := resp.StatusCode

	return b, c, err
}

func TestHealthCheck(t *testing.T) {
	target := "http://localhost:8080/healthz"
	body, code, err := testSendRequest("GET", target, nil, gsi.HealthCheck)
	want := 200
	assert.Nilf(t, err, "Failed parsing mock request to %s: %s", target, err)
	assert.EqualValuesf(t, want, code, "Want code %d, got %d, body: %s", want, code, body)
}

func TestReadyCheck(t *testing.T) {
	target := "http://localhost:8080/readyz"
	body, code, err := testSendRequest("GET", target, nil, gsi.ReadyCheck)
	want := 200

	assert.Nilf(t, err, "Failed parsing mock request to %s: %s", target, err)
	assert.EqualValuesf(t, want, code, "Want code %d, got %d, body: %s", want, code, body)
}

// Set up echo fixture and things
func TestMain(m *testing.M) {
	ge = echo.New()

	testClonePath, tfRemove := testdata.GetTempFolder("_api")
	defer tfRemove()
	// Set test configs here
	// os.Setenv("GITREPO_URL", testdata.TestGitRepoURL)


	config.SetTestEnv()
	config.ValidateENV()
	config.ReadConfig()
	flag.Parse()
	gsi.Handler = gs2k
	swagger := MakeSwagger()

	// Log all requests
	ge.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	ge.Use(middleware.OapiRequestValidator(swagger))

	// We now register above as the handler for the interface
	RegisterHandlers(ge, gs2k)
	os.Exit(m.Run())
}
