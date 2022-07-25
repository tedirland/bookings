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
	if !isValid {
		t.Error("got invalid, should have been valid")
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/myendpoint", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("error saying required fields are not present when they are")
	}

}

// Has checks if form fields is in post and not empty

func TestFormHas(t *testing.T) {
	// need to create a rquest and a postform
	postedData := url.Values{}
	form := New(postedData)
	has := form.Has("a")
	if has {
		t.Error("Reported field present that wasn't")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	// reinit form with new request
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/myendpoint", nil)
	form := New(r.PostForm)

	form.MinLength("whatever", 5)
	if form.Valid() {
		t.Error("Minlength passed on a field that does not exist")
	}

	isError := form.Errors.Get("whatever")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("a", "hello")
	// reinit form with new request
	form = New(postedData)

	form.MinLength("a", 100)
	if form.Valid() {
		t.Error("MinLength reported as too short but it is correct length")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc")

	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("some_field")
	if form.Valid() {
		t.Error("Form shows valid email on a field that does not exist")
	}

	postedData = url.Values{}
	postedData.Add("a", "ted")
	// reinit form with new request
	form = New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("Form shows valid email on a field that is not an email")
	}
	postedData = url.Values{}
	postedData.Add("a", "ted@ted.com")
	// reinit form with new request
	form = New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("IsEmail failed on a value that was an email")
	}

}
