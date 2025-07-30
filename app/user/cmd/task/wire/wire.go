//go:build wireinject
// +build wireinject

package wire

import (
	"simplex/app/user/internal/repo"
	"simplex/app/user/internal/srv"
	"simplex/app/user/internal/task"
	"simplex/pkg/app"
	"simplex/pkg/logx"
	"simplex/pkg/sid"
	"simplex/repository"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repo.NewUserRepository,
)

var taskSet = wire.NewSet(
	task.NewTask,
	task.NewUserTask,
)
var serverSet = wire.NewSet(
	srv.NewTaskServer,
)

// build App
func newApp(
	task *srv.TaskServer,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("user-task"),
	)
}

func NewWire(*viper.Viper, *logx.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		taskSet,
		serverSet,
		newApp,
		sid.NewSid,
	))
}
