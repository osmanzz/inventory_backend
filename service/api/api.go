package api

import (
	"github.com/osmanzz/inventory_backend/resource/uiresource"
	"github.com/osmanzz/inventory_backend/usecase"
	"net/http"
)

func SelecUserHandler(req *http.Request, writer http.ResponseWriter, uires uiresource.UIResource,useCase usecase.HttpHandlerUseCase) (interface{},error){

	data := uires.CreateUseCaseData()
	resp,err := useCase(data,uires.Usecase)
	uires.RenderJSONDefault(writer,resp,err)
	return resp,nil
}

func LoginHandler(req *http.Request, writer http.ResponseWriter, uires uiresource.UIResource,useCase usecase.HttpHandlerUseCase) (interface{},error) {


	data := uires.CreateUseCaseData()
	data.HTTPData = req
	resp,err := useCase(data,uires.Usecase)
	uires.RenderJSONDefault(writer,resp,err)
	return resp,nil
}
