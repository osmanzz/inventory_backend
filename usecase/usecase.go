package usecase

import (
	"encoding/json"
	"errors"

	"Osman/backendapps/db"
	"Osman/backendapps/resource/usecaseResource"
	"net/http"
)

type HttpHandlerUseCase func(*UsecaseData, usecaseResource.UsecaseResource) (interface{}, error)

var (
	HTTPServiceType ServiceType = "HTTPServiceType"
)

type UsecaseData struct {
	ServiceType *ServiceType
	HTTPData    *http.Request
}
type ServiceType string

func GetUsecase(repo db.Repo) usecaseResource.UsecaseResource {
	return usecaseResource.UsecaseResource{
		Repo: repo,
	}
}

func (u *UsecaseData) Cast(target interface{}) error {
	return u.cast(target)
}

func (u *UsecaseData) cast(target interface{}) error {

	if u.ServiceType == nil {
		return errors.New("service type not found!")
	}

	switch *u.ServiceType {
	case HTTPServiceType :
		return u.castHTTPRequest(target)
	default:
		return errors.New("unimplemented service type for casting")
	}
}

func (u *UsecaseData) castHTTPRequest(target interface{}) error {
	if u.HTTPData == nil {
		return errors.New("request data is nil!")
	}

	if u.HTTPData.Body != nil {
		err := json.NewDecoder(u.HTTPData.Body).Decode(target)
		if err != nil {
			return err
		}
	}
	return nil

}
