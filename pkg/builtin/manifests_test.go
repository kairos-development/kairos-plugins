package builtin

import (
	"testing"

	"github.com/kairos-development/kairos-contracts/pkg/contracts"
	"github.com/stretchr/testify/assert"
)

func TestAllReturnsBuiltInManifestsCopy(t *testing.T) {
	manifests := All()
	assert.Equal(t, []contracts.PluginManifest{GridManifest, DCAManifest, TrendFollowingManifest}, manifests)

	manifests[0].Name = "mutated"
	assert.Equal(t, "grid", GridManifest.Name)
}

func TestBuiltInManifestsAreTradeStrategies(t *testing.T) {
	for _, manifest := range All() {
		assert.NotEmpty(t, manifest.Name)
		assert.Equal(t, manifest.Name, manifest.Strategy)
		if manifest.Name == "trend_following" {
			assert.Equal(t, "trend_following", manifest.Strategy)
		}
		assert.Equal(t, contracts.ABIVersion{Major: 1, Minor: 0}, manifest.ABIVersion)
		assert.Equal(t, contracts.StrategyModeTrade, manifest.StrategyMode)
		assert.Equal(t, "on_tick", manifest.Entrypoint)
	}
}
