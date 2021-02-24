package main

import (
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/graph"
)

func main() {
	// Config
	config.Init(".")

	// Facades
	departments.Init()

	// GraphQL
	graph.Init()
}
