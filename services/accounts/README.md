# services/accounts

Accounts service interface and implementation.

TODO:

- Define AccountsService interface
- Add models and implementation
- Add mocks and tests

## Create an account

POST https://{youraccountname}.api-us1.com/api/3/accounts

POST /accounts (Example REQUEST)

```json
{
  "account": {
    "name": "Example Account",
    "accountUrl": "https://www.example.com",
    "owner": 1,
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "500-1000"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP"
      }
    ]
  }
}
```

POST /accounts (Example RESPONSE)

```json
{
  "account": {
    "name": "Example Account",
    "accountUrl": "https://www.example.com",
    "createdTimestamp": "2019-06-12T16:52:16-05:00",
    "updatedTimestamp": "2019-06-12T16:52:16-05:00",
    "links": [],
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "501 - 1000",
        "accountId": "1"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP",
        "accountId": "1"
      }
    ],
    "id": "1"
  }
}
```

```bash
curl --request POST \
     --url https://youraccountname.api-us1.com/api/3/accounts \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "account": {
    "owner": 1,
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "500-1000",
        "fieldCurrency": "AU"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP"
      }
    ],
    "name": "Example Account",
    "accountUrl": "https://www.example.com"
  }
}
'
```

201 Created

```json
{
  "account": {
    "name": "Example Account",
    "accountUrl": "https://www.example.com",
    "createdTimestamp": "2019-06-12T16:52:16-05:00",
    "updatedTimestamp": "2019-06-12T16:52:16-05:00",
    "links": [],
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "501 - 1000",
        "accountId": "1"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP",
        "accountId": "1"
      }
    ],
    "id": "1"
  }
}
```

422 Unprocessable Entity (Example RESPONSE)

```json
{
  "errors": [
    {
      "title": "The account name was not provided.",
      "detail": "",
      "code": "field_missing",
      "source": {
        "pointer": "/data/attributes/name"
      }
    }
  ]
}
```

## Update an account

PUT https://{youraccountname}.api-us1.com/api/3/accounts/{id}

PUT /accounts (Example REQUEST)

```json
{
  "account": {
    "name": "Example Account",
    "accountUrl": "https://www.example.com",
    "owner": 1,
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "500-1000"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP"
      }
    ]
  }
}
```

PUT /accounts (Example RESPONSE)

```json
{
  "account": {
    "name": "Exmaple Account",
    "accountUrl": "https://www.example.com",
    "createdTimestamp": "2019-04-03T13:57:31-05:00",
    "updatedTimestamp": "2019-06-12T16:55:32-05:00",
    "links": [],
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "501 - 1000",
        "accountId": "1"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP",
        "accountId": "1"
      }
    ],
    "id": "1"
  }
}
```

````bash
curl --request PUT \
     --url https://youraccountname.api-us1.com/api/3/accounts/id \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "account": {
## Examples

There are currently no dedicated examples for the Accounts service in `examples/`.
You can use the example programs in the `examples/` directory as a reference for creating
and wiring a core client and typed service. Example usage to construct a core client and a
service (pseudo-code):

```go
// build a network-capable core client
core, err := client.NewCoreClient(os.Getenv("ACTIVE_URL"), os.Getenv("ACTIVE_TOKEN"))
if err != nil {
    // handle error
}

// create the accounts service when available
// svc := accounts.NewRealService(core)

// call the service methods
// resp, apiResp, err := svc.ListAccounts(ctx)
````

If you'd like an example added here (create, list, get), I can add a small runnable
example under `examples/accounts_list` that demonstrates basic usage.
"fields": [
{
"customFieldId": 9,
"fieldValue": "500-1000"
},
{
"customFieldId": 20,
"fieldValue": 1234,
"fieldCurrency": "GBP"
}
],
"owner": 1,
"name": "Exmaple Account",
"accountUrl": "https://www.example.com"
}
}
'

````

200 OK

```json
{
  "account": {
    "name": "Exmaple Account",
    "accountUrl": "https://www.example.com",
    "createdTimestamp": "2019-04-03T13:57:31-05:00",
    "updatedTimestamp": "2019-06-12T16:55:32-05:00",
    "links": [],
    "fields": [
      {
        "customFieldId": 9,
        "fieldValue": "501 - 1000",
        "accountId": "1"
      },
      {
        "customFieldId": 20,
        "fieldValue": 1234,
        "fieldCurrency": "GBP",
        "accountId": "1"
      }
    ],
    "id": "1"
  }
}
````

400 Bad Request (Example RESPONSE)

```json
{}
```

## Retrieve an account

GET https://{youraccountname}.api-us1.com/api/3/accounts/{id}

```json
{
  "account": {
    "name": "Example Account",
    "accountUrl": "https://www.example.com",
    "createdTimestamp": "2019-05-15T15:58:16-05:00",
    "updatedTimestamp": "2019-05-15T15:58:16-05:00",
    "links": [],
    "id": "1"
  }
}
```

## Delete an account

DELETE https://{youraccountname}.api-us1.com/api/3/accounts/{id}

```bash
curl --request DELETE \
     --url https://youraccountname.api-us1.com/api/3/accounts/id \
     --header 'accept: application/json'
```

200

```json
{}
```

400

```json
{}
```

### List all accounts

GET https://{youraccountname}.api-us1.com/api/3/accounts
Retrieve all existing accounts

```json
{
    "accounts": [
        {
            "name": "First Example Account",
            "accountUrl": null,
            "createdTimestamp": "2019-04-29T07:51:31-05:00",
            "updatedTimestamp": "2019-04-29T07:51:31-05:00",
            "contactCount": "1",
            "dealCount": "3",
            "links": [
                "notes": "https://:account.api-us1.com/api/:version/accounts/1/notes",
                "accountCustomFieldData": "https://:account.api-us1.com/api/:version/accounts/1/accountCustomFieldData",
                "accountContacts": "https://:account.api-us1.com/api/:version/accounts/1/accountContacts"],
            "id": "1"
        },
        {
            "name": "Second Example Account",
            "accountUrl": null,
            "createdTimestamp": "2019-04-29T07:51:32-05:00",
            "updatedTimestamp": "2019-04-29T07:51:32-05:00",
            "contactCount": "2",
            "dealCount": "5",
            "links": [
                "notes": "https://:account.api-us1.com/api/:version/accounts/2/notes",
                "accountCustomFieldData": "https://:account.api-us1.com/api/:version/accounts/2/accountCustomFieldData",
                "accountContacts": "https://:account.api-us1.com/api/:version/accounts/2/accountContacts"],
            "id": "2"
        }
    ],
    "meta": {
      "total": "2"
    }
}
```

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/accounts \
     --header 'accept: application/json'
```

200 OK

```json
{
    "accounts": [
        {
            "name": "First Example Account",
            "accountUrl": null,
            "createdTimestamp": "2019-04-29T07:51:31-05:00",
            "updatedTimestamp": "2019-04-29T07:51:31-05:00",
            "contactCount": "1",
            "dealCount": "3",
            "links": [
                "notes": "https://:account.api-us1.com/api/:version/accounts/1/notes",
                "accountCustomFieldData": "https://:account.api-us1.com/api/:version/accounts/1/accountCustomFieldData",
                "accountContacts": "https://:account.api-us1.com/api/:version/accounts/1/accountContacts"],
            "id": "1"
        },
        {
            "name": "Second Example Account",
            "accountUrl": null,
            "createdTimestamp": "2019-04-29T07:51:32-05:00",
            "updatedTimestamp": "2019-04-29T07:51:32-05:00",
            "contactCount": "2",
            "dealCount": "5",
            "links": [
                "notes": "https://:account.api-us1.com/api/:version/accounts/2/notes",
                "accountCustomFieldData": "https://:account.api-us1.com/api/:version/accounts/2/accountCustomFieldData",
                "accountContacts": "https://:account.api-us1.com/api/:version/accounts/2/accountContacts"],
            "id": "2"
        }
    ],
    "meta": {
      "total": "2"
    }
}
```

400

```json
{}
```

### Create an account note

POST https://{youraccountname}.api-us1.com/api/3/accounts/{id}/notes
Create a new note for an account

POST /accounts/:id/notes (Example REQUEST)

```json
{
  "note": {
    "note": "Note for the account"
  }
}
```

POST /accounts/:id/notes (Example RESPONSE)

```json
{
  "accounts": [
    {
      "name": "Museuem of Science and Industry",
      "accountUrl": "www.msi.com",
      "createdTimestamp": "2019-04-03T13:29:35-05:00",
      "updatedTimestamp": "2019-06-12T16:59:54-05:00",
      "id": "1"
    }
  ],
  "note": {
    "cdate": "2017-05-31T09:54:30-05:00",
    "id": "2",
    "links": {
      "activities": "/api/3/notes/2/activities",
      "mentions": "/api/3/notes/2/mentions",
      "notes": "/api/3/notes/2/notes",
      "owner": "/api/3/notes/2/owner",
      "user": "/api/3/notes/2/user"
    },
    "mdate": "2017-05-31T09:54:30-05:00",
    "note": "Note for the account",
    "owner": {
      "id": "1",
      "type": "account"
    },
    "relid": "1",
    "reltype": "CustomerAccount",
    "user": "1",
    "userid": "1"
  }
}
```

```bash
curl --request POST \
     --url https://youraccountname.api-us1.com/api/3/accounts/id/notes \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": {
    "note": "Note for the account"
  }
}
'
```

201 Created

```json
{
  "accounts": [
    {
      "name": "Museuem of Science and Industry",
      "accountUrl": "www.msi.com",
      "createdTimestamp": "2019-04-03T13:29:35-05:00",
      "updatedTimestamp": "2019-06-12T16:59:54-05:00",
      "id": "1"
    }
  ],
  "note": {
    "cdate": "2017-05-31T09:54:30-05:00",
    "id": "2",
    "links": {
      "activities": "/api/3/notes/2/activities",
      "mentions": "/api/3/notes/2/mentions",
      "notes": "/api/3/notes/2/notes",
      "owner": "/api/3/notes/2/owner",
      "user": "/api/3/notes/2/user"
    },
    "mdate": "2017-05-31T09:54:30-05:00",
    "note": "Note for the account",
    "owner": {
      "id": "1",
      "type": "account"
    },
    "relid": "1",
    "reltype": "CustomerAccount",
    "user": "1",
    "userid": "1"
  }
}
```

## Update an account note

PUT https://{youraccountname}.api-us1.com/api/3/accounts/{id}/notes/{noteid}

Update an existing note for a account

PUT /accounts/:id/notes/:noteid (Example REQUEST)

```json
{
  "note": {
    "note": "Updated note for the account"
  }
}
```

PUT /accounts/:id/notes/:noteid (Example RESPONSE

```json
{
  "accounts": [
    {
      "name": "Example Account",
      "accountUrl": "https://www.example.url",
      "createdTimestamp": "2019-04-03T13:29:35-05:00",
      "updatedTimestamp": "2019-06-12T16:59:54-05:00",
      "links": {
        "notes": "https://hosted.localdev/api/3/accounts/1/notes"
      },
      "id": "1"
    }
  ],
  "note": {
    "cdate": "2017-06-01T13:42:13-05:00",
    "id": "2",
    "links": {
      "activities": "/api/3/notes/2/activities",
      "mentions": "/api/3/notes/2/mentions",
      "notes": "/api/3/notes/2/notes",
      "owner": "/api/3/notes/2/owner",
      "user": "/api/3/notes/2/user"
    },
    "mdate": "2017-06-01T13:42:13-05:00",
    "note": "Update with more info",
    "owner": {
      "id": "1",
      "type": "account"
    },
    "relid": "1",
    "reltype": "CustomerAccount",
    "user": "1",
    "userid": "1"
  }
}
```

```bash
curl --request PUT \
     --url https://youraccountname.api-us1.com/api/3/accounts/id/notes/noteid \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": {
    "note": "Update with more info"
  }
}
'
```

200 OK

```json
{
  "accounts": [
    {
      "name": "Example Account",
      "accountUrl": "https://www.example.url",
      "createdTimestamp": "2019-04-03T13:29:35-05:00",
      "updatedTimestamp": "2019-06-12T16:59:54-05:00",
      "links": {
        "notes": "https://hosted.localdev/api/3/accounts/1/notes"
      },
      "id": "1"
    }
  ],
  "note": {
    "cdate": "2017-06-01T13:42:13-05:00",
    "id": "2",
    "links": {
      "activities": "/api/3/notes/2/activities",
      "mentions": "/api/3/notes/2/mentions",
      "notes": "/api/3/notes/2/notes",
      "owner": "/api/3/notes/2/owner",
      "user": "/api/3/notes/2/user"
    },
    "mdate": "2017-06-01T13:42:13-05:00",
    "note": "Update with more info",
    "owner": {
      "id": "1",
      "type": "account"
    },
    "relid": "1",
    "reltype": "CustomerAccount",
    "user": "1",
    "userid": "1"
  }
}
```
