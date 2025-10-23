package contactautomation

import (
	"context"
	"errors"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	th "github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestAccessors_and_nil_receiver_errors(t *testing.T) {
	// Accessors return empty slices when nil
	var l *ListContactAutomationsResponse
	assert.Empty(t, l.ContactAutomationsOrEmpty())

	var a *AutomationCountsResponse
	assert.Empty(t, a.CountsOrEmpty())

	// nil receiver service methods should return not implemented error messages
	var s *service
	_, _, err := s.GetCounts(context.Background())
	assert.Error(t, err)
	_, _, err = s.GetAutomationEntry(context.Background(), "x")
	assert.Error(t, err)
	_, _, err = s.AddContactToAutomation(context.Background(), &CreateContactAutomationRequest{})
	assert.Error(t, err)
	_, err = s.RemoveContactFromAutomation(context.Background(), "x")
	assert.Error(t, err)
	_, _, err = s.ListContactAutomations(context.Background(), "x")
	assert.Error(t, err)
}

func TestClient_errors_propagate(t *testing.T) {
	// MockDoer that returns an error to ensure service propagates it
	md := &th.MockDoer{Err: errors.New("boom")}
	svc := NewRealServiceFromDoer(md)

	_, _, err := svc.GetCounts(context.Background())
	assert.Error(t, err)

	_, _, err = svc.GetAutomationEntry(context.Background(), "id")
	assert.Error(t, err)

	_, _, err = svc.AddContactToAutomation(context.Background(), &CreateContactAutomationRequest{ContactAutomation: ContactAutomationPayload{Contact: "c1", Automation: "a1"}})
	assert.Error(t, err)

	_, err = svc.RemoveContactFromAutomation(context.Background(), "id")
	assert.Error(t, err)

	_, _, err = svc.ListContactAutomations(context.Background(), "c1")
	assert.Error(t, err)
}

func TestRecordingDoer_requests(t *testing.T) {
	rd := &th.RecordingDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"contactAutomations": [{"id":"ca1","automation":"a1"}]}`)}
	svc := NewRealServiceFromDoer(rd)

	out, apiResp, err := svc.ListContactAutomations(context.Background(), "c1")
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	// verify RecordingDoer captured the path
	assert.Contains(t, rd.LastPath, "contacts/c1/contactAutomations")
	if out != nil && len(out.ContactAutomationsOrEmpty()) > 0 {
		assert.Equal(t, "ca1", out.ContactAutomationsOrEmpty()[0].ID)
	}
}

func TestNewRealService_constructor(t *testing.T) {
	cc, err := client.NewCoreClient("http://example.com", "tok")
	assert.NoError(t, err)
	svc := NewRealService(cc)
	assert.NotNil(t, svc)
}
