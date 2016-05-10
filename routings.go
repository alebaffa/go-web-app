package main

import (
	"net/http"

	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

type Controller struct {
	newsletter *mgo.Collection
}

func (c *Controller) ViewAllIssues(res http.ResponseWriter, req *http.Request) {
	//var issues = make([]Issue, 0)
	newsletter := GetAllIssues(c)

	r := render.New(render.Options{
		IndentJSON: true,
	})

	r.JSON(res, 200, newsletter.Issues)
}
