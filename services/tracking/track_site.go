//go:build ignore

package tracking

import "context"

// TrackSite registers a site visit or tracking pixel event.
func (s *service) TrackSite(ctx context.Context, payload interface{}) (*client.APIResponse, error) {
	return nil, nil
}
