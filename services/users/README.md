# services/users

Create a user
post
https://{youraccountname}.api-us1.com/api/3/users
Create a new user

POST /users (Example REQUEST)

```json
{
  "user": {
    "username": "jdoe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "password": "myPa$$w0rd",
    "group": 4
  }
}
```

POST /users (Example RESPONSE)

```json
{
  "user": {
    "username": "jdoe",
    "email": "johndoe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "lang": "english",
    "localZoneid": "America/New_York",
    "cdate": "2022-02-02T16:01:44-06:00",
    "udate": "2022-02-02T16:01:44-06:00",
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/3/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/3/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/3/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/3/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/3/configs",
      "dealConnection": "https://:account.api-us1.com/api/3/users/3/dealConnection",
      "userConversationsPermission": "https://:account.api-us1.com/api/3/users/3/userConversationsPermission",
      "seatUser": "https://:account.api-us1.com/api/3/users/3/seatUser"
    },
    "id": "3"
  }
}
```

201 Created

```json
{
  "user": {
    "username": "jdoe",
    "email": "johndoe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "lang": "english",
    "localZoneid": "America/New_York",
    "cdate": "2022-02-02T16:01:44-06:00",
    "udate": "2022-02-02T16:01:44-06:00",
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/3/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/3/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/3/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/3/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/3/configs",
      "dealConnection": "https://:account.api-us1.com/api/3/users/3/dealConnection",
      "userConversationsPermission": "https://:account.api-us1.com/api/3/users/3/userConversationsPermission",
      "seatUser": "https://:account.api-us1.com/api/3/users/3/seatUser"
    },
    "id": "3"
  }
}
```

400 Bad Request

```json
{}
```

Retrieve a user
get
https://{youraccountname}.api-us1.com/api/3/users/{id}
Retrieve an existing user

```json
{
  "user": {
    "username": "admin",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

200 OK

```json
{
  "user": {
    "username": "admin",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

400 Bad Request

```json
{}
```

Retrieve a user by email
get
https://{youraccountname}.api-us1.com/api/3/users/email/{email}
Retrieve an existing user by looking up their email address

```json
{
  "user": {
    "username": "somebody",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

200 OK

```json
{
  "user": {
    "username": "somebody",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

400 Bad Request

```json
{}
```

Retrieve a user by username
get
https://{youraccountname}.api-us1.com/api/3/users/username/{username}
Retrieve an existing user by looking up their username

```json
{
  "user": {
    "username": "admin",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

200 OK

```json
{
  "user": {
    "username": "admin",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

400 Bad Request

```json
{}
```

Retrieve logged-in user
get
https://{youraccountname}.api-us1.com/api/3/users/me
Retrieve the logged-in user

```json
{
  "user": {
    "username": "jdoe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

200 OK

```json
{
  "user": {
    "username": "jdoe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": null,
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/1/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/1/configs"
    },
    "id": "1"
  }
}
```

400 Bad Request

```json
{}
```

Update a user
put
https://{youraccountname}.api-us1.com/api/3/users/{id}
Update an existing user

PUT /users/:id (Example REQUEST)

```json
{
  "user": {
    "username": "jdoe2",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe2@example.com",
    "password": "myPa$$w0rd",
    "group": 3
  }
}
```

PUT /users/:id (Example RESPONSE)

```json
{
  "userGroups": [
    {
      "userid": "3",
      "groupid": "3",
      "links": {
        "group": "https://:account.api-us1.com/api/3/userGroups/3/group",
        "user": "https://:account.api-us1.com/api/3/userGroups/3/user"
      },
      "id": "3",
      "group": "3",
      "user": "3"
    }
  ],
  "user": {
    "username": "user",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": "",
    "userGroup": "3",
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/3/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/3/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/3/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/3/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/3/configs"
    },
    "id": "3"
  }
}
```

200 OK

```json
{
  "userGroups": [
    {
      "userid": "3",
      "groupid": "3",
      "links": {
        "group": "https://:account.api-us1.com/api/3/userGroups/3/group",
        "user": "https://:account.api-us1.com/api/3/userGroups/3/user"
      },
      "id": "3",
      "group": "3",
      "user": "3"
    }
  ],
  "user": {
    "username": "user",
    "firstName": "John",
    "lastName": "Doe",
    "email": "johndoe@example.com",
    "phone": "",
    "signature": "",
    "userGroup": "3",
    "links": {
      "lists": "https://:account.api-us1.com/api/3/users/3/lists",
      "userGroup": "https://:account.api-us1.com/api/3/users/3/userGroup",
      "dealGroupTotals": "https://:account.api-us1.com/api/3/users/3/dealGroupTotals",
      "dealGroupUsers": "https://:account.api-us1.com/api/3/users/3/dealGroupUsers",
      "configs": "https://:account.api-us1.com/api/3/users/3/configs"
    },
    "id": "3"
  }
}
```

400 Bad Request

```json
{}
```

Delete a user
delete
https://{youraccountname}.api-us1.com/api/3/users/{id}
Delete an existing user

200 OK

```json
{}
```

400 Bad Request

```json
{}
```

List all users
get
https://{youraccountname}.api-us1.com/api/3/users
List all existing users

```json
{
  "users": [
    {
      "username": "admin",
      "firstName": "John",
      "lastName": "Doe",
      "email": "johndoe@activecampaign.com",
      "phone": "",
      "signature": null,
      "links": {
        "lists": "https://:account.api-us1.com/api/3/users/1/lists",
        "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
        "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
        "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
        "configs": "https://:account.api-us1.com/api/3/users/1/configs"
      },
      "id": "1"
    },
    {
      "username": "janedoe",
      "firstName": "Jane",
      "lastName": "Doe",
      "email": "janedoe@activecampaign.com",
      "phone": "",
      "signature": null,
      "links": {
        "lists": "https://:account.api-us1.com/api/3/users/5/lists",
        "userGroup": "https://:account.api-us1.com/api/3/users/5/userGroup",
        "dealGroupTotals": "https://:account.api-us1.com/api/3/users/5/dealGroupTotals",
        "dealGroupUsers": "https://:account.api-us1.com/api/3/users/5/dealGroupUsers",
        "configs": "https://:account.api-us1.com/api/3/users/5/configs"
      },
      "id": "5"
    }
  ],
  "meta": {
    "total": "2"
  }
}
```

200 OK

```json
{
  "users": [
    {
      "username": "admin",
      "firstName": "John",
      "lastName": "Doe",
      "email": "johndoe@activecampaign.com",
      "phone": "",
      "signature": null,
      "links": {
        "lists": "https://:account.api-us1.com/api/3/users/1/lists",
        "userGroup": "https://:account.api-us1.com/api/3/users/1/userGroup",
        "dealGroupTotals": "https://:account.api-us1.com/api/3/users/1/dealGroupTotals",
        "dealGroupUsers": "https://:account.api-us1.com/api/3/users/1/dealGroupUsers",
        "configs": "https://:account.api-us1.com/api/3/users/1/configs"
      },
      "id": "1"
    },
    {
      "username": "janedoe",
      "firstName": "Jane",
      "lastName": "Doe",
      "email": "janedoe@activecampaign.com",
      "phone": "",
      "signature": null,
      "links": {
        "lists": "https://:account.api-us1.com/api/3/users/5/lists",
        "userGroup": "https://:account.api-us1.com/api/3/users/5/userGroup",
        "dealGroupTotals": "https://:account.api-us1.com/api/3/users/5/dealGroupTotals",
        "dealGroupUsers": "https://:account.api-us1.com/api/3/users/5/dealGroupUsers",
        "configs": "https://:account.api-us1.com/api/3/users/5/configs"
      },
      "id": "5"
    }
  ],
  "meta": {
    "total": "2"
  }
}
```

400 Bad Request

```json
{}
```

Create a group
post
https://{youraccountname}.api-us1.com/api/3/groups

POST /groups (Example REQUEST)

```json
{
  "group": {
    "title": "My Groups Title"
  }
}
```

POST /groups (Example RESPONSE)

```json
{
  "group": {
    "title": "My Groups Title",
    "p_admin": 1,
    "links": {
      "userGroups": "https://:account.api-us1.com/api/:version/groups/7/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/:version/groups/7/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/:version/groups/7/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/:version/groups/7/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/:version/groups/7/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/:version/groups/7/automationGroups"
    },
    "id": "7"
  }
}
```

200

```json
{
  "group": {
    "title": "My Groups Title",
    "p_admin": 1,
    "links": {
      "userGroups": "https://:account.api-us1.com/api/:version/groups/7/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/:version/groups/7/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/:version/groups/7/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/:version/groups/7/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/:version/groups/7/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/:version/groups/7/automationGroups"
    },
    "id": "7"
  }
}
```

Retrieve a group
get
https://{youraccountname}.api-us1.com/api/3/groups/{id}

```json
{
  "group": {
    "title": "TEST TITLE",
    "descript": "This is a group for admin users (people that can manage content)",
    "unsubscribelink": "0",
    "optinconfirm": "0",
    "p_admin": "1",
    "pgListAdd": "1",
    "pgListEdit": "1",
    "pgListDelete": "1",
    "pgListHeaders": "1",
    "pgListEmailaccount": "1",
    "pgListBounce": "1",
    "pgMessageAdd": "1",
    "pgMessageEdit": "1",
    "pgMessageDelete": "1",
    "pgMessageSend": "1",
    "pgContactAdd": "1",
    "pgContactEdit": "1",
    "pgContactDelete": "1",
    "pgContactMerge": "1",
    "pgContactImport": "1",
    "pgContactApprove": "1",
    "pgContactExport": "1",
    "pgContactSync": "1",
    "pgContactFilters": "1",
    "pgContactActions": "0",
    "pgContactFields": "1",
    "pg_user_add": "1",
    "pg_user_edit": "1",
    "pg_user_delete": "1",
    "pgGroupAdd": "1",
    "pgGroupEdit": "1",
    "pgGroupDelete": "1",
    "pgTemplateAdd": "1",
    "pgTemplateEdit": "1",
    "pgTemplateDelete": "1",
    "pgPersonalizationAdd": "1",
    "pgPersonalizationEdit": "1",
    "pgPersonalizationDelete": "1",
    "pgAutomationManage": "1",
    "pgFormEdit": "1",
    "pgReportsCampaign": "1",
    "pgReportsList": "1",
    "pgReportsUser": "1",
    "pgReportsTrend": "1",
    "pgStartupReports": "1",
    "pgStartupGettingstarted": "1",
    "pgDeal": "1",
    "pgDealDelete": "1",
    "pgDealReassign": "1",
    "pgDealGroupAdd": "1",
    "pgDealGroupEdit": "1",
    "pgDealGroupDelete": "1",
    "pgSavedResponsesManage": "1",
    "sdate": "2018-09-20T10:51:14-05:00",
    "reqApproval": "0",
    "reqApproval1st": "2",
    "reqApprovalNotify": "",
    "socialdata": "0",
    "links": {
      "userGroups": "https://:account.api-us1.com/api/3/groups/3/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/3/groups/3/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/3/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/3/groups/3/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/3/groups/3/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/3/groups/3/automationGroups"
    },
    "id": "3"
  }
}
```

200 OK

```json
{
  "group": {
    "title": "TEST TITLE",
    "descript": "This is a group for admin users (people that can manage content)",
    "unsubscribelink": "0",
    "optinconfirm": "0",
    "p_admin": "1",
    "pgListAdd": "1",
    "pgListEdit": "1",
    "pgListDelete": "1",
    "pgListHeaders": "1",
    "pgListEmailaccount": "1",
    "pgListBounce": "1",
    "pgMessageAdd": "1",
    "pgMessageEdit": "1",
    "pgMessageDelete": "1",
    "pgMessageSend": "1",
    "pgContactAdd": "1",
    "pgContactEdit": "1",
    "pgContactDelete": "1",
    "pgContactMerge": "1",
    "pgContactImport": "1",
    "pgContactApprove": "1",
    "pgContactExport": "1",
    "pgContactSync": "1",
    "pgContactFilters": "1",
    "pgContactActions": "0",
    "pgContactFields": "1",
    "pg_user_add": "1",
    "pg_user_edit": "1",
    "pg_user_delete": "1",
    "pgGroupAdd": "1",
    "pgGroupEdit": "1",
    "pgGroupDelete": "1",
    "pgTemplateAdd": "1",
    "pgTemplateEdit": "1",
    "pgTemplateDelete": "1",
    "pgPersonalizationAdd": "1",
    "pgPersonalizationEdit": "1",
    "pgPersonalizationDelete": "1",
    "pgAutomationManage": "1",
    "pgFormEdit": "1",
    "pgReportsCampaign": "1",
    "pgReportsList": "1",
    "pgReportsUser": "1",
    "pgReportsTrend": "1",
    "pgStartupReports": "1",
    "pgStartupGettingstarted": "1",
    "pgDeal": "1",
    "pgDealDelete": "1",
    "pgDealReassign": "1",
    "pgDealGroupAdd": "1",
    "pgDealGroupEdit": "1",
    "pgDealGroupDelete": "1",
    "pgSavedResponsesManage": "1",
    "sdate": "2018-09-20T10:51:14-05:00",
    "reqApproval": "0",
    "reqApproval1st": "2",
    "reqApprovalNotify": "",
    "socialdata": "0",
    "links": {
      "userGroups": "https://:account.api-us1.com/api/3/groups/3/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/3/groups/3/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/3/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/3/groups/3/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/3/groups/3/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/3/groups/3/automationGroups"
    },
    "id": "3"
  }
}
```

403 Forbidden

```json
{
  "message": "No Result found for Group with id 10"
}
```

Update a group
put
https://{youraccountname}.api-us1.com/api/3/groups/{id}

PUT /groups/:id (Example REQUEST)

```json
{
  "group": {
    "title": "My Groups Title"
  }
}
```

PUT /groups/:id (Example RESPONSE)

```json
{
  "group": {
    "title": "My Groups Title",
    "p_admin": 1,
    "links": {
      "userGroups": "https://:account.api-us1.com/api/:version/groups/7/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/:version/groups/7/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/:version/groups/7/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/:version/groups/7/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/:version/groups/7/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/:version/groups/7/automationGroups"
    },
    "id": "7"
  }
}
```

200 OK

```json
{
  "group": {
    "title": "My Groups Title",
    "p_admin": 1,
    "links": {
      "userGroups": "https://:account.api-us1.com/api/:version/groups/7/userGroups",
      "groupLimit": "https://:account.api-us1.com/api/:version/groups/7/groupLimit",
      "dealGroupGroups": "https://:account.api-us1.com/api/:version/groups/7/dealGroupGroups",
      "listGroups": "https://:account.api-us1.com/api/:version/groups/7/listGroups",
      "addressGroups": "https://:account.api-us1.com/api/:version/groups/7/addressGroups",
      "automationGroups": "https://:account.api-us1.com/api/:version/groups/7/automationGroups"
    },
    "id": "7"
  }
}
```

403 Forbidden

```json
{
  "message": "No Result found for Group with id 8"
}
```

Delete a group
delete
https://{youraccountname}.api-us1.com/api/3/groups/{id}

200 OK

```json
{}
```

400 Bad Request

```json
{}
```

List all groups
get
https://{youraccountname}.api-us1.com/api/3/groups

```json
{
  "groups": [
    {
      "title": "TEST TITLE",
      "descript": null,
      "unsubscribelink": "0",
      "optinconfirm": "0",
      "p_admin": "1",
      "pgListAdd": "0",
      "pgListEdit": "0",
      "pgListDelete": "0",
      "pgListHeaders": "0",
      "pgListEmailaccount": "0",
      "pgListBounce": "0",
      "pgMessageAdd": "0",
      "pgMessageEdit": "0",
      "pgMessageDelete": "0",
      "pgMessageSend": "0",
      "pgContactAdd": "0",
      "pgContactEdit": "0",
      "pgContactDelete": "0",
      "pgContactMerge": "0",
      "pgContactImport": "0",
      "pgContactApprove": "0",
      "pgContactExport": "0",
      "pgContactSync": "0",
      "pgContactFilters": "0",
      "pgContactActions": "0",
      "pgContactFields": "0",
      "pg_user_add": "0",
      "pg_user_edit": "0",
      "pg_user_delete": "0",
      "pgGroupAdd": "0",
      "pgGroupEdit": "0",
      "pgGroupDelete": "0",
      "pgTemplateAdd": "0",
      "pgTemplateEdit": "0",
      "pgTemplateDelete": "0",
      "pgPersonalizationAdd": "0",
      "pgPersonalizationEdit": "0",
      "pgPersonalizationDelete": "0",
      "pgAutomationManage": "0",
      "pgFormEdit": "0",
      "pgReportsCampaign": "0",
      "pgReportsList": "0",
      "pgReportsUser": "0",
      "pgReportsTrend": "1",
      "pgStartupReports": "0",
      "pgStartupGettingstarted": "1",
      "pgDeal": "1",
      "pgDealDelete": "1",
      "pgDealReassign": "1",
      "pgDealGroupAdd": "1",
      "pgDealGroupEdit": "1",
      "pgDealGroupDelete": "1",
      "pgSavedResponsesManage": "0",
      "sdate": null,
      "reqApproval": "0",
      "reqApproval1st": "2",
      "reqApprovalNotify": "",
      "socialdata": "0",
      "links": {
        "userGroups": "https://:account.api-us1.com/api/3/groups/7/userGroups",
        "groupLimit": "https://:account.api-us1.com/api/3/groups/7/groupLimit",
        "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/7/dealGroupGroups",
        "listGroups": "https://:account.api-us1.com/api/3/groups/7/listGroups",
        "addressGroups": "https://:account.api-us1.com/api/3/groups/7/addressGroups",
        "automationGroups": "https://:account.api-us1.com/api/3/groups/7/automationGroups"
      },
      "id": "7"
    }
  ],
  "meta": {
    "total": "1"
  }
}
```

200 OK

```json
{
  "groups": [
    {
      "title": "TEST TITLE",
      "descript": "This is a group for admin users (people that can manage content)",
      "unsubscribelink": "0",
      "optinconfirm": "0",
      "p_admin": "1",
      "pgListAdd": "1",
      "pgListEdit": "1",
      "pgListDelete": "1",
      "pgListHeaders": "1",
      "pgListEmailaccount": "1",
      "pgListBounce": "1",
      "pgMessageAdd": "1",
      "pgMessageEdit": "1",
      "pgMessageDelete": "1",
      "pgMessageSend": "1",
      "pgContactAdd": "1",
      "pgContactEdit": "1",
      "pgContactDelete": "1",
      "pgContactMerge": "1",
      "pgContactImport": "1",
      "pgContactApprove": "1",
      "pgContactExport": "1",
      "pgContactSync": "1",
      "pgContactFilters": "1",
      "pgContactActions": "0",
      "pgContactFields": "1",
      "pg_user_add": "1",
      "pg_user_edit": "1",
      "pg_user_delete": "1",
      "pgGroupAdd": "1",
      "pgGroupEdit": "1",
      "pgGroupDelete": "1",
      "pgTemplateAdd": "1",
      "pgTemplateEdit": "1",
      "pgTemplateDelete": "1",
      "pgPersonalizationAdd": "1",
      "pgPersonalizationEdit": "1",
      "pgPersonalizationDelete": "1",
      "pgAutomationManage": "1",
      "pgFormEdit": "1",
      "pgReportsCampaign": "1",
      "pgReportsList": "1",
      "pgReportsUser": "1",
      "pgReportsTrend": "1",
      "pgStartupReports": "1",
      "pgStartupGettingstarted": "1",
      "pgDeal": "1",
      "pgDealDelete": "1",
      "pgDealReassign": "1",
      "pgDealGroupAdd": "1",
      "pgDealGroupEdit": "1",
      "pgDealGroupDelete": "1",
      "pgSavedResponsesManage": "1",
      "sdate": "2018-09-20T10:51:14-05:00",
      "reqApproval": "0",
      "reqApproval1st": "2",
      "reqApprovalNotify": "",
      "socialdata": "0",
      "links": {
        "userGroups": "https://:account.api-us1.com/api/3/groups/3/userGroups",
        "groupLimit": "https://:account.api-us1.com/api/3/groups/3/groupLimit",
        "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/3/dealGroupGroups",
        "listGroups": "https://:account.api-us1.com/api/3/groups/3/listGroups",
        "addressGroups": "https://:account.api-us1.com/api/3/groups/3/addressGroups",
        "automationGroups": "https://:account.api-us1.com/api/3/groups/3/automationGroups"
      },
      "id": "3"
    },
    {
      "title": "TEST TITLE",
      "descript": null,
      "unsubscribelink": "0",
      "optinconfirm": "0",
      "p_admin": "1",
      "pgListAdd": "0",
      "pgListEdit": "0",
      "pgListDelete": "0",
      "pgListHeaders": "0",
      "pgListEmailaccount": "0",
      "pgListBounce": "0",
      "pgMessageAdd": "0",
      "pgMessageEdit": "0",
      "pgMessageDelete": "0",
      "pgMessageSend": "0",
      "pgContactAdd": "0",
      "pgContactEdit": "0",
      "pgContactDelete": "0",
      "pgContactMerge": "0",
      "pgContactImport": "0",
      "pgContactApprove": "0",
      "pgContactExport": "0",
      "pgContactSync": "0",
      "pgContactFilters": "0",
      "pgContactActions": "0",
      "pgContactFields": "0",
      "pg_user_add": "0",
      "pg_user_edit": "0",
      "pg_user_delete": "0",
      "pgGroupAdd": "0",
      "pgGroupEdit": "0",
      "pgGroupDelete": "0",
      "pgTemplateAdd": "0",
      "pgTemplateEdit": "0",
      "pgTemplateDelete": "0",
      "pgPersonalizationAdd": "0",
      "pgPersonalizationEdit": "0",
      "pgPersonalizationDelete": "0",
      "pgAutomationManage": "0",
      "pgFormEdit": "0",
      "pgReportsCampaign": "0",
      "pgReportsList": "0",
      "pgReportsUser": "0",
      "pgReportsTrend": "1",
      "pgStartupReports": "0",
      "pgStartupGettingstarted": "1",
      "pgDeal": "1",
      "pgDealDelete": "1",
      "pgDealReassign": "1",
      "pgDealGroupAdd": "1",
      "pgDealGroupEdit": "1",
      "pgDealGroupDelete": "1",
      "pgSavedResponsesManage": "0",
      "sdate": null,
      "reqApproval": "0",
      "reqApproval1st": "2",
      "reqApprovalNotify": "",
      "socialdata": "0",
      "links": {
        "userGroups": "https://:account.api-us1.com/api/3/groups/5/userGroups",
        "groupLimit": "https://:account.api-us1.com/api/3/groups/5/groupLimit",
        "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/5/dealGroupGroups",
        "listGroups": "https://:account.api-us1.com/api/3/groups/5/listGroups",
        "addressGroups": "https://:account.api-us1.com/api/3/groups/5/addressGroups",
        "automationGroups": "https://:account.api-us1.com/api/3/groups/5/automationGroups"
      },
      "id": "5"
    },
    {
      "title": "TEST TITLE",
      "descript": null,
      "unsubscribelink": "0",
      "optinconfirm": "0",
      "p_admin": "1",
      "pgListAdd": "0",
      "pgListEdit": "0",
      "pgListDelete": "0",
      "pgListHeaders": "0",
      "pgListEmailaccount": "0",
      "pgListBounce": "0",
      "pgMessageAdd": "0",
      "pgMessageEdit": "0",
      "pgMessageDelete": "0",
      "pgMessageSend": "0",
      "pgContactAdd": "0",
      "pgContactEdit": "0",
      "pgContactDelete": "0",
      "pgContactMerge": "0",
      "pgContactImport": "0",
      "pgContactApprove": "0",
      "pgContactExport": "0",
      "pgContactSync": "0",
      "pgContactFilters": "0",
      "pgContactActions": "0",
      "pgContactFields": "0",
      "pg_user_add": "0",
      "pg_user_edit": "0",
      "pg_user_delete": "0",
      "pgGroupAdd": "0",
      "pgGroupEdit": "0",
      "pgGroupDelete": "0",
      "pgTemplateAdd": "0",
      "pgTemplateEdit": "0",
      "pgTemplateDelete": "0",
      "pgPersonalizationAdd": "0",
      "pgPersonalizationEdit": "0",
      "pgPersonalizationDelete": "0",
      "pgAutomationManage": "0",
      "pgFormEdit": "0",
      "pgReportsCampaign": "0",
      "pgReportsList": "0",
      "pgReportsUser": "0",
      "pgReportsTrend": "1",
      "pgStartupReports": "0",
      "pgStartupGettingstarted": "1",
      "pgDeal": "1",
      "pgDealDelete": "1",
      "pgDealReassign": "1",
      "pgDealGroupAdd": "1",
      "pgDealGroupEdit": "1",
      "pgDealGroupDelete": "1",
      "pgSavedResponsesManage": "0",
      "sdate": null,
      "reqApproval": "0",
      "reqApproval1st": "2",
      "reqApprovalNotify": "",
      "socialdata": "0",
      "links": {
        "userGroups": "https://:account.api-us1.com/api/3/groups/6/userGroups",
        "groupLimit": "https://:account.api-us1.com/api/3/groups/6/groupLimit",
        "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/6/dealGroupGroups",
        "listGroups": "https://:account.api-us1.com/api/3/groups/6/listGroups",
        "addressGroups": "https://:account.api-us1.com/api/3/groups/6/addressGroups",
        "automationGroups": "https://:account.api-us1.com/api/3/groups/6/automationGroups"
      },
      "id": "6"
    },
    {
      "title": "TEST TITLE",
      "descript": null,
      "unsubscribelink": "0",
      "optinconfirm": "0",
      "p_admin": "1",
      "pgListAdd": "0",
      "pgListEdit": "0",
      "pgListDelete": "0",
      "pgListHeaders": "0",
      "pgListEmailaccount": "0",
      "pgListBounce": "0",
      "pgMessageAdd": "0",
      "pgMessageEdit": "0",
      "pgMessageDelete": "0",
      "pgMessageSend": "0",
      "pgContactAdd": "0",
      "pgContactEdit": "0",
      "pgContactDelete": "0",
      "pgContactMerge": "0",
      "pgContactImport": "0",
      "pgContactApprove": "0",
      "pgContactExport": "0",
      "pgContactSync": "0",
      "pgContactFilters": "0",
      "pgContactActions": "0",
      "pgContactFields": "0",
      "pg_user_add": "0",
      "pg_user_edit": "0",
      "pg_user_delete": "0",
      "pgGroupAdd": "0",
      "pgGroupEdit": "0",
      "pgGroupDelete": "0",
      "pgTemplateAdd": "0",
      "pgTemplateEdit": "0",
      "pgTemplateDelete": "0",
      "pgPersonalizationAdd": "0",
      "pgPersonalizationEdit": "0",
      "pgPersonalizationDelete": "0",
      "pgAutomationManage": "0",
      "pgFormEdit": "0",
      "pgReportsCampaign": "0",
      "pgReportsList": "0",
      "pgReportsUser": "0",
      "pgReportsTrend": "1",
      "pgStartupReports": "0",
      "pgStartupGettingstarted": "1",
      "pgDeal": "1",
      "pgDealDelete": "1",
      "pgDealReassign": "1",
      "pgDealGroupAdd": "1",
      "pgDealGroupEdit": "1",
      "pgDealGroupDelete": "1",
      "pgSavedResponsesManage": "0",
      "sdate": null,
      "reqApproval": "0",
      "reqApproval1st": "2",
      "reqApprovalNotify": "",
      "socialdata": "0",
      "links": {
        "userGroups": "https://:account.api-us1.com/api/3/groups/7/userGroups",
        "groupLimit": "https://:account.api-us1.com/api/3/groups/7/groupLimit",
        "dealGroupGroups": "https://:account.api-us1.com/api/3/groups/7/dealGroupGroups",
        "listGroups": "https://:account.api-us1.com/api/3/groups/7/listGroups",
        "addressGroups": "https://:account.api-us1.com/api/3/groups/7/addressGroups",
        "automationGroups": "https://:account.api-us1.com/api/3/groups/7/automationGroups"
      },
      "id": "7"
    }
  ],
  "meta": {
    "total": "4"
  }
}
```

List All Group Limits
get
https://{youraccountname}.api-us1.com/api/3/groupLimits

```json
{
  "groupLimits": [
    {
      "groupid": "1",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/1/group"
      },
      "id": "1",
      "group": "1"
    },
    {
      "groupid": "2",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:accounts.api-us1.com/api/3/groupLimits/2/group"
      },
      "id": "2",
      "group": "2"
    },
    {
      "groupid": "3",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/3/group"
      },
      "id": "3",
      "group": "3"
    },
    {
      "groupid": "4",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/4/group"
      },
      "id": "4",
      "group": "4"
    }
  ],
  "meta": {
    "total": "4"
  }
}
```

200 OK

```json
{
  "groupLimits": [
    {
      "groupid": "1",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/1/group"
      },
      "id": "1",
      "group": "1"
    },
    {
      "groupid": "2",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:accounts.api-us1.com/api/3/groupLimits/2/group"
      },
      "id": "2",
      "group": "2"
    },
    {
      "groupid": "3",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/3/group"
      },
      "id": "3",
      "group": "3"
    },
    {
      "groupid": "4",
      "limitMail": "0",
      "limitMailType": "month",
      "limitContact": "0",
      "limitList": "0",
      "limitCampaign": "0",
      "limitCampaignType": "month",
      "limitAttachment": "-1",
      "limitUser": "0",
      "abuseRatio": "4",
      "forceSenderInfo": "0",
      "links": {
        "group": "https://:account.api-us1.com/api/3/groupLimits/4/group"
      },
      "id": "4",
      "group": "4"
    }
  ],
  "meta": {
    "total": "4"
  }
}
```
