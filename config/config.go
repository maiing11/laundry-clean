package config

import "go.uber.org/fx"

// Module export depedency

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewDatabase),
	fx.Provide(NewRequestHandler),
)
