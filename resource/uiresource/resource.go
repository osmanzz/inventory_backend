package uiresource

import (
	cons_http "Osman/github.com/inventory_backend/cons/http"
	"Osman/github.com/inventory_backend/resource/usecaseResource"
	"Osman/github.com/inventory_backend/usecase"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UIResource struct {
	Usecase     usecaseResource.UsecaseResource
	Router      *httprouter.Router
	ServiceType *usecase.ServiceType
}

func (c *UIResource) RenderJSONDefault(w http.ResponseWriter, data interface{}, err error) {
	var header Header
	status := cons_http.StatusOK
	header.Status = status
	if err != nil {
		status := cons_http.StatusInternalServerError
		header.Status = status
		header.Reason = err.Error()
	}

	resp := APIResponse{
		Header: header,
		Data:   data,
	}
	byteData, _ := json.Marshal(resp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(byteData)
}

func (c *UIResource) CreateUseCaseData() *usecase.UsecaseData {
	return &usecase.UsecaseData{
		ServiceType: c.ServiceType,
	}
}
