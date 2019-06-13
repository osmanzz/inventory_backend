package main

import (
	"github.com/osmanzz/inventory_backend/db"
	new "github.com/osmanzz/inventory_backend/init"
	"github.com/osmanzz/inventory_backend/resource/uiresource"
	"github.com/osmanzz/inventory_backend/resource/usecaseResource"
	"github.com/osmanzz/inventory_backend/server"
	usecases "github.com/osmanzz/inventory_backend/usecase"
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
