package main

type Link struct {
	Contributor string
	Nature      string
	Title       string
	Link        string
}

type Issue struct {
	Id     int
	Status string
	Links  []Link
}

type Newsletter struct {
	Issues []Issue
}
