package active_campaign

// root_stub.go provides minimal, buildable stubs for the module root package
// so example programs that import the root module can compile during migration.

// ClientOpts are a simplified placeholder for example compatibility.
type ClientOpts struct {
	BaseUrl    string
	Token      string
	HttpClient interface{}
}

// Response is a placeholder for the legacy Response wrapper used by examples.
type Response struct{}

// Client is a minimal stub with a Contacts field to match old usage in examples.
type Client struct {
	Contacts *ContactsAPI
	Tags     *TagsAPI
}

// ContactsAPI is a minimal stub providing SearchEmail used by examples.
type ContactsAPI struct{}

// SearchEmail is a placeholder that returns nils.
func (c *ContactsAPI) SearchEmail(email string) (interface{}, *Response, error) {
	return nil, nil, nil
}

// GetContactLists placeholder
func (c *ContactsAPI) GetContactLists(contactID string) (interface{}, *Response, error) {
	return nil, nil, nil
}

// GetContactTags placeholder
func (c *ContactsAPI) GetContactTags(contactID string) (interface{}, *Response, error) {
	return nil, nil, nil
}

// RemoveContactTag placeholder
func (c *ContactsAPI) RemoveContactTag(contactID, tagID string) (*Response, error) {
	return nil, nil
}

// NewClient returns a minimal Client stub for examples.
func NewClient(opts *ClientOpts) (*Client, error) {
	return &Client{Contacts: &ContactsAPI{}, Tags: &TagsAPI{}}, nil
}

// TagsAPI placeholder
type TagsAPI struct{}

// RemoveTag placeholder
func (t *TagsAPI) RemoveTag(tagID string) (*Response, error) {
	return nil, nil
}
