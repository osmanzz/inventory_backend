package server

import (
	"Osman/backendapps/resource/uiresource"
	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	ui       uiresource.UIResource
	router   *httprouter.Router
	handlers []*handler
}
