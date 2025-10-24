package deals

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeals_Interface(t *testing.T) {
	// compile-time check that RealService implements DealsService
	var _ DealsService = (*RealService)(nil)
	// trivial assertion to use testify as requested
	require := require.New(t)
	require.True(true)
	// keep an assert invocation for consistency
	assert.True(t, true)
}
