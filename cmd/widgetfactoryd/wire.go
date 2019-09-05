//+build wireinject

package main

import "github.com/google/wire"

func Setup() (*MainStuff, error) {

	wire.Build(
		NewSampleController(),
		NewMainStuff,
		NewDBConnString,
		NewDBDriverName,
		NewRouter(),
	)

	return &MainStuff{}, nil
}
