package main

import (
	"fmt"
	"net/http"
	"os"
	"socialNetwork/internal/server"
	"socialNetwork/pkg/assert"
	config "socialNetwork/pkg/config"
	"testing"
)

func init() {
	// Initialization code that runs before tests
	fmt.Println("Initializing resources before tests...")
}

// Declare global variables for the app and test server
var (
	app *server.App
	cfg *config.Conf
	ts  *testServer
)

// TestMain is executed before any tests are run
func TestMain(m *testing.M) {
	// Initialize the application and server once
	app, cfg = server.NewTestApplication()
	ts = newTestServer(nil, app.InitRoutes(cfg))

	// Run the tests
	exitCode := m.Run()

	// Close the server and clean up
	ts.Close()

	// Exit with the test result code
	os.Exit(exitCode)
}

func TestPing(t *testing.T) {
	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
