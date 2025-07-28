//go:build wireinject
// +build wireinject

package wire

import (
	"simplex/app/user/internal/hdl"
	"simplex/app/user/internal/job"
	"simplex/app/user/internal/repo"
	"simplex/app/user/internal/srv"
	"simplex/app/user/internal/svc"
	"simplex/pkg/app"
	"simplex/pkg/jwt"
	"simplex/pkg/log"
	"simplex/pkg/server/http"
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
	httpServer *http.Server,
	jobServer *srv.JobServer,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, jobServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
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
