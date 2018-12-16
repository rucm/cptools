package atcoder

import (
	infrastracture "github.com/rucm/cptools/atcoder/infrastracture"
	interactor "github.com/rucm/cptools/atcoder/interactor"
	repository "github.com/rucm/cptools/atcoder/repository"
	"github.com/rucm/cptools/usecase"
)

// AuthenticationController :
type AuthenticationController struct {
	Interactor usecase.Authentication
}

// NewAuthenticationController :
func NewAuthenticationController() *AuthenticationController {

	return &AuthenticationController{
		Interactor: &interactor.AuthenticationInteractor{
			Handler: &infrastracture.RequestHandler{
				URL:         "https://practice.contest.atcoder.jp/login",
				Method:      "POST",
				ContentType: "application/x-www-form-urlencoded",
			},
			Repo: &repository.SessionFileRepository{
				Filaname: "atcoder.session",
			},
		},
	}
}

// Login :
func (controller *AuthenticationController) Login(user string, password string) {

	param := infrastracture.NewParameter()
	param.Set("name", user)
	param.Set("password", password)
	controller.Interactor.Login(param)
}

// Logout :
func (controller *AuthenticationController) Logout() {

	controller.Interactor.Logout()
}
