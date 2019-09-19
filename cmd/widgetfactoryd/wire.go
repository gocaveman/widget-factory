//+build wireinject

package main

import "github.com/google/wire"
import "github.com/gocaveman/widget-factory/cmd/widgetfactoryd/store"

func Setup() (*MainStuff, error) {

	wire.Build(
		NewDBConn,
		store.NewStore,
		NewWidgetController,
		NewMainStuff,
		NewDBConnString,
		NewDBDriverName,
		NewWidgetRouter,
	)

	return &MainStuff{}, nil
}
