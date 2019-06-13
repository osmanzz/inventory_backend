package main

import (
	"Osman/github.com/inventory_backend/db"
	new "Osman/github.com/inventory_backend/init"
	"Osman/github.com/inventory_backend/resource/uiresource"
	"Osman/github.com/inventory_backend/resource/usecaseResource"
	"Osman/github.com/inventory_backend/server"
	usecases "Osman/github.com/inventory_backend/usecase"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	database := new.NewDB()
	repo := db.GetOrderDB(database)
	usecase := usecases.GetUsecase(repo)
	UIResource := newUIResource(usecase)

	httpServer := server.InitHttp(UIResource)
	httpServer.Run()
}

func newUIResource(usecase usecaseResource.UsecaseResource) uiresource.UIResource {
	return uiresource.UIResource{
		Usecase:     usecase,
		Router:      httprouter.New(),
		ServiceType: &usecases.HTTPServiceType,
	}
}
