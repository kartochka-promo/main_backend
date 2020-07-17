package usecase

import (
	"github.com/satori/go.uuid"

	interfaces "2020_1_drop_table/internal/app/admin_session"
)

type CookieLogic struct {
	cookieStorage interfaces.Repository
}

func (cl *CookieLogic) CreateSession(adminUsername string) (string, error) {
	cookieValue := uuid.NewV4().String()
	return cookieValue, cl.cookieStorage.AddCookie(cookieValue, adminUsername)
}

func (cl *CookieLogic) DeleteSession(cookieValue string) error {
	return cl.cookieStorage.DeleteCookie(cookieValue)
}

func (cl *CookieLogic) CheckSession(cookieValue string) (string, error) {
	return cl.cookieStorage.GetUsernameByCookie(cookieValue)
}

func (cl *CookieLogic) GetUsernameByCookie(cookieValue string) (string, error) {
	return cl.cookieStorage.GetUsernameByCookie(cookieValue)
}
