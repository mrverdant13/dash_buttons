package main

import (
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
	"github.com/mrverdant13/dash_buttons/backend/graph"
	"github.com/mrverdant13/dash_buttons/backend/internal/pkg/database"
)

func main() {
	// Config
	config.Init(".")

	// Internal
	database.Init(true)

	// Facades
	auth.Init()
	departments.Init()
	provinces.Init()
	users.Init()

	// GraphQL
	graph.Init()
}
