package admin_session

type Repository interface {
	AddCookie(cookieValue, adminUsername string) error
	GetUsernameByCookie(cookieValue string) (string, error)
	DeleteCookie(cookieValue string) error
}
