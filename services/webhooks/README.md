# services/webhooks

## Webhooks

Webhooks provide the ability to receive real-time data updates about your various ActiveCampaign events.

You may choose to receive data based on certain actions (contact subscribes, contact unsubscribes, campaign opens, deal adds, SMS sends, etc..) and have applicable data sent to a URL of your choice. You can then use your own custom code to read, save, and do whatever you want with that data. This is a powerful option that allows you to keep all of your data in sync and opens up various integration options.

With every webhook you create, you can choose when it should actually fire. Perhaps you only want to receive data when a contact is added from the API. You can simply specify subscribe as the event and api as the source when you create your webhook. You can specify multiple events and sources for each webhook if you wish. All event and source options are listed below.

We guarantee at least once delivery on webhooks. In some rare cases, you may receive a webhook event more than once, so it’s important to create an idempotent system. see link for example.

Webhooks are never retried.

Webhook payload fields are listed in this document: webhook payloads.

### Events

| Event                 | Description              |
| --------------------- | ------------------------ |
| `forward`             | Campaign forwarded       |
| `open`                | Campaign opened          |
| `share`               | Campaign shared          |
| `sent`                | Campaign started sending |
| `subscribe`           | Contact subscribed/added |
| `subscriber_note`     | Contact note created     |
| `contact_tag_added`   | Contact tag added        |
| `contact_tag_removed` | Contact tag removed      |
| `unsubscribe`         | Contact unsubscribed     |
| `update`              | Contact updated          |
| `deal_add`            | Deal created             |
| `deal_note_add`       | Deal note created        |
| `deal_pipeline_add`   | Deal pipeline added      |
| `deal_stage_add`      | Deal stage added         |
| `deal_task_add`       | Deal task created        |
| `deal_task_complete`  | Deal task completed      |
| `deal_tasktype_add`   | Deal task type added     |
| `deal_update`         | Deal updated             |
| `bounce`              | Email bounced            |
| `reply`               | Email reply received     |
| `click`               | Link clicked in campaign |
| `list_add`            | New list created         |
| `sms_reply`           | SMS reply received       |
| `sms_sent`            | SMS sent                 |
| `sms_unsub`           | SMS unsubscribe via SMS  |

Sources

| Source   | Description                           |
| -------- | ------------------------------------- |
| `public` | Triggered by public/contact action    |
| `admin`  | Triggered by user in the admin UI     |
| `api`    | Triggered by an API call              |
| `system` | Triggered by automated/system actions |

### Create a webhook

POST https://{youraccountname}.api-us1.com/api/3/webhooks

Create a new webhook

POST /webhooks (Example REQUEST)

Body parameters

| Field     | Type          | Required | Description                                                                                                                                                                                                                                            |
| --------- | ------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `name`    | string        | yes      | A name (label) for this webhook.                                                                                                                                                                                                                       |
| `url`     | string        | yes      | Destination URL for webhook POSTs (HTTPS endpoints must be on port 443).                                                                                                                                                                               |
| `events`  | array[string] | yes      | List of events that will trigger this webhook (see Events table above).                                                                                                                                                                                |
| `sources` | array[string] | yes      | Sources that trigger the webhook (`public`, `admin`, `api`, `system`).                                                                                                                                                                                 |
| `listid`  | int32         | no       | Optional list id to scope events to a single list. Note: `listid` has no effect for certain events (for example: `subscriber_note`, `contact_tag_added`, `contact_tag_removed`, `contact_task_add`, `list_add`, `sms_*`, `deal_*`, `account_*`, etc.). |

Example request body

```json
{
  "webhook": {
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"]
  }
}
```

POST /webhooks (Example RESPONSE)

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"],
    "links": [],
    "id": "1"
  }
}
```

```bash
curl --request POST \
     --url https://youraccountname.api-us1.com/api/3/webhooks \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "webhook": {
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": [
      "subscribe",
      "unsubscribe",
      "sent"
    ],
    "sources": [
      "public",
      "system"
    ]
  }
}
'
```

201 Created

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"],
    "links": [],
    "id": "1"
  }
}
```

### Retrieve a webhook

GET https://{youraccountname}.api-us1.com/api/3/webhooks/{id}

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["admin", "api", "system"],
    "links": [],
    "id": "1"
  }
}
```

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/webhooks/id \
     --header 'accept: application/json'
```

200 OK

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["admin", "api", "system"],
    "links": [],
    "id": "1"
  }
}
```

### Update a webhook

PUT https://{youraccountname}.api-us1.com/api/3/webhooks/{id}

Path parameters

| Name | Type  | Required | Description                 |
| ---- | ----- | -------- | --------------------------- |
| `id` | int32 | yes      | ID of the webhook to update |

Body parameters

| Field     | Type          | Required | Description                                  |
| --------- | ------------- | -------- | -------------------------------------------- |
| `name`    | string        | no       | A name (label) for the webhook.              |
| `url`     | string        | no       | Destination URL for webhook POSTs.           |
| `events`  | array[string] | no       | Events that will trigger the webhook.        |
| `sources` | array[string] | no       | Sources that trigger the webhook.            |
| `listid`  | int32         | no       | Optional list id (required for some events). |

PUT /webhooks/:id (Example REQUEST)

```json
{
  "webhook": {
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"]
  }
}
```

PUT /webhooks/:id (Example RESPONSE)

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"],
    "links": [],
    "id": "1"
  }
}
```

```bash
curl --request PUT \
     --url https://youraccountname.api-us1.com/api/3/webhooks/id \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "webhook": {
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": [
      "subscribe",
      "unsubscribe",
      "sent"
    ],
    "sources": [
      "public",
      "system"
    ]
  }
}
'
```

200 OK

```json
{
  "webhook": {
    "cdate": "2016-01-01T12:00:00-00:00",
    "listid": "0",
    "name": "My Hook",
    "url": "http://example.com/my-hook",
    "events": ["subscribe", "unsubscribe", "sent"],
    "sources": ["public", "system"],
    "links": [],
    "id": "1"
  }
}
```

400 Bad Request

```json
{}
```

### Delete a webhook

DELETE https://{youraccountname}.api-us1.com/api/3/webhooks/{id}

Delete an existing webhook

```bash
curl --request DELETE \
     --url https://youraccountname.api-us1.com/api/3/webhooks/id \
     --header 'accept: application/json'
```

200 OK

```json
{}
```

### List all webhooks

GET https://{youraccountname}.api-us1.com/api/3/webhooks

List all existing webhooks

Query parameters

| Query             | Type   | Description                            |
| ----------------- | ------ | -------------------------------------- |
| `filters[name]`   | string | Filter by webhook name                 |
| `filters[url]`    | string | Filter by webhook url                  |
| `filters[listid]` | string | Filter by webhook's associated list id |

filters[listid]
string
Filter by webhook's associated list

```json
{
  "webhooks": [
    {
      "cdate": "2016-01-01T12:00:00-00:00",
      "listid": "0",
      "name": "My Hook",
      "url": "http://example.com/my-hook",
      "events": ["subscribe", "unsubscribe", "sent"],
      "sources": ["public", "system"],
      "links": [],
      "id": "1"
    },
    {
      "cdate": "2016-01-01T12:00:00-00:00",
      "listid": "0",
      "name": "My Hook 2",
      "url": "http://example.com/my-hook-2",
      "events": ["subscribe"],
      "sources": ["admin"],
      "links": [],
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
     --url 'https://youraccountname.api-us1.com/api/3/webhooks?filters[name]=filters&filters[url]=filters&filters[listid]=filters' \
     --header 'accept: application/json'
```

200 OK

```json
{
  "webhooks": [
    {
      "cdate": "2016-01-01T12:00:00-00:00",
      "listid": "0",
      "name": "My Hook",
      "url": "http://example.com/my-hook",
      "events": ["subscribe", "unsubscribe", "sent"],
      "sources": ["public", "system"],
      "links": [],
      "id": "1"
    },
    {
      "cdate": "2016-01-01T12:00:00-00:00",
      "listid": "0",
      "name": "My Hook 2",
      "url": "http://example.com/my-hook-2",
      "events": ["subscribe"],
      "sources": ["admin"],
      "links": [],
      "id": "2"
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

### List all webhook events

GET https://{youraccountname}.api-us1.com/api/3/webhook/events

List all available webhook events

```json
{
  "webhookEvents": [
    "bounce",
    "click",
    "contact_tag_added",
    "contact_tag_removed",
    "deal_add",
    "deal_note_add",
    "deal_pipeline_add",
    "deal_stage_add",
    "deal_task_add",
    "deal_task_complete",
    "deal_tasktype_add",
    "deal_update",
    "forward",
    "list_add",
    "open",
    "reply",
    "sent",
    "share",
    "sms_reply",
    "sms_sent",
    "sms_unsub",
    "subscribe",
    "subscriber_note",
    "unsubscribe",
    "update"
  ],
  "meta": {
    "total": 25
  }
}
```

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/webhook/events \
     --header 'accept: application/json'
```

200 OK

```json
{
  "webhookEvents": [
    "bounce",
    "click",
    "contact_tag_added",
    "contact_tag_removed",
    "deal_add",
    "deal_note_add",
    "deal_pipeline_add",
    "deal_stage_add",
    "deal_task_add",
    "deal_task_complete",
    "deal_tasktype_add",
    "deal_update",
    "forward",
    "list_add",
    "open",
    "reply",
    "sent",
    "share",
    "sms_reply",
    "sms_sent",
    "sms_unsub",
    "subscribe",
    "subscriber_note",
    "unsubscribe",
    "update"
  ],
  "meta": {
    "total": 25
  }
}
```

## Webhook payloads (representative)

You can choose which events to receive. Each webhook delivers a POST request to your configured URL with event-specific fields. Below are representative payload tables for common events — if you want tables for the remaining events I can add them.

### Account created (type: `account*add`)

| Field                         | Description                                           |
| ----------------------------- | ----------------------------------------------------- |
| `type`                        | Event type (e.g. `account*add`)                       |
| `date_time`                   | Timestamp when the webhook was triggered              |
| `initiated_from`              | Who triggered the event (`system`, `public`, `admin`) |
| `initiated_by`                | Source/section that triggered the event               |
| `list`                        | Always `0` for account webhooks                       |
| `account[id]`                 | Account ID                                            |
| `account[name]`               | Account name                                          |
| `account[account_url]`        | Account URL                                           |
| `account[created_timestamp]`  | Account created timestamp                             |
| `account[updated_timestamp]`  | Account updated timestamp                             |
| `account[fields][<field_id>]` | Account custom fields (indexed by custom field id)    |

### Account updated (type: `account*update`)

| Field                         | Description                                           |
| ----------------------------- | ----------------------------------------------------- |
| `type`                        | Event type (e.g. `account*update`)                    |
| `date_time`                   | Timestamp when the webhook was triggered              |
| `initiated_from`              | Who triggered the event (`system`, `public`, `admin`) |
| `initiated_by`                | Source/section that triggered the event               |
| `list`                        | Always `0` for account webhooks                       |
| `account[id]`                 | Account ID                                            |
| `account[name]`               | Account name                                          |
| `account[account_url]`        | Account URL                                           |
| `account[created_timestamp]`  | Account created timestamp                             |
| `account[updated_timestamp]`  | Account updated timestamp                             |
| `account[fields][<field_id>]` | Account custom fields (indexed by custom field id)    |

### Contact added to account (type: `account*contact_add`)

| Field                         | Description                                           |
| ----------------------------- | ----------------------------------------------------- |
| `type`                        | Event type (e.g. `account*contact_add`)               |
| `date_time`                   | Timestamp when the webhook was triggered              |
| `initiated_from`              | Who triggered the event (`system`, `public`, `admin`) |
| `initiated_by`                | Source/section that triggered the event               |
| `list`                        | Always `0` for this webhook                           |
| `contact[id]`                 | Contact system ID                                     |
| `contact[email]`              | Contact email                                         |
| `contact[first_name]`         | Contact first name                                    |
| `contact[last_name]`          | Contact last name                                     |
| `contact[phone]`              | Contact phone number                                  |
| `contact[tags]`               | Comma-separated list of contact tags                  |
| `contact[orgname]`            | Contact's organization name                           |
| `contact[ip]`                 | Contact IP address                                    |
| `contact[fields][<field_id>]` | Contact custom fields (indexed by custom field id)    |
| `customer_account_id`         | Customer's account id (if present)                    |

### Contact subscribed (type: `subscribe`)

| Field                         | Description                                             |
| ----------------------------- | ------------------------------------------------------- |
| `type`                        | Event type (`subscribe`)                                |
| `date_time`                   | Timestamp when the webhook was triggered                |
| `initiated_by`                | Source/section that triggered the event                 |
| `list`                        | List id(s) associated with the subscription (often `0`) |
| `form[id]`                    | Subscription form id (if applicable)                    |
| `contact[id]`                 | Contact system ID                                       |
| `contact[email]`              | Contact email                                           |
| `contact[first_name]`         | Contact first name                                      |
| `contact[last_name]`          | Contact last name                                       |
| `contact[phone]`              | Contact phone number (for SMS flows)                    |
| `contact[tags]`               | Comma-separated list of contact tags                    |
| `contact[ip]`                 | Contact IP address                                      |
| `contact[fields][<field_id>]` | Contact custom fields                                   |

### Contact unsubscribed (type: `unsubscribe`)

| Field                 | Description                                                          |
| --------------------- | -------------------------------------------------------------------- |
| `type`                | Event type (`unsubscribe`)                                           |
| `date_time`           | Timestamp when the webhook was triggered                             |
| `initiated_by`        | Source/section that triggered the event                              |
| `list`                | List id (if initiated from admin) or `list[][id]` for public actions |
| `campaign[id]`        | Campaign id associated with the unsubscribe (if applicable)          |
| `unsubscribe[reason]` | Reason for unsubscribing (if supplied)                               |
| `contact[id]`         | Contact system ID                                                    |
| `contact[email]`      | Contact email                                                        |

### Contact updated (type: `update`)

| Field                         | Description                                           |
| ----------------------------- | ----------------------------------------------------- |
| `type`                        | Event type (`update`)                                 |
| `date_time`                   | Timestamp when the webhook was triggered              |
| `initiated_by`                | Source/section that triggered the event               |
| `contact[id]`                 | Contact system ID                                     |
| `contact[email]`              | Contact email                                         |
| `contact[first_name]`         | Contact first name                                    |
| `contact[last_name]`          | Contact last name                                     |
| `contact[fields][<field_id>]` | Contact custom fields (updated values may be present) |

Note: the webhook payload field names use bracket notation (for example `contact[id]` or `account[fields][123]`) to represent nested data. The tables above show the most-common fields for representative events; many other events follow similar patterns (prefixed keys, nested arrays for custom fields and links)
