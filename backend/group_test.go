package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"socialNetwork/pkg/assert"
)

func TestGetGroups(t *testing.T) {

	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
	t.Run("Loging For Group Test", func(t *testing.T) {

		reqBody := struct {
			Nickname string `json:"nickname"`
			Password string `json:"password"`
		}{
			Nickname: existNeckName,
			Password: validPassword,
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		code, _, _ := ts.postJSON(t, "/api/login", jsonBody)
		assert.Equal(t, code, http.StatusOK)
	})
	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
	tests := []struct {
		name     string
		limit    int
		offset   int
		wantCode int
	}{
		{
			name:     "User Group",
			limit:    10,
			offset:   10,
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := struct {
				Limit  int `json:"limit"`
				Offset int `json:"offset"`
			}{
				Limit:  10,
				Offset: 10,
			}

			jsonBody, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatal(err)
			}

			//code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodPost)
			code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodGet)
			assert.Equal(t, code, tt.wantCode)
		})
	}
}
