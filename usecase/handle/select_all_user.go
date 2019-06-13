package handle

import (
	"github.com/osmanzz/inventory_backend/resource/usecaseResource"
	"github.com/osmanzz/inventory_backend/usecase"
)

func SelectAllUserHandle(ucData *usecase.UsecaseData, resource usecaseResource.UsecaseResource) (interface{}, error) {

	data, err := resource.Repo.SelectUser()
	if err != nil {
		return nil, err
	}

	return data, nil
}
