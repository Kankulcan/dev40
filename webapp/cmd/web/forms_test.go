package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("no data yet")

	if has{
		t.Error("form shows has field it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")

	if !has{
		t.Error("shows form does not have fiels when it should")
	}
}

func TestForm_Required(t *testing.T){
	r := httptest.NewRequest("POST", "/seila", nil)
	form := NewForm(r.PostForm)

	form.Required("a","b","c")
	if form.Valid() {
		t.Error("form show valid when required fiels are missing")
	}

	postedDate := url.Values{}
	postedDate.Add("a","a")
	postedDate.Add("b","b")
	postedDate.Add("c","c")

	r, _ = http.NewRequest("POST", "/seila", nil)
	r.PostForm = postedDate

	form = NewForm(r.PostForm)
	form.Required("a","b","c")
	if !form.Valid() {
		t.Error("shows post does not have required fields, when it does")
	}
}