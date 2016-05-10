package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllIssues(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	Controller.ViewAllIssues(res, req)
}
