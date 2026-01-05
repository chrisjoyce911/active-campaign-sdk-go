package active_campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient_NilOpts_ReturnsStub(t *testing.T) {
	c, err := NewClient(nil)
	assert.NoError(t, err)
	if assert.NotNil(t, c) {
		// stubbed Contacts should be present but not wired
		out, resp, err := c.Contacts.SearchEmail("x@example.com")
		assert.Nil(t, out)
		assert.Nil(t, resp)
		assert.Nil(t, err)
	}
}

func TestNewClient_InvalidBaseURL_ReturnsError(t *testing.T) {
	c, err := NewClient(&ClientOpts{BaseUrl: ":"})
	assert.Error(t, err)
	assert.Nil(t, c)
}

func TestContactsAPI_Placeholders_ReturnNil(t *testing.T) {
	c := &ContactsAPI{}
	out, resp, err := c.GetContactLists("1")
	assert.Nil(t, out)
	assert.Nil(t, resp)
	assert.Nil(t, err)

	out2, resp2, err2 := c.TagsGet("1")
	assert.Nil(t, out2)
	assert.Nil(t, resp2)
	assert.Nil(t, err2)

	rem, err3 := c.RemoveContactTag("1", "2")
	assert.Nil(t, rem)
	assert.Nil(t, err3)
}

func TestTagsAPI_RemoveTag_ReturnNil(t *testing.T) {
	tAPI := &TagsAPI{}
	resp, err := tAPI.RemoveTag("t1")
	assert.Nil(t, resp)
	assert.Nil(t, err)
}
