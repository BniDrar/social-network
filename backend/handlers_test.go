package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"socialNetwork/internal/server"
	"socialNetwork/pkg/assert"
	config "socialNetwork/pkg/config"
)

const (
	existNeckName = "existuser@example.com"
	existLast     = "user"
	existFirst    = "exist"
	existEmail    = "existuser@example.com"

	validNeckName    = "validuser"
	validFirst       = "valid"
	validLast        = "user"
	validDateOfBirth = "01/01/2000"
	validPassword    = "validPa$$word1"
	validEmail       = "validuser@example.com"

	InvalidEmail = "BadEmail"
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
	defer ts.Close()

	// Exit with the test result code
	os.Exit(exitCode)
}

// func TestPing(t *testing.T) {
// 	// app, cfg := server.NewTestApplication()
// 	// ts := newTestServer(t, app.InitRoutes(cfg)) // http.Server{}
// 	// defer ts.Close()
// 	code, _, body := ts.get(t, "/ping")
// 	assert.Equal(t, code, http.StatusOK)
// 	assert.Equal(t, body, "OK")
// }

// func TestRegister(t *testing.T) {
// 	// Create the application struct containing our mocked dependencies and set
// 	// up the test server for running an end-to-end test.
// 	// app, cfg := server.NewTestApplication()
// 	// ts := newTestServer(t, app.InitRoutes(cfg))
// 	// defer ts.Close()

// 	// Make a GET /user/signup request and then extract the CSRF token from the
// 	// response body.
// 	_, _, _ = ts.get(t, "/api/register")

// 	// csrfToken := extractCSRFToken(t, body)
// 	// Log the CSRF token value in our test output using the t.Logf() function.
// 	// The t.Logf() function works in the same way as fmt.Printf(), but writes
// 	// the provided message to the test output.

// 	tests := []struct {
// 		name        string
// 		nickname    string
// 		email       string
// 		password    string
// 		first       string
// 		last        string
// 		dateOfBirth string
// 		wantCode    int
// 	}{
// 		{
// 			name:        "Valid submission",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusCreated,
// 		},
// 		{
// 			name:        "Empty First Name",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    validPassword,
// 			first:       "",
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Empty LastName",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        "",
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Empty NickName",
// 			nickname:    "",
// 			email:       validEmail,
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Empty email",
// 			nickname:    validNeckName,
// 			email:       "",
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Empty password",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    "",
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Invalid email",
// 			nickname:    validNeckName,
// 			email:       InvalidEmail,
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Short password",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    "short",
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "password With No Digit",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    "passwordWithNoDigit",
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "password_with_no_upper_case1",
// 			nickname:    validNeckName,
// 			email:       validEmail,
// 			password:    validPassword,
// 			first:       validFirst,
// 			last:        validLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 		{
// 			name:        "Duplicate email",
// 			nickname:    existNeckName,
// 			email:       existEmail,
// 			password:    validPassword,
// 			first:       existFirst,
// 			last:        existLast,
// 			dateOfBirth: validDateOfBirth,
// 			wantCode:    http.StatusBadRequest,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 				Nickname    string `json:"nickname"`
// 				Email       string `json:"email"`
// 				Password    string `json:"password"`
// 				First       string `json:"first"`
// 				Last        string `json:"last"`
// 				DateOfBirth string `json:"date_of_birth"`
// 			}{
// 				Nickname:    tt.nickname,
// 				Email:       tt.email,
// 				Password:    tt.password,
// 				First:       tt.first,
// 				Last:        tt.last,
// 				DateOfBirth: tt.dateOfBirth,
// 			}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			code, _, _ := ts.postJSON(t, "/api/register", jsonBody)

// 			assert.Equal(t, code, tt.wantCode)
// 			if tt.nickname != existNeckName {
// 				err = app.User.DeleteUserByNickName(tt.nickname)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 			}
// 		})
// 	}
// }

// func TestLogin(t *testing.T) {
// 	// Create the application struct containing our mocked dependencies and set
// 	// up the test server for running an end-to-end test.
// 	// app, cfg := server.NewTestApplication()
// 	// ts := newTestServer(t, app.InitRoutes(cfg))
// 	// defer ts.Close()

// 	// Make a GET /user/signup request and then extract the CSRF token from the
// 	// response body.
// 	_, _, _ = ts.get(t, "/api/login")

// 	// csrfToken := extractCSRFToken(t, body)
// 	// Log the CSRF token value in our test output using the t.Logf() function.
// 	// The t.Logf() function works in the same way as fmt.Printf(), but writes
// 	// the provided message to the test output.

// 	tests := []struct {
// 		name     string
// 		nickname string
// 		password string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Valid login with nickname",
// 			nickname: existNeckName,
// 			password: validPassword,
// 			wantCode: http.StatusOK,
// 		},
// 		{
// 			name:     "valid login with email",
// 			nickname: existEmail,
// 			password: validPassword,
// 			wantCode: http.StatusOK,
// 		},
// 		{
// 			name:     "Empty Nickname",
// 			nickname: "",
// 			password: validPassword,
// 			wantCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:     "Empty Password",
// 			nickname: existNeckName,
// 			password: "",
// 			wantCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:     "Empty Neck Name And Password",
// 			nickname: "",
// 			password: "",
// 			wantCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:     "Wrgong nickname",
// 			nickname: "MadeUpNeckname",
// 			password: validPassword,
// 			wantCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:     "Wrong Password",
// 			nickname: existNeckName,
// 			password: "MadeUpPassword",
// 			wantCode: http.StatusBadRequest,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 				Nickname string `json:"nickname"`
// 				Password string `json:"password"`
// 			}{
// 				Nickname: tt.nickname,
// 				Password: tt.password,
// 			}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			code, _, _ := ts.postJSON(t, "/api/login", jsonBody)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }

// func TestUserPing(t *testing.T) {
// 	// Create the application struct containing our mocked dependencies and set
// 	// up the test server for running an end-to-end test.
// 	// app, cfg := server.NewTestApplication()
// 	// ts := newTestServer(t, app.InitRoutes(cfg))
// 	// defer ts.Close()

// 	// Make a GET /user/signup request and then extract the CSRF token from the
// 	// response body.
// 	_, _, _ = ts.get(t, "/api/login")

// 	// csrfToken := extractCSRFToken(t, body)
// 	// Log the CSRF token value in our test output using the t.Logf() function.
// 	// The t.Logf() function works in the same way as fmt.Printf(), but writes
// 	// the provided message to the test output.

// 	tests := []struct {
// 		name     string
// 		nickname string
// 		password string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Logout Test",
// 			nickname: existNeckName,
// 			password: validPassword,
// 			wantCode: http.StatusInternalServerError,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 				Nickname string `json:"nickname"`
// 				Password string `json:"password"`
// 			}{
// 				Nickname: tt.nickname,
// 				Password: tt.password,
// 			}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			code, header, _ := ts.postJSON(t, "/api/login", jsonBody)
// 			assert.Equal(t, code, tt.wantCode)

// 			// --- Logout Phase ---
// 			// Get session cookie from login response
// 			cookie := extractSessionCookie(header)
// 			if cookie == nil {
// 				t.Fatal("No session cookie found in login response")
// 			}
// 			// --- Logout Phase ---
// 			logoutReq, _ := http.NewRequest(http.MethodPost, ts.URL+"/ping/user", nil) // ✅ Fix: use full URL
// 			logoutReq.AddCookie(cookie)                                                // Attach session cookie
// 			logoutResp, err := ts.Client().Do(logoutReq)
// 			log.Println("ping request", logoutReq)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

//				// --- Validate Logout Response ---
//				log.Println("Logout Response Code:", logoutResp.StatusCode)
//				assert.Equal(t, logoutResp.StatusCode, tt.wantCode)
//			})
//		}
//	}
func TestUserLogout(t *testing.T) {
	tests := []struct {
		name     string
		nickname string
		password string
		wantCode int
	}{
		{
			name:     "Successful Logout",
			nickname: existNeckName,
			password: validPassword,
			wantCode: http.StatusOK, // Expect successful login
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// --- Login Phase ---
			loginBody := struct {
				Nickname string `json:"nickname"`
				Password string `json:"password"`
			}{
				Nickname: tt.nickname,
				Password: tt.password,
			}

			jsonBody, _ := json.Marshal(loginBody)
			loginCode, loginHeader, _ := ts.postJSON(t, "/api/login", jsonBody)
			assert.Equal(t, loginCode, tt.wantCode) // Ensure login succeeds

			// --- Extract Session Cookie ---
			cookie := extractSessionCookie(loginHeader)
			if cookie == nil {
				t.Fatal("No session cookie found")
			}
			// --- Logout Phase ---
			logoutReq, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/logout", nil)
			logoutReq.AddCookie(cookie)
			logoutResp, err := ts.Client().Do(logoutReq)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, logoutResp.StatusCode, tt.wantCode) // Check logout succeeded

			// --- Validate Session Invalidation ---
			// Try accessing a protected route (e.g., /ping/user) after logout
			protectedReq, _ := http.NewRequest(http.MethodGet, ts.URL+"/ping/user", nil)
			protectedReq.AddCookie(cookie) // Use the same (now invalid) cookie
			protectedResp, err := ts.Client().Do(protectedReq)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, protectedResp.StatusCode, http.StatusForbidden) // Expect 401
		})
	}
}

// Helper functions
// func TestPosts(t *testing.T) {
// 	// Create a new instance of our application struct which uses the mocked
// 	// dependencies.
// 	app, cfg := server.NewTestApplication()
// 	// Establish a new test server for running end-to-end tests.
// 	ts := newTestServer(t, app.InitRoutes(cfg))
// 	defer ts.Close()
// 	// Set up some table-driven tests to check the responses sent by our
// 	// application for different URLs.
// 	tests := []struct {
// 		name     string
// 		urlPath  string
// 		wantCode int
// 		wantBody string
// 	}{
// 		{
// 			name:     "Valid ID",
// 			urlPath:  "/snippet/view?id=1",
// 			wantCode: http.StatusOK,
// 			wantBody: "An old silent pond...",
// 		},
// 		{
// 			name:     "Non-existent ID",
// 			urlPath:  "/snippet/view?id=2",
// 			wantCode: http.StatusNotFound,
// 		},
// 		{
// 			name:     "Negative ID",
// 			urlPath:  "/snippet/view?id-1",
// 			wantCode: http.StatusNotFound,
// 		},
// 		{
// 			name:     "Decimal ID",
// 			urlPath:  "/snippet/view?id=1.23",
// 			wantCode: http.StatusNotFound,
// 		},
// 		{
// 			name:     "String ID",
// 			urlPath:  "/snippet/view?id=foo",
// 			wantCode: http.StatusNotFound,
// 		},
// 		{
// 			name:     "Empty ID",
// 			urlPath:  "/snippet/view/",
// 			wantCode: http.StatusNotFound,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			code, _, body := ts.get(t, tt.urlPath)
// 			fmt.Println(body)
// 			assert.Equal(t, code, tt.wantCode)
// 			if tt.wantBody != "" {
// 				assert.StringContains(t, body, tt.wantBody)
// 			}
// 		})
// 	}
// }
