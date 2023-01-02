package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if isValid == false {
		t.Error("got invalid when should have valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shols valid when required fields missing")
	}

	postData := url.Values{}

	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postData

	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required fileds when it does")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("watever")

	if has {
		t.Error("forms shows has when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.Has("a")
	if has == false {
		t.Error("shows form doesnt not have fileds when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)

	if form.Valid() {
		t.Error("shows form min length for non existence field")
	}

	isError := form.Errors.Get("x")

	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("fi", "some")

	form = New(postedValues)

	form.MinLength("fi", 100)
	if form.Valid() {
		t.Error("shows min len of 100 when data is shorter")
	}

	postedValues = url.Values{}

	postedValues.Add("mm", "abc123")
	form = New(postedValues)
	form.MinLength("mm", 1)

	if !form.Valid() {
		t.Error("shows min len of 1 is not when it is")
	}

	isError = form.Errors.Get("mm")
	if isError != "" {
		t.Error("should not have error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValue := url.Values{}
	form := New(postedValue)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("form shows valid email for non existent filed")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "me@He.com")
	form = New(postedValue)

	form.IsEmail("email")

	if form.Valid() == false {
		t.Error("got an invalid email when we should not have")
	}

	postedValue = url.Values{}
	postedValue.Add("x", "me.com")
	form = New(postedValue)

	form.IsEmail("x")

	if form.Valid() == true {
		t.Error("got  valid for  invalid email when we should not have")
	}

}
