package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"socialNetwork/pkg/assert"
)

func TestRegister(t *testing.T) {
	// t.Parallel() // Run sub-tests concurrently

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
		{
			name:        "Empty First Name",
			nickname:    validNeckName,
			email:       validEmail,
			password:    validPassword,
			first:       "",
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Empty LastName",
			nickname:    validNeckName,
			email:       validEmail,
			password:    validPassword,
			first:       validFirst,
			last:        "",
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Empty NickName",
			nickname:    "",
			email:       validEmail,
			password:    validPassword,
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Empty email",
			nickname:    validNeckName,
			email:       "",
			password:    validPassword,
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Empty password",
			nickname:    validNeckName,
			email:       validEmail,
			password:    "",
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Invalid email",
			nickname:    validNeckName,
			email:       InvalidEmail,
			password:    validPassword,
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Short password",
			nickname:    validNeckName,
			email:       validEmail,
			password:    "short",
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "password With No Digit",
			nickname:    validNeckName,
			email:       validEmail,
			password:    "passwordWithNoDigit",
			first:       validFirst,
			last:        validLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
		{
			name:        "Duplicate email",
			nickname:    existNeckName,
			email:       existEmail,
			password:    validPassword,
			first:       existFirst,
			last:        existLast,
			dateOfBirth: validDateOfBirth,
			wantCode:    http.StatusBadRequest,
		},
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

			code, _, _ := ts.postJSON(t, "/api/register", jsonBody)

			assert.Equal(t, code, tt.wantCode)
			if tt.nickname != existNeckName {
				err = app.User.DeleteUserByNickName(tt.nickname)
				if err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func TestLogin(t *testing.T) {

	tests := []struct {
		name     string
		nickname string
		password string
		wantCode int
	}{
		{
			name:     "Valid login with nickname",
			nickname: existNeckName,
			password: validPassword,
			wantCode: http.StatusOK,
		},
		{
			name:     "valid login with email",
			nickname: existEmail,
			password: validPassword,
			wantCode: http.StatusOK,
		},
		{
			name:     "Empty Nickname",
			nickname: "",
			password: validPassword,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Empty Password",
			nickname: existNeckName,
			password: "",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Empty Neck Name And Password",
			nickname: "",
			password: "",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Wrgong nickname",
			nickname: "MadeUpNeckname",
			password: validPassword,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Wrong Password",
			nickname: existNeckName,
			password: "MadeUpPassword",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := struct {
				Nickname string `json:"nickname"`
				Password string `json:"password"`
			}{
				Nickname: tt.nickname,
				Password: tt.password,
			}

			jsonBody, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatal(err)
			}

			code, _, _ := ts.postJSON(t, "/api/login", jsonBody)
			assert.Equal(t, code, tt.wantCode)
		})
	}
}

func TestUserPing(t *testing.T) {

	tests := []struct {
		name     string
		nickname string
		password string
		wantCode int
	}{
		{
			name:     "Logout User Ping",
			nickname: existNeckName,
			password: validPassword,
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := struct {
				Nickname string `json:"nickname"`
				Password string `json:"password"`
			}{
				Nickname: tt.nickname,
				Password: tt.password,
			}

			jsonBody, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatal(err)
			}

			code, _, _ := ts.postJSON(t, "/api/login", jsonBody)
			assert.Equal(t, code, tt.wantCode)

			Req, _ := http.NewRequest(http.MethodGet, ts.URL+"/ping/user", nil) // Use GET instead of POST

			Resp, err := ts.Client().Do(Req) // Send request

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, Resp.StatusCode, tt.wantCode)
		})
	}
}
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
			wantCode: http.StatusOK,
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
			loginCode, _, _ := ts.postJSON(t, "/api/login", jsonBody)
			assert.Equal(t, loginCode, tt.wantCode) // Ensure login succeeds

			// --- Logout Phase ---
			logoutReq, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/logout", nil)
			logoutResp, err := ts.Client().Do(logoutReq)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, logoutResp.StatusCode, tt.wantCode) // Check logout succeeded

			// --- Validate Session Invalidation ---
			// Try accessing a protected route (e.g., /ping/user) after logout
			protectedReq, _ := http.NewRequest(http.MethodGet, ts.URL+"/ping/user", nil)
			protectedResp, err := ts.Client().Do(protectedReq)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, protectedResp.StatusCode, http.StatusForbidden) // Expect 401
		})
	}
}

func TestProfile(t *testing.T) {

	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
	t.Run("Loging For GetGroup By ID Test", ts.login)
	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
	tests := []struct {
		name     string
		wantCode int
	}{
		{
			name:     "Get Profile",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, _ := ts.JSONRequest(t, "/api/profile?userid=1", nil, http.MethodGet)
			assert.Equal(t, code, tt.wantCode)
		})
	}
}

// func TestFollow(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Follow Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Get Follow",
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			code, _, _ := ts.JSONRequest(t, "/api/follow?followed?=2", nil, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }

// func TestFollows(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Follow Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Get Follow",
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			code, _, _ := ts.JSONRequest(t, "/api/followers", nil, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }
