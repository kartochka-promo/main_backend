package admin_session

type UseCase interface {
	CreateSession(adminUsername string) (string, error)
	DeleteSession(cookieValue string) error
	CheckSession(cookieValue string) (string, error)
	GetUsernameByCookie(cookieValue string) (string, error)
}
