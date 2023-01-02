package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

// slice of struct
var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"general", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"rs", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"saj", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"mrp", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "john"},
		{key: "last_name", value: "sinha"},
		{key: "email", value: "sinha@test.com"},
		{key: "phone", value: "7895998998"},
	}, http.StatusOK},
}

func TestNewHandler(t *testing.T) {
	routes := geRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			response, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s , expected %d but ogd %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		} else {

			values := url.Values{}

			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			response, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if response.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s , expected %d but ogd %d", e.name, e.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
