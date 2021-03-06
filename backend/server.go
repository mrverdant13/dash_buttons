package main

import (
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/facades/districts"
	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
	"github.com/mrverdant13/dash_buttons/backend/graph"
	"github.com/mrverdant13/dash_buttons/backend/internal/database"
)

func main() {
	// Config
	config.Init(".")

	// Internal
	database.Init(database.OrmMigration)

	// Facades
	auth.Init()
	departments.Init()
	provinces.Init()
	districts.Init()
	users.Init(true)

	// GraphQL
	graph.Init()
}
