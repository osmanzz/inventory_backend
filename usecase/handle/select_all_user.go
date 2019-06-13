package handle

import (
	"Osman/backendapps/resource/usecaseResource"
	"Osman/backendapps/usecase"
)

func SelectAllUserHandle(ucData *usecase.UsecaseData,resource usecaseResource.UsecaseResource) (interface{},error){

	data,err  := resource.Repo.SelectUser()
	if err != nil {
		return nil,err
	}

	return data , nil
}