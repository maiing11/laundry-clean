package repository

import "go.uber.org/fx"

// Module export depedency
var Module = fx.Options(
	fx.Provide(NewRepository),
)
