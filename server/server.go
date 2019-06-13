package server

import (
	"Osman/github.com/inventory_backend/resource/uiresource"
	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	ui       uiresource.UIResource
	router   *httprouter.Router
	handlers []*handler
}
