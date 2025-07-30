package srv

import (
	apiV1 "simplex/app/user/api/v1"
	"simplex/app/user/internal/hdl"
	"simplex/app/user/internal/md"
	"simplex/pkg/jwt"
	"simplex/pkg/logx"
	"simplex/pkg/serv/http_serv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	logger *logx.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *hdl.UserHandler,
) *http_serv.Server {
	if conf.GetString("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	s := http_serv.NewServer(
		gin.Default(),
		logger,
		http_serv.WithServerHost(conf.GetString("http.host")),
		http_serv.WithServerPort(conf.GetInt("http.port")),
	)

	s.Use(
		md.CORSMiddleware(),
		md.ResponseLogMiddleware(logger),
		md.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(md.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(md.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
