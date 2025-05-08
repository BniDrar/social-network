package main

// const commentAPI = "/ping"

// func TestComment(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Group Test", func(t *testing.T) {

// 		reqBody := struct {
// 			Nickname string `json:"nickname"`
// 			Password string `json:"password"`
// 		}{
// 			Nickname: existNeckName,
// 			Password: validPassword,
// 		}

// 		jsonBody, err := json.Marshal(reqBody)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		code, _, _ := ts.postJSON(t, "/api/login", jsonBody)
// 		assert.Equal(t, code, http.StatusOK)
// 	})
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		wantCode int
// 	}{
// 		{
// 			name:     "User Group",
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 			}{}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			code, _, _ := ts.postJSON(t, commentAPI, jsonBody)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }
