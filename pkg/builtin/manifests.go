package builtin

import "github.com/kairos-development/kairos-contracts/pkg/contracts"

var (
	// GridManifest is the built-in grid strategy descriptor.
	GridManifest = contracts.PluginManifest{
		Name:         "grid",
		Strategy:     "grid",
		ABIVersion:   contracts.ABIVersion{Major: 1, Minor: 0},
		StrategyMode: contracts.StrategyModeTrade,
		Entrypoint:   "on_tick",
	}
	// DCAManifest is the built-in DCA strategy descriptor.
	DCAManifest = contracts.PluginManifest{
		Name:         "dca",
		Strategy:     "dca",
		ABIVersion:   contracts.ABIVersion{Major: 1, Minor: 0},
		StrategyMode: contracts.StrategyModeTrade,
		Entrypoint:   "on_tick",
	}
	// TrendFollowingManifest is the built-in trend strategy descriptor.
	TrendFollowingManifest = contracts.PluginManifest{
		Name:         "trend_following",
		Strategy:     "trend_following",
		ABIVersion:   contracts.ABIVersion{Major: 1, Minor: 0},
		StrategyMode: contracts.StrategyModeTrade,
		Entrypoint:   "on_tick",
	}
)

// All returns the built-in strategy manifest catalog.
func All() []contracts.PluginManifest {
	return []contracts.PluginManifest{GridManifest, DCAManifest, TrendFollowingManifest}
}
