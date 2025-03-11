package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"socialNetwork/internal/server"
	"socialNetwork/pkg/assert"
)

func TestPing(t *testing.T) {
	app, cfg := server.NewTestApplication()
	ts := newTestServer(t, app.InitRoutes(cfg)) // http.Server{}
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

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

func TestRegister(t *testing.T) {
	defer DropTestDB()
	// Create the application struct containing our mocked dependencies and set
	// up the test server for running an end-to-end test.
	app, cfg := server.NewTestApplication()
	ts := newTestServer(t, app.InitRoutes(cfg))
	defer ts.Close()

	// Make a GET /user/signup request and then extract the CSRF token from the
	// response body.
	_, _, _ = ts.get(t, "/api/register")

	// csrfToken := extractCSRFToken(t, body)
	// Log the CSRF token value in our test output using the t.Logf() function.
	// The t.Logf() function works in the same way as fmt.Printf(), but writes
	// the provided message to the test output.
	const (
		validNeckName    = "nickname"
		validFirst       = "firstname"
		validLast        = "lastname"
		validDateOfBirth = "01/01/2000"
		validName        = "chiwahed"
		validPassword    = "validPa$$word1"
		validEmail       = "hadak@example.com"
	)

	tests := []struct {
		name        string
		nickname    string
		email       string
		password    string
		first       string
		last        string
		dateOfBirth string
		wantCode    int
	}{
		{
			name:        "Valid submission",
			nickname:    validNeckName,
			email:       validEmail,
			password:    validPassword,
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusCreated,
		},
		// {
		// 	name:         "Empty name",
		// 	userName:     "",
		// 	userEmail:    validEmail,
		// 	userPassword: validPassword,
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
		// {
		// 	name:         "Empty email",
		// 	userName:     validName,
		// 	userEmail:    "",
		// 	userPassword: validPassword,
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
		// {
		// 	name:         "Empty password",
		// 	userName:     validName,
		// 	userEmail:    validEmail,
		// 	userPassword: "",
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
		// {
		// 	name:         "Invalid email",
		// 	userName:     validName,
		// 	userEmail:    "bob@example.",
		// 	userPassword: validPassword,
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
		// {
		// 	name:         "Short password",
		// 	userName:     validName,
		// 	userEmail:    validEmail,
		// 	userPassword: "pa$$",
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
		// {
		// 	name:         "Duplicate email",
		// 	userName:     validName,
		// 	userEmail:    "dupe@example.com",
		// 	userPassword: validPassword,
		// 	wantCode:     http.StatusUnprocessableEntity,
		// 	wantFormTag:  formTag,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := struct {
				Nickname    string `json:"nickname"`
				Email       string `json:"email"`
				Password    string `json:"password"`
				First       string `json:"first"`
				Last        string `json:"last"`
				DateOfBirth string `json:"date_of_birth"`
			}{
				Nickname:    tt.nickname,
				Email:       tt.email,
				Password:    tt.password,
				First:       tt.first,
				Last:        tt.last,
				DateOfBirth: tt.dateOfBirth,
			}

			jsonBody, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatal(err)
			}

			code, _, body := ts.postJSON(t, "/api/register", jsonBody)

			assert.Equal(t, code, tt.wantCode)
			// if tt.wantFormTag != "" {
			// 	assert.StringContains(t, body, tt.wantFormTag)
			// }
			fmt.Println("the response body:", body)
		})
	}
}
