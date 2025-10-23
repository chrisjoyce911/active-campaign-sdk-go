# services/groups

Documentation for the Groups endpoints (create, get, list, update, delete).

## Create a group

Endpoint

POST https://{youraccountname}.api-us1.com/api/3/groups

Request body

Content type: application/json

Example (minimal):

```json
{
  "group": {
    "title": "My Groups Title"
  }
}
```

Fields

| Field               | Type    | Default | Description                                         |
| ------------------- | ------- | ------- | --------------------------------------------------- |
| `title`             | string  | n/a     | Title of the group (required when creating)         |
| `descript`          | string  | ""      | Group description                                   |
| `reqApprovalNotify` | string  | ""      | Email address to notify for approval-related issues |
| `reqApproval`       | boolean | false   | Whether campaigns require approval                  |
| `reqApproval1st`    | boolean | false   | Whether the first campaign requires approval        |

Permission flags

The group exposes many boolean permission flags (default: `false`) that control features and access. The table below lists the common flags and a short description.

| Flag                      | Description                                              |
| ------------------------- | -------------------------------------------------------- |
| `pgMessageAdd`            | Permission for adding messages                           |
| `pgMessageEdit`           | Permission for editing messages                          |
| `pgMessageDelete`         | Permission for deleting messages                         |
| `pgMessageSend`           | Permission for sending messages                          |
| `pgListAdd`               | Permission for adding lists                              |
| `pgListEdit`              | Permission for editing lists                             |
| `pgListDelete`            | Permission for deleting lists                            |
| `pgListHeaders`           | Permission for managing custom email headers             |
| `pgListEmailaccount`      | Permission for managing Unsubscribe By Email             |
| `pgListBounce`            | Permission for accessing list bounce settings            |
| `pgContactAdd`            | Permission for adding contacts                           |
| `pgContactEdit`           | Permission for editing contacts                          |
| `pgContactDelete`         | Permission for deleting contacts                         |
| `pgContactMerge`          | Permission for merging contacts                          |
| `pgContactImport`         | Permission for importing contacts                        |
| `pgContactApprove`        | Permission for approving contacts                        |
| `pgContactExport`         | Permission for exporting contacts                        |
| `pgContactSync`           | Permission for syncing contacts                          |
| `pgContactFilters`        | Permission for managing contact list segments            |
| `pgContactActions`        | Permission for managing contact actions                  |
| `pgContactFields`         | Permission for managing contact custom fields            |
| `pg_user_add`             | Permission for adding users                              |
| `pg_user_edit`            | Permission for editing users                             |
| `pg_user_delete`          | Permission for deleting users                            |
| `pgGroupAdd`              | Permission for adding groups                             |
| `pgGroupEdit`             | Permission for editing groups                            |
| `pgGroupDelete`           | Permission for deleting groups                           |
| `pgTemplateAdd`           | Permission for adding templates                          |
| `pgTemplateEdit`          | Permission for editing templates                         |
| `pgTemplateDelete`        | Permission for deleting templates                        |
| `pgPersonalizationAdd`    | Permission for adding personalization tags               |
| `pgPersonalizationEdit`   | Permission for editing personalization tags              |
| `pgPersonalizationDelete` | Permission for deleting personalization tags             |
| `pgAutomationManage`      | Permission for managing automations                      |
| `pgFormEdit`              | Permission for editing subscription forms                |
| `pgReportsCampaign`       | Permission for viewing campaign reports                  |
| `pgReportsList`           | Permission for viewing list reports                      |
| `pgReportsUser`           | Permission for viewing user reports                      |
| `pgStartupReports`        | Campaign ID/state used for startup reports               |
| `pgReportsTrend`          | Permission for viewing trend reports                     |
| `pgStartupGettingstarted` | Whether to show the getting-started tutorial on overview |
| `pgDeal`                  | Permission for viewing deals                             |
| `pgDealDelete`            | Permission for deleting deals                            |
| `pgDealReassign`          | Permission for reassigning deals                         |
| `pgDealGroupAdd`          | Permission for adding deal groups                        |
| `pgDealGroupEdit`         | Permission for editing deal groups                       |
| `pgDealGroupDelete`       | Permission for deleting deal groups                      |
| `pgSavedResponsesManage`  | Permission for managing saved responses                  |
| `pgTagManage`             | Permission for managing tags                             |
| `socialdata`              | Controls social link display in campaign emails          |

Example request (full):

```json
{
  "group": {
    "pgMessageAdd": 1,
    "unsubscribelink": "0",
    "optinconfirm": "0",
    "pgListAdd": 1,
    "pgListEdit": 1,
    "pgListDelete": 1,
    "pgListHeaders": 1,
    "pgListEmailaccount": 1,
    "pgListBounce": 1,
    "pgMessageEdit": 1,
    "pgMessageDelete": 1,
    "pgMessageSend": 1,
    "pgContactAdd": 1,
    "pgContactEdit": 1,
    "pgContactDelete": 1,
    "pgContactMerge": 1,
    "pgContactImport": 1,
    "pgContactApprove": 1,
    "pgContactExport": 1,
    "pgContactSync": 1,
    "pgContactFilters": 1,
    "pgContactActions": 1,
    "pgContactFields": 1,
    "pg_user_add": "0",
    "pg_user_edit": "0",
    "pg_user_delete": "0",
    "pgGroupAdd": 1,
    "pgGroupEdit": 1,
    "pgGroupDelete": 1,
    "pgTemplateAdd": 1,
    "pgTemplateEdit": 1,
    "pgTemplateDelete": 1,
    "pgPersonalizationAdd": 1,
    "pgPersonalizationEdit": 1,
    "pgPersonalizationDelete": 1,
    "pgAutomationManage": 1,
    "pgFormEdit": 1,
    "pgReportsCampaign": 1,
    "pgReportsList": 1,
    "pgReportsUser": 1,
    "pgStartupReports": 1,
    "pgReportsTrend": 1,
    "pgStartupGettingstarted": 1,
    "pgDeal": 1,
    "pgDealDelete": 1,
    "pgDealReassign": 1,
    "pgDealGroupAdd": 1,
    "pgDealGroupEdit": 1,
    "pgDealGroupDelete": 1,
    "pgSavedResponsesManage": 1,
    "pgTagManage": false,
    "reqApproval": 1,
    "reqApproval1st": 1,
    "socialdata": 0,
    "title": "TEST TITLE",
    "descript": "Description Text",
    "reqApprovalNotify": "JohnDoe@gmail.com"
  }
}
```

Example response (201/200):

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

cURL example

```bash
curl --request POST \
     --url https://youraccountname.api-us1.com/api/3/groups \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data @- <<'JSON'
{
  "group": {
    "title": "TEST TITLE",
    "descript": "Description Text",
    "reqApprovalNotify": "JohnDoe@gmail.com"
  }
}
JSON
```

Success

200 OK (or 201 Created depending on the API)

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

### Retrieve a group

GET https://{youraccountname}.api-us1.com/api/3/groups/{id}

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

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/groups/id \
     --header 'accept: application/json'
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
  "message": "No Result found for SubscriberSeries with id 3"
}
```

## Update a group

PUT https://{youraccountname}.api-us1.com/api/3/groups/{id}

### Path parameters

| Name | Type  | Required | Description               |
| ---- | ----- | -------- | ------------------------- |
| `id` | int32 | yes      | ID of the group to update |

### Body (group)

The request body accepts a `group` object. Core top-level fields are listed below; most permission-related boolean flags are described in the "Permission flags" table above.

| Field               | Type    | Default | Description                                  |
| ------------------- | ------- | ------- | -------------------------------------------- |
| `title`             | string  | n/a     | Group title                                  |
| `descript`          | string  | ""      | Group description                            |
| `unsubscribelink`   | boolean | false   | Force unsubscribe links                      |
| `optinconfirm`      | boolean | false   | Force opt-in confirmation for this group     |
| `reqApproval`       | boolean | false   | Whether campaigns require approval           |
| `reqApproval1st`    | boolean | false   | Whether the first campaign requires approval |
| `reqApprovalNotify` | string  | ""      | Email to notify for approval issues          |

For permission-style flags (for example `pgListAdd`, `pgMessageSend`, `pgContactAdd`), see the "Permission flags" table above. Set these to `1`/`0` or `true`/`false` depending on the API flavor you use.

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

```bash
curl --request PUT \
     --url https://youraccountname.api-us1.com/api/3/groups/id \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "group": {
    "title": "TEST TITLE",
    "descript": "Description Text",
    "unsubscribelink": "0",
    "optinconfirm": "0",
    "pgListAdd": 1,
    "pgListEdit": 1,
    "pgListDelete": 1,
    "pgListHeaders": 1,
    "pgListEmailaccount": 1,
    "pgListBounce": 1,
    "pgMessageAdd": 1,
    "pgMessageEdit": 1,
    "pgMessageDelete": 1,
    "pgMessageSend": 1,
    "pgContactAdd": 1,
    "pgContactEdit": 1,
    "pgContactDelete": 1,
    "pgContactMerge": 1,
    "pgContactImport": 1,
    "pgContactApprove": 1,
    "pgContactExport": 1,
    "pgContactSync": 1,
    "pgContactFilters": 1,
    "pgContactActions": 1,
    "pgContactFields": 1,
    "pg_user_add": "0",
    "pg_user_edit": "0",
    "pg_user_delete": "0",
    "pgGroupAdd": 1,
    "pgGroupEdit": 1,
    "pgGroupDelete": 1,
    "pgTemplateAdd": 1,
    "pgTemplateEdit": 1,
    "pgTemplateDelete": 1,
    "pgPersonalizationAdd": 1,
    "pgPersonalizationEdit": 1,
    "pgPersonalizationDelete": 1,
    "pgAutomationManage": 1,
    "pgFormEdit": 1,
    "pgReportsCampaign": 1,
    "pgReportsList": 1,
    "pgReportsUser": 1,
    "pgReportsTrend": 1,
    "pgStartupReports": 1,
    "pgStartupGettingstarted": 1,
    "pgDeal": 1,
    "pgDealDelete": 1,
    "pgDealReassign": 1,
    "pgDealGroupAdd": 1,
    "pgDealGroupEdit": 1,
    "pgDealGroupDelete": 1,
    "pgSavedResponsesManage": 1,
    "reqApproval": 1,
    "reqApproval1st": 1,
    "reqApprovalNotify": "JohnDoe@gmail.com",
    "socialdata": 0
  }
}
'
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

## Delete a group

DELETE https://{youraccountname}.api-us1.com/api/3/groups/{id}

```bash
curl --request DELETE \
     --url https://youraccountname.api-us1.com/api/3/groups/id \
     --header 'accept: application/json'
```

200 OK

```json
{}
```

400 Bad Request

```json
{}
```

## List all groups

GET https://{youraccountname}.api-us1.com/api/3/groups

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

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/groups \
     --header 'accept: application/json'
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

## List All Group Limits

GET https://{youraccountname}.api-us1.com/api/3/groupLimits

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

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/groupLimits \
     --header 'accept: application/json'
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
