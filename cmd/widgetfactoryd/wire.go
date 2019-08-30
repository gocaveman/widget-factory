//+build wireinject

package main

import "github.com/google/wire"

func Setup() (*MainStuff, error) {

	wire.Build(
		NewMainStuff,
		NewDBConnString,
		NewDBDriverName,
	)

	return &MainStuff{}, nil
}
