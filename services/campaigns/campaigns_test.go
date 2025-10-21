package campaigns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign_NotImplemented(t *testing.T) {
	s := &service{}

	tests := []struct {
		name string
	}{
		{name: "returns not implemented"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := s.CreateCampaign(context.Background(), nil)
			if assert.Error(t, err) {
				assert.Contains(t, err.Error(), "not implemented")
			}
		})
	}
}
