package main

import (
    "html"
    "io"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/http/httptest"
    "net/url"
    "regexp"
    "testing"
    "time"


    "cci007/snippetbox/pkg/models/mock"

    "github.com/golangcollege/sessions"

)




var csrfTokenRX = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)


func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, []byte) {
    rs, err := ts.Client().PostForm(ts.URL+urlPath, form)
    if err != nil {
        t.Fatal(err)
    }

    // Read the response body.
    defer rs.Body.Close()
    body, err := io.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }

    // Return the response status, headers and body.
    return rs.StatusCode, rs.Header, body
}




func extractCSRFToken(t *testing.T, body []byte) string {
    // Use the FindSubmatch method to extract the token from the HTML body.
    // Note that this returns an array with the entire matched pattern in the
    // first position, and the values of any captured data in the subsequent
    // positions.
    matches := csrfTokenRX.FindSubmatch(body)
    if len(matches) < 2 {
        t.Fatal("no csrf token found in body")
    }

    return html.UnescapeString(string(matches[1]))
}

// Create a newTestApplication helper which returns an instance of our
// application struct containing mocked dependencies.
func newTestApplication(t *testing.T) *application {
    // Create an instance of the template cache.
    templateCache, err := newTemplateCache("./../../ui/html/")
    if err != nil {
        t.Fatal(err)
    }

    // Create a session manager instance, with the same settings as production.
    session := sessions.New([]byte("3dSm5MnygFHh7XidAtbskXrjbwfoJcbJ"))
    session.Lifetime = 12 * time.Hour
    session.Secure = true


    return &application{
        errorLog: log.New(io.Discard, "", 0),
        infoLog:  log.New(io.Discard, "", 0),
        session:       session,
        snippets:      &mock.SnippetModel{},
        templateCache: templateCache,
        users:         &mock.UserModel{},
    }
}

type testServer struct {
    *httptest.Server
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
    rs, err := ts.Client().Get(ts.URL + urlPath)
    if err != nil {
        t.Fatal(err)
    }

    defer rs.Body.Close()
    body, err := io.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }

    return rs.StatusCode, rs.Header, body
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
    ts := httptest.NewTLSServer(h)

    jar, err := cookiejar.New(nil)
    if err != nil {
        t.Fatal(err)
    }
    ts.Client().Jar = jar

    ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
    }

    return &testServer{ts}
}

