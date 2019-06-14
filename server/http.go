package server

import (
	"github.com/osmanzz/inventory_backend/resource/uiresource"
	"github.com/osmanzz/inventory_backend/service/api"
	"github.com/osmanzz/inventory_backend/usecase"
	"github.com/osmanzz/inventory_backend/usecase/handle"
	"github.com/rs/cors"
	"net/http"
)

type httpHandleFunc func(req *http.Request, writer http.ResponseWriter, uires uiresource.UIResource,usecase usecase.HttpHandlerUseCase) (interface{}, error)
type HandlerErrFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

func InitHttp(uires uiresource.UIResource) httpServer {
	return httpServer{
		ui : uires,
		router:   uires.Router,
		handlers: make([]*handler, 0),
	}
}

type handler struct {
	method          string
	path            string
	httpHandlerFunc httpHandleFunc
	usecase         usecase.HttpHandlerUseCase
}

func (h *httpServer) Run() {
	h.RegisterRestfulAPI()
	h.route()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	handler := c.Handler(h.router)
	http.ListenAndServe(":8081", handler)
}

func (h *httpServer) addHandler(method, path string, handleFunc httpHandleFunc, usecase usecase.HttpHandlerUseCase) {
	hand := &handler{
		method:          method,
		path:            path,
		httpHandlerFunc: handleFunc,
		usecase:         usecase,
	}
	h.handlers = append(h.handlers, hand)

}
func doFilter(hr HandlerErrFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		hr(writer, req)
	}
}
func (h *httpServer) route() {
	for i := range h.handlers {
		hand := h.handlers[i]
		h.router.HandlerFunc(hand.method, hand.path, doFilter(func(writer http.ResponseWriter, req *http.Request) (interface{}, error) {
			return hand.httpHandlerFunc(req, writer,h.ui, hand.usecase)
		}))
	}
}

func (h *httpServer) RegisterRestfulAPI() {
	h.addHandler("POST", "/user", api.SelecUserHandler, handle.SelectAllUserHandle)
	h.addHandler("POST","/login",api.LoginHandler,handle.LoginHandle)

}
