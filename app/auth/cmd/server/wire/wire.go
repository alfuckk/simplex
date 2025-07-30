//go:build wireinject
// +build wireinject

package wire

import (
	"simplex/app/auth/internal/hdl"
	"simplex/app/auth/internal/job"
	"simplex/app/auth/internal/repo"
	"simplex/app/auth/internal/srv"
	"simplex/app/auth/internal/svc"
	"simplex/pkg/app"
	"simplex/pkg/jwt"
	"simplex/pkg/logx"
	"simplex/pkg/serv/http_serv"
	"simplex/pkg/sid"
	"simplex/repository"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	//repository.NewMongo,
	repository.NewRepository,
	repository.NewTransaction,
	repo.NewUserRepository,
)

var serviceSet = wire.NewSet(
	svc.NewService,
	svc.NewUserService,
)

var handlerSet = wire.NewSet(
	hdl.NewHandler,
	hdl.NewUserHandler,
)

var jobSet = wire.NewSet(
	job.NewJob,
	job.NewUserJob,
)
var serverSet = wire.NewSet(
	srv.NewHTTPServer,
	srv.NewJobServer,
)

// build App
func newApp(
	httpServer *http_serv.Server,
	jobServer *srv.JobServer,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, jobServer),
		app.WithName("user-server"),
	)
}

func NewWire(*viper.Viper, *logx.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		jobSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
