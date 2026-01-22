package contacts

import (
	"context"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContacts_SimpleEndpointsAndAccessors(t *testing.T) {
	require := require.New(t)
	// AddFieldOption
	md := &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"fieldOption":{"id":"fo1"}}`)}
	require.NotNil(md)
	svc := NewRealServiceFromDoer(md)
	require.NotNil(svc)
	foOut, apiResp, err := svc.AddFieldOption(context.Background(), &FieldOptionPayload{Value: "v1"})
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	_ = foOut

	// CreateContactTag
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 201}, Body: []byte(`{"contactTag":{"contact":"c1","tag":"t1"}}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	ctOut, apiResp2, err2 := svc.CreateContactTag(context.Background(), &ContactTagRequest{ContactTag: ContactTagPayload{Contact: "c1", Tag: "t1"}})
	assert.NoError(t, err2)
	assert.Equal(t, 201, apiResp2.StatusCode)
	_ = ctOut

	// CreateField (delegates to CreateCustomField)
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"field":{"id":"f1"}}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	fldOut, apiResp3, err3 := svc.CreateField(context.Background(), &FieldPayload{Title: "x"})
	assert.NoError(t, err3)
	assert.Equal(t, 200, apiResp3.StatusCode)
	_ = fldOut

	// DeleteField (delegates)
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 204}}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	delResp, err4 := svc.DeleteField(context.Background(), "f1")
	assert.NoError(t, err4)
	assert.Equal(t, 204, delResp.StatusCode)

	// DeleteFieldGroup
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	dgResp, err5 := svc.DeleteFieldGroup(context.Background(), "g1")
	assert.NoError(t, err5)
	assert.Equal(t, 200, dgResp.StatusCode)

	// GetContactBounceLogs
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"logs": []}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	blOut, blResp, blErr := svc.GetContactBounceLogs(context.Background(), "123")
	assert.NoError(t, blErr)
	assert.Equal(t, 200, blResp.StatusCode)
	_ = blOut

	// GetContactByEmailWithTags
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"contacts": []}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	byOut, byResp, byErr := svc.GetContactByEmailWithTags(context.Background(), "x@example.com")
	assert.NoError(t, byErr)
	assert.Equal(t, 200, byResp.StatusCode)
	_ = byOut

	// GetContactGoals
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"goals": []}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	gOut, gResp, gErr := svc.GetContactGoals(context.Background(), "c1")
	assert.NoError(t, gErr)
	assert.Equal(t, 200, gResp.StatusCode)
	_ = gOut

	// GetContactLists
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"contactLists": []}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	lOut, lResp, lErr := svc.GetContactLists(context.Background(), "c1")
	assert.NoError(t, lErr)
	assert.Equal(t, 200, lResp.StatusCode)
	_ = lOut

	// GetContactLogs
	md = &testhelpers.MockDoer{Resp: &client.APIResponse{StatusCode: 200}, Body: []byte(`{"logs": []}`)}
	require.NotNil(md)
	svc = NewRealServiceFromDoer(md)
	require.NotNil(svc)
	lgOut, lgResp, lgErr := svc.GetContactLogs(context.Background(), "c1")
	assert.NoError(t, lgErr)
	assert.Equal(t, 200, lgResp.StatusCode)
	_ = lgOut

	// NewService placeholder should return nil (impl.go TODO)
	ns := NewService(&http.Client{}, "https://example.com", "tok")
	assert.Nil(t, ns)
}

func TestContacts_ListContactsAndAccessors(t *testing.T) {
	// ListContacts query params
	hd := &testhelpers.HTTPDoer{BaseURL: "https://example.com/api/3/", RespStatus: 200, RespBody: []byte(`{"contacts":[{"id":"c1"}]}`)}
	svc := NewRealServiceFromDoer(hd)
	opts := map[string]string{"search": "term", "limit": "5"}
	out, apiResp, err := svc.ListContacts(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 200, apiResp.StatusCode)
	// verify query params were encoded
	q := hd.LastRequest.URL.Query()
	assert.Equal(t, "term", q.Get("search"))
	assert.Equal(t, "5", q.Get("limit"))
	_ = out

	// Accessor helpers for nil and non-nil cases
	var ctResp *ContactTagsResponse
	empty := ctResp.ContactTagsOrEmpty()
	assert.Equal(t, 0, len(empty))

	tcr := &ContactTagsResponse{ContactTags: &[]ContactTag{{ID: "t1"}}}
	got := tcr.ContactTagsOrEmpty()
	assert.Equal(t, 1, len(got))

	cr := &CreateContactResponse{}
	assert.Equal(t, 0, len(cr.ContactAutomationsOrEmpty()))
	assert.Equal(t, 0, len(cr.ContactDataOrEmpty()))
	assert.Equal(t, 0, len(cr.ContactListsOrEmpty()))
	assert.Equal(t, 0, len(cr.FieldValuesOrEmpty()))
	assert.Equal(t, 0, len(cr.GeoAddressesOrEmpty()))
	assert.Equal(t, 0, len(cr.GeoIpsOrEmpty()))
	assert.Equal(t, 0, len(cr.ScoreValuesOrEmpty()))

	csr := &ContactSearchResponse{}
	assert.Equal(t, 0, len(csr.ScoreValuesOrEmpty()))
}
