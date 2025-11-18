# mocks/contacts

Function-field mock for `contacts.ContactsService`.

```go
import (
  contactsmock "github.com/chrisjoyce911/active-campaign-sdk-go/mocks/contacts"
)

svc := &contactsmock.Service{
  SearchByEmailFunc: func(ctx context.Context, email string) (*contacts.ContactSearchResponse, *client.APIResponse, error) {
    return &contacts.ContactSearchResponse{Contacts: []contacts.Contact{{ID: "c1"}}}, nil, nil
  },
}
```

Only set fields you need; defaults are no-op.
