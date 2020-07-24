package delivery

import "github.com/labstack/echo"

/* SVA For admin + main admin operations */
func (ams AdminMainService) Setup(server *echo.Echo) {
	/* SVA */
	server.POST("/api/v1/admin/login", ams.CreateAdmin)
	server.PUT("/api/v1/admin/login", ams.Authentication)
	server.DELETE("/api/v1/admin/login", ams.Logout)
	///* only updating admin password */
	server.PUT("/api/v1/admin/update", ams.UpdateAdmin)
}

/* main admin operations with cafe */
func (acs AdminCafeService) Setup(server *echo.Echo) {
	/* Singe Cafe Methods */
	server.POST("/api/v1/admin/database/cafe", acs.CreateCafe)
	server.GET("/api/v1/admin/database/cafe/:cafeID", acs.GetCafe)
	server.PUT("/api/v1/admin/database/cafe/:cafeID", acs.UpdateCafe)
	server.DELETE("/api/v1/admin/database/cafe/:cafeID", acs.DeleteCafe)
	/* Get All Cafes Methods */
	server.GET("/api/v1/admin/database/cafes:limit:offset", acs.GetCafes)
}
