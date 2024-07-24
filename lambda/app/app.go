package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// initialize db store --> gets passed down into the API handler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewAPIHandler(db)

	return App {
		ApiHandler: apiHandler,
	}

}