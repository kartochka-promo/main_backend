package delivery

import "github.com/labstack/echo"

func (as AdminService) Setup(server *echo.Echo) {
	/* SVA */
	server.POST("/api/v1/login",as.CreateAdmin)
	server.PUT("/api/v1/login",as.Authentication)
	server.DELETE("/api/v1/login",as.Logout)
	/* only updating admin pass */
	server.PUT("/api/v1/admin",as.UpdateAdmin)
}
