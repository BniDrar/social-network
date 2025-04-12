package main

// func TestGetPosts(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Posts Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		limit    int
// 		offset   int
// 		wantCode int
// 	}{
// 		{
// 			name:     "User Posts",
// 			limit:    10,
// 			offset:   0,
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			query := "?limit=10&offset=0"
// 			url := "/api/posts" + query

// 			code, _, _ := ts.JSONRequest(t, url, nil, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})

// 	}
// }

// const postAPI = "/ping"

// func TestPost(t *testing.T) {

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

// 			code, _, _ := ts.postJSON(t, postAPI, jsonBody)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }
