//go:build wireinject
// +build wireinject

package wire

import (
	"simplex/app/geoip/internal/repo"
	"simplex/app/geoip/internal/srv"
	"simplex/pkg/app"
	"simplex/pkg/log"
	"simplex/repository"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repo.NewUserRepository,
)
var serverSet = wire.NewSet(
	srv.NewMigrateServer,
)

// build App
func newApp(
	migrateServer *srv.MigrateServer,
) *app.App {
	return app.NewApp(
		app.WithServer(migrateServer),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serverSet,
		newApp,
	))
}
