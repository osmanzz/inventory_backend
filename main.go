package main

import (
	"Osman/backendapps/db"
	new "Osman/backendapps/init"
	"Osman/backendapps/resource/uiresource"
	"Osman/backendapps/resource/usecaseResource"
	"Osman/backendapps/server"
	usecases "Osman/backendapps/usecase"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	database:= new.NewDB()
	repo := db.GetOrderDB(database)
	usecase := usecases.GetUsecase(repo)
	UIResource := newUIResource(usecase)

	httpServer:= server.InitHttp(UIResource)
	httpServer.Run()
}

func newUIResource(usecase usecaseResource.UsecaseResource) uiresource.UIResource {
	return uiresource.UIResource{
		Usecase: usecase,
		Router: httprouter.New(),
		ServiceType: &usecases.HTTPServiceType,
	}
}
