package routes

import (
	"donasitamanzakattest/app/controllers"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/config"
	"donasitamanzakattest/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Route struct {
	startTime  time.Time
	appVersion string
	cfg        *config.Config
	repo       *repositories.RepositoryContext
	router     *gin.RouterGroup
	auth       *middleware.Auth
	ctrl       *controllers.Controller
}

func NewRoute(startTime time.Time, appVersion string, cfg *config.Config, repo *repositories.RepositoryContext, router *gin.RouterGroup) *Route {
	return &Route{startTime: startTime, appVersion: appVersion, cfg: cfg, repo: repo, router: router}
}

func (r *Route) RegisterCoreServicesRoutes() {
	ctrl := controllers.NewController(r.startTime, r.appVersion, r.cfg, r.repo)
	log.Warn().Msg("running route ....")

	auth := middleware.NewAuth(r.cfg, r.repo)

	if r.auth == nil {
		r.auth = auth
	}

	if r.ctrl == nil {
		r.ctrl = ctrl
	}

	r.router.GET("", ctrl.HealthController)
	r.initRoute()
}

func (r *Route) initRoute() {
	categories := r.router.Group("/categories")
	{
		categories.POST("", r.ctrl.CreateCategory)
		categories.GET("", r.ctrl.ListCategories)
		categories.GET("/:id", r.ctrl.GetCategory)
		categories.PUT("/:id", r.ctrl.UpdateCategory)
		categories.DELETE("/:id", r.ctrl.DeleteCategory)
	}

	programs := r.router.Group("/programs")
	{
		programs.POST("", r.ctrl.CreateProgram)
		programs.GET("", r.ctrl.ListPrograms)
		programs.GET("/:id", r.ctrl.GetProgram)
		programs.PUT("/:id", r.ctrl.UpdateProgram)
		programs.DELETE("/:id", r.ctrl.DeleteProgram)
	}

	news := r.router.Group("/news")
	{
		news.POST("", r.ctrl.CreateNews)
		news.GET("", r.ctrl.ListNews)
		news.GET("/:id", r.ctrl.GetNews)
		news.PUT("/:id", r.ctrl.UpdateNews)
		news.DELETE("/:id", r.ctrl.DeleteNews)
	}

	r.public()
}

func (r *Route) public() {
	public := r.router.Group("/public")
	r.categoriesPublic(public)
	r.authPublic(public)
}

func (r *Route) categoriesPublic(public *gin.RouterGroup) {
	{
		categories := public.Group("/categories")
		categories.GET("", r.ctrl.PublicListCategoriesController)
		//categories.GET("/:id", r.ctrl.PublicGetCategory)
	}
}

func (r *Route) authPublic(public *gin.RouterGroup) {
	{
		auth := public.Group("/auth")
		auth.POST("/registration", r.ctrl.RegistrationController)
		auth.POST("/login", r.ctrl.LoginController)
	}
}
