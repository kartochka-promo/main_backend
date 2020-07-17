package admin_middleware

import (
	"net/http"

	"github.com/labstack/echo"

	sessions "2020_1_drop_table/internal/app/admin_session"
	"2020_1_drop_table/internal/pkg/responses"
)

type MiddlewareHandler struct {
	sessChecker sessions.UseCase
}

func (mh MiddlewareHandler) SetMiddleware(server *echo.Echo) {
	server.Use(mh.SetCorsMiddleware)
	server.Use(mh.PanicMiddleWare)
	server.Use(mh.CheckAuthentication)
}

func (mh *MiddlewareHandler) SetCorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-Csrf-Token, csrf-token, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Vary", "Cookie")
		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}

		return next(c)
	}
}

func (mh MiddlewareHandler) PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() error {
			if err := recover(); err != nil {
				return c.JSON(http.StatusInternalServerError, responses.HttpError{
					Code:    http.StatusInternalServerError,
					Message: "server panic",
				})
			}
			return nil
		}()
		return next(c)
	}
}

func (mh *MiddlewareHandler) CheckAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(rwContext echo.Context) error {
		var (
			err      error
			userName string
			cookie   *http.Cookie
		)

		if rwContext.Path() == "login" {
			return next(rwContext)
		}

		if cookie, err = rwContext.Cookie("cookie_value"); err != nil || cookie.Path != "/admin" {
			return rwContext.NoContent(http.StatusExpectationFailed)
		}

		if userName, err = mh.sessChecker.CheckSession(cookie.Value); err != nil || userName == "" {
			return rwContext.NoContent(http.StatusUnauthorized)
		}
		rwContext.Set("username", userName)
		return next(rwContext)
	}

}
