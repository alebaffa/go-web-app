package main

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"html/template"
)

type Controller struct {
	newsletter *mgo.Collection
}

func (c *Controller) ViewAllIssues(w http.ResponseWriter, r *http.Request) {

	newsletter := GetAllIssues(c)
	t, _ := template.ParseFiles("templates/list_issues.html")
	err := t.Execute(w, newsletter)
	if err != nil {
		panic(err)
	}
}