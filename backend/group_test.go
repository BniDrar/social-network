package main

// func TestGetGroups(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Group Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		limit    int
// 		offset   int
// 		wantCode int
// 	}{
// 		{
// 			name:     "User Group",
// 			limit:    10,
// 			offset:   0,
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 				Limit  int `json:"limit"`
// 				Offset int `json:"offset"`
// 			}{
// 				Limit:  10,
// 				Offset: 0,
// 			}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			//code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodPost)
// 			code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }

// func TestGetGroupById(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For GetGroup By ID Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Get Group By ID",
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

// 			//code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodPost)
// 			code, _, _ := ts.JSONRequest(t, "/api/group?id=1", jsonBody, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }

// func TestCreateGroup(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For GetGroup By ID Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		wantCode int
// 	}{
// 		{
// 			name:     "Create Group",
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

// 			//code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodPost)
// 			code, _, _ := ts.JSONRequest(t, "/api/group/create", jsonBody, http.MethodPost)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }

// func TestGetAllGroups(t *testing.T) {

// 	/*_________________THE FIRST STEP IS TO LOGIN_______________*/
// 	t.Run("Loging For Get All Groups Test", ts.login)
// 	/*___________MAKE A TEST TABLE OF ALL POSSIBLE CASES_________*/
// 	tests := []struct {
// 		name     string
// 		limit    int
// 		offset   int
// 		wantCode int
// 	}{
// 		{
// 			name:     "User Group",
// 			limit:    10,
// 			offset:   0,
// 			wantCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody := struct {
// 				Limit  int `json:"limit"`
// 				Offset int `json:"offset"`
// 			}{
// 				Limit:  10,
// 				Offset: 0,
// 			}

// 			jsonBody, err := json.Marshal(reqBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			//code, _, _ := ts.JSONRequest(t, "/api/user_groups", jsonBody, http.MethodPost)
// 			code, _, _ := ts.JSONRequest(t, "/api/groups", jsonBody, http.MethodGet)
// 			assert.Equal(t, code, tt.wantCode)
// 		})
// 	}
// }
