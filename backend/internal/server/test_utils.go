package server

// New import

// func newTestApplication(t *testing.T) *App {
// 	// And a form decoder.
// 	// formDecoder := form.NewDecoder()
// 	// And a session manager instance. Note that we use the same settings as
// 	// production, except that we *don't* set a Store for the session manager.
// 	// If no store is set, the SCS package will default to using a transient
// 	// in-memory store, which is ideal for testing purposes.
// 	sessionManager := scs.New()
// 	sessionManager.Lifetime = 12 * time.Hour
// 	sessionManager.Cookie.Secure = true

// 	app := NewApp(nil, sessionManager)

// 	return app
// }

// Define a custom testServer type which embeds a httptest.Server instance.
// type testServer struct {
// 	*httptest.Server
// }

// // Implement a get() method on our custom testServer type. This makes a GET
// // request to a given url path using the test server client, and returns the
// // response status code, headers and body.
// func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
// 	rs, err := ts.Client().Get(ts.URL + urlPath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer rs.Body.Close()
// 	body, err := io.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	bytes.TrimSpace(body)
// 	return rs.StatusCode, rs.Header, string(body)
// }

// // Create a newTestServer helper which initalizes and returns a new instance
// // of our custom testServer type.
// func newTestServer(t *testing.T, h http.Handler) *testServer {
// 	// Initialize the test server as normal.
// 	ts := httptest.NewTLSServer(h)
// 	// Initialize a new cookie jar.
// 	jar, err := cookiejar.New(nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// Add the cookie jar to the test server client. Any response cookies will
// 	// now be stored and sent with subsequent requests when using this client.
// 	ts.Client().Jar = jar
// 	// Disable redirect-following for the test server client by setting a custom
// 	// CheckRedirect function. This function will be called whenever a 3xx
// 	// response is received by the client, and by always returning a
// 	// http.ErrUseLastResponse error it forces the client to immediately return
// 	// the received response.
// 	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
// 		return http.ErrUseLastResponse
// 	}
// 	return &testServer{ts}
// }

// // Create a postForm method for sending POST requests to the test server. The
// // final parameter to this method is a url.Values object which can contain any
// // form data that you want to send in the request body.
// func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, string) {
// 	rs, err := ts.Client().PostForm(ts.URL+urlPath, form)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// Read the response body from the test server.
// 	defer rs.Body.Close()
// 	body, err := io.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	bytes.TrimSpace(body)
// 	// Return the response status, headers and body.
// 	return rs.StatusCode, rs.Header, string(body)
// }
