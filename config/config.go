package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewConfig),
	fx.Provide(NewInfra),
)
