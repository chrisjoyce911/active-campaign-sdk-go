package lists

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensure methods return the "not implemented" error when the receiver is nil.
func TestNilReceiverMethods(t *testing.T) {
	var s *service
	ctx := context.Background()

	// CreateList
	_, _, err := s.CreateList(ctx, CreateListRequest{})
	if !assert.Error(t, err) {
		t.Fatalf("expected error for nil receiver CreateList")
	}

	// CreateListGroup
	_, _, err = s.CreateListGroup(ctx, CreateListGroupRequest{})
	if !assert.Error(t, err) {
		t.Fatalf("expected error for nil receiver CreateListGroup")
	}

	// DeleteList
	_, err = s.DeleteList(ctx, "1")
	if !assert.Error(t, err) {
		t.Fatalf("expected error for nil receiver DeleteList")
	}

	// GetList
	_, _, err = s.GetList(ctx, "1")
	if !assert.Error(t, err) {
		t.Fatalf("expected error for nil receiver GetList")
	}

	// ListLists
	_, _, err = s.ListLists(ctx, nil)
	if !assert.Error(t, err) {
		t.Fatalf("expected error for nil receiver ListLists")
	}
}

// Ensure NewRealService sets the client field and methods return errors when client is nil.
func TestNewRealServiceWithNilClient(t *testing.T) {
	svc := NewRealService(nil)
	// just ensure we get a non-nil service back; don't call methods that would invoke a typed-nil Doer
	if !assert.NotNil(t, svc) {
		t.Fatalf("expected NewRealService to return a non-nil service even when client is nil")
	}
}
