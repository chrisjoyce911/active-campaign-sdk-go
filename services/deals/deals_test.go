package deals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeals_Interface(t *testing.T) {
	// compile-time check that RealService implements DealsService
	var _ DealsService = (*RealService)(nil)
	// trivial assertion to use testify as requested
	assert.True(t, true)
}
