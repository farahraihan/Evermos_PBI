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

	a_hnd "evermos_pbi/internal/features/address/handler"
	a_rep "evermos_pbi/internal/features/address/repository"
	a_srv "evermos_pbi/internal/features/address/service"

	c_hnd "evermos_pbi/internal/features/categories/handler"
	c_rep "evermos_pbi/internal/features/categories/repository"
	c_srv "evermos_pbi/internal/features/categories/service"

	p_hnd "evermos_pbi/internal/features/products/handler"
	p_rep "evermos_pbi/internal/features/products/repository"
	p_srv "evermos_pbi/internal/features/products/service"

	l_hnd "evermos_pbi/internal/features/logproduct/handler"
	l_rep "evermos_pbi/internal/features/logproduct/repository"
	l_srv "evermos_pbi/internal/features/logproduct/service"

	t_hnd "evermos_pbi/internal/features/transaction/handler"
	t_rep "evermos_pbi/internal/features/transaction/repository"
	t_srv "evermos_pbi/internal/features/transaction/service"

	d_rep "evermos_pbi/internal/features/detailtransaction/repository"
	d_srv "evermos_pbi/internal/features/detailtransaction/service"

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

	aq := a_rep.NewAddressQuery(db)
	as := a_srv.NewAdreessService(aq)
	ah := a_hnd.NewAddressHandler(as, jwt)

	cq := c_rep.NewCategoryQuery(db)
	cs := c_srv.NewCategoryService(cq, us)
	ch := c_hnd.NewCategoryHandler(cs, jwt)

	lq := l_rep.NewLogProductQuery(db)
	ls := l_srv.NewLogProductService(lq)
	lh := l_hnd.NewLogProductHandler(ls)

	pq := p_rep.NewProductQuery(db)
	ps := p_srv.NewProductService(pq, cloud, ss, ls)
	ph := p_hnd.NewProductHandler(ps, jwt)

	dq := d_rep.NewDetailTransactionQuery(db)
	ds := d_srv.NewDetailTransactionService(dq)

	tq := t_rep.NewTransactionQuery(db)
	ts := t_srv.NewTransactionService(tq, ds)
	th := t_hnd.NewTransactionHandler(ts, jwt)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, sh, ah, ch, lh, ph, th)

}
