package srv

import (
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

	v1 := s.Group("/v1")
	{
		// Non-strict permission routing group
		userRouter := v1.Group("/").Use(md.NoStrictAuth(jwt, logger))
		{
			userRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(md.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
