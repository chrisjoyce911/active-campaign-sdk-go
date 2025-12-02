package main

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	contactsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/contacts"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func TestRunDryRun(t *testing.T) {
	buf := &bytes.Buffer{}
	svc := &contactsmock.Service{}
	cfg := runConfig{ContactID: "1", ListID: "2", Status: ""}

	if err := run(context.Background(), svc, cfg, false, buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := buf.String(); !strings.Contains(got, "dry-run") {
		t.Fatalf("expected dry-run output, got %q", got)
	}
}

func TestRunApplySuccess(t *testing.T) {
	buf := &bytes.Buffer{}
	called := false
	svc := &contactsmock.Service{
		AddContactToListFunc: func(ctx context.Context, req *contacts.AddContactToListPayload) (*contacts.AddContactToListResponse, *client.APIResponse, error) {
			called = true
			if req.Contact != "1" || req.List != "2" || req.Status != "1" {
				t.Fatalf("unexpected payload: %+v", req)
			}
			return &contacts.AddContactToListResponse{ContactList: &contacts.ContactList{ID: "cl1"}}, &client.APIResponse{StatusCode: 201}, nil
		},
	}

	cfg := runConfig{ContactID: "1", ListID: "2"}
	if err := run(context.Background(), svc, cfg, true, buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !called {
		t.Fatalf("expected AddContactToList to be called")
	}
	if got := buf.String(); !strings.Contains(got, "cl1") {
		t.Fatalf("expected output to contain membership id, got %q", got)
	}
}

func TestRunApplyError(t *testing.T) {
	svc := &contactsmock.Service{
		AddContactToListFunc: func(ctx context.Context, req *contacts.AddContactToListPayload) (*contacts.AddContactToListResponse, *client.APIResponse, error) {
			return nil, &client.APIResponse{StatusCode: 400}, errors.New("boom")
		},
	}
	cfg := runConfig{ContactID: "1", ListID: "2", Status: "2"}

	err := run(context.Background(), svc, cfg, true, nil)
	if err == nil || !strings.Contains(err.Error(), "status 400") {
		t.Fatalf("expected wrapped status error, got %v", err)
	}
}

func TestRunMissingInputs(t *testing.T) {
	svc := &contactsmock.Service{}
	err := run(context.Background(), svc, runConfig{}, true, nil)
	if err == nil {
		t.Fatalf("expected error for missing ids")
	}
}
