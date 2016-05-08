package main

import "gopkg.in/mgo.v2/bson"

func GetAllIssues(c *Controller) Newsletter {
	newsletter := Newsletter{}

	if err := c.newsletter.Find(bson.M{}).Select(bson.M{"issues": 1}).One(&newsletter); err != nil {
		return newsletter
	}
	return newsletter
}
