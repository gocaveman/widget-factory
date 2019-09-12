//+build wireinject

package main

import "github.com/google/wire"

func Setup() (*MainStuff, error) {

	wire.Build(
		NewWidgetController,
		NewMainStuff,
		NewDBConnString,
		NewDBDriverName,
		NewWidgetRouter,
	)

	return &MainStuff{}, nil
}
