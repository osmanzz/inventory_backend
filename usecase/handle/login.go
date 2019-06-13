package handle

import (
	"Osman/backendapps/resource/usecaseResource"
	"Osman/backendapps/usecase"
)
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Success int `json:"success"`
	Message string `json:"message"`
}
func LoginHandle(ucData *usecase.UsecaseData,resource usecaseResource.UsecaseResource) (interface{},error){

	request := new(LoginRequest)
	response:= new(LoginResponse)
	err := ucData.Cast(request)

	if err != nil {
		response.Success = 0
		response.Message = err.Error()
		return response, err
	}

	data,err  := resource.Repo.SelectUserByUsername(request.Username,request.Password)
	if err != nil {
		response.Success = 0
		response.Message = err.Error()
		return response,err
	}

	if data != nil{
		response.Success = 1;
		response.Message = "Success"
	}
	return response , nil
}