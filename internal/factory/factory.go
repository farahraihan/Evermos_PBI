package factory

import (
	"evermos_pbi/config"
	"evermos_pbi/internal/routes"
	"evermos_pbi/internal/utils"

	u_hnd "evermos_pbi/internal/features/users/handler"
	u_rep "evermos_pbi/internal/features/users/repository"
	u_srv "evermos_pbi/internal/features/users/service"

	s_hnd "evermos_pbi/internal/features/stores/handler"
	s_rep "evermos_pbi/internal/features/stores/repository"
	s_srv "evermos_pbi/internal/features/stores/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	db, _ := config.ConnectDB()

	pu := utils.NewPassUtil()
	jwt := utils.NewJwtUtility()
	cloud := utils.NewCloudinaryUtility()

	sq := s_rep.NewStoreQuery(db)
	ss := s_srv.NewStoreService(sq, cloud)
	sh := s_hnd.NewStoreHandler(ss, jwt)

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserServices(uq, pu, jwt, cloud, ss)
	uh := u_hnd.NewUserHandler(us, jwt)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, sh)

}
