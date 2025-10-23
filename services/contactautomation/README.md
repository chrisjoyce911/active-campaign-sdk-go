# services/contactautomation

## Add a contact to an automation

POST https://{youraccountname}.api-us1.com/api/3/contactAutomations

POST /contactAutomations (Example REQUEST)

Body Params
contactAutomation
object

contactAutomation object
contact
int32
required
Contact ID of the Contact, to be linked to the contactAutomation

automation
int32
required
Automation ID of the automation, to be linked to the contactAutomation

```json
{
  "contactAutomation": {
    "contact": "117",
    "automation": "42"
  }
}
```

POST /contactAutomations (Example RESPONSE)

```json
{
  "contacts": [
    {
      "cdate": "2018-08-06T16:26:06-05:00",
      "email": "johndoe@example.com",
      "phone": "",
      "firstName": "",
      "lastName": "",
      "orgid": "0",
      "segmentio_id": "",
      "bounced_hard": "0",
      "bounced_soft": "0",
      "bounced_date": null,
      "ip": "2130706433",
      "ua": null,
      "hash": "054aa0acede49e07a844420c879b3c30",
      "socialdata_lastcheck": null,
      "email_local": "",
      "email_domain": "",
      "sentcnt": "0",
      "rating_tstamp": null,
      "gravatar": "0",
      "deleted": "0",
      "anonymized": "0",
      "adate": null,
      "udate": null,
      "edate": null,
      "deleted_at": null,
      "created_utc_timestamp": "2018-09-21 12:04:40",
      "updated_utc_timestamp": "2018-09-21 12:04:40",
      "links": {
        "bounceLogs": "https://:account.api-us1.com/api/:version/contacts/64/bounceLogs",
        "contactAutomations": "https://:account.api-us1.com/api/:version/contacts/64/contactAutomations",
        "contactData": "https://:account.api-us1.com/api/:version/contacts/64/contactData",
        "contactGoals": "https://:account.api-us1.com/api/:version/contacts/64/contactGoals",
        "contactLists": "https://:account.api-us1.com/api/:version/contacts/64/contactLists",
        "contactLogs": "https://:account.api-us1.com/api/:version/contacts/64/contactLogs",
        "contactTags": "https://:account.api-us1.com/api/:version/contacts/64/contactTags",
        "contactDeals": "https://:account.api-us1.com/api/:version/contacts/64/contactDeals",
        "deals": "https://:account.api-us1.com/api/:version/contacts/64/deals",
        "fieldValues": "https://:account.api-us1.com/api/:version/contacts/64/fieldValues",
        "geoIps": "https://:account.api-us1.com/api/:version/contacts/64/geoIps",
        "notes": "https://:account.api-us1.com/api/:version/contacts/64/notes",
        "organization": "https://:account.api-us1.com/api/:version/contacts/64/organization",
        "plusAppend": "https://:account.api-us1.com/api/:version/contacts/64/plusAppend",
        "trackingLogs": "https://:account.api-us1.com/api/:version/contacts/64/trackingLogs",
        "scoreValues": "https://:account.api-us1.com/api/:version/contacts/64/scoreValues"
      },
      "id": "64",
      "organization": null
    }
  ],
  "contactAutomation": {
    "contact": "64",
    "seriesid": "2",
    "startid": 0,
    "status": 1,
    "lastblock": "4",
    "completedElements": "1",
    "totalElements": "2",
    "completed": 0,
    "completeValue": 50,
    "links": {
      "automation": "https://:account.api-us1.com/api/:version/contactAutomations/3/automation",
      "contact": "https://:account.api-us1.com/api/:version/contactAutomations/3/contact",
      "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/3/contactGoals"
    },
    "id": "3",
    "automation": "2"
  }
}
```

```bash
curl --request POST \
     --url https://youraccountname.api-us1.com/api/3/contactAutomations \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "contactAutomation": {
    "contact": "117",
    "automation": "42"
  }
}
'
```

201 Created

```json
{
  "contacts": [
    {
      "cdate": "2018-08-06T16:26:06-05:00",
      "email": "johndoe@example.com",
      "phone": "",
      "firstName": "",
      "lastName": "",
      "orgid": "0",
      "segmentio_id": "",
      "bounced_hard": "0",
      "bounced_soft": "0",
      "bounced_date": null,
      "ip": "2130706433",
      "ua": null,
      "hash": "054aa0acede49e07a844420c879b3c30",
      "socialdata_lastcheck": null,
      "email_local": "",
      "email_domain": "",
      "sentcnt": "0",
      "rating_tstamp": null,
      "gravatar": "0",
      "deleted": "0",
      "anonymized": "0",
      "adate": null,
      "udate": null,
      "edate": null,
      "deleted_at": null,
      "created_utc_timestamp": "2018-09-21 12:04:40",
      "updated_utc_timestamp": "2018-09-21 12:04:40",
      "links": {
        "bounceLogs": "https://:account.api-us1.com/api/:version/contacts/64/bounceLogs",
        "contactAutomations": "https://:account.api-us1.com/api/:version/contacts/64/contactAutomations",
        "contactData": "https://:account.api-us1.com/api/:version/contacts/64/contactData",
        "contactGoals": "https://:account.api-us1.com/api/:version/contacts/64/contactGoals",
        "contactLists": "https://:account.api-us1.com/api/:version/contacts/64/contactLists",
        "contactLogs": "https://:account.api-us1.com/api/:version/contacts/64/contactLogs",
        "contactTags": "https://:account.api-us1.com/api/:version/contacts/64/contactTags",
        "contactDeals": "https://:account.api-us1.com/api/:version/contacts/64/contactDeals",
        "deals": "https://:account.api-us1.com/api/:version/contacts/64/deals",
        "fieldValues": "https://:account.api-us1.com/api/:version/contacts/64/fieldValues",
        "geoIps": "https://:account.api-us1.com/api/:version/contacts/64/geoIps",
        "notes": "https://:account.api-us1.com/api/:version/contacts/64/notes",
        "organization": "https://:account.api-us1.com/api/:version/contacts/64/organization",
        "plusAppend": "https://:account.api-us1.com/api/:version/contacts/64/plusAppend",
        "trackingLogs": "https://:account.api-us1.com/api/:version/contacts/64/trackingLogs",
        "scoreValues": "https://:account.api-us1.com/api/:version/contacts/64/scoreValues"
      },
      "id": "64",
      "organization": null
    }
  ],
  "contactAutomation": {
    "contact": "64",
    "seriesid": "2",
    "startid": 0,
    "status": 1,
    "lastblock": "4",
    "completedElements": "1",
    "totalElements": "2",
    "completed": 0,
    "completeValue": 50,
    "links": {
      "automation": "https://:account.api-us1.com/api/:version/contactAutomations/3/automation",
      "contact": "https://:account.api-us1.com/api/:version/contactAutomations/3/contact",
      "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/3/contactGoals"
    },
    "id": "3",
    "automation": "2"
  }
}
```

403 Forbidden

```json
{
  "message": "Could not create SubscriberSeries"
}
```

## Retrieve an automation a contact is in

GET https://{youraccountname}.api-us1.com/api/3/contactAutomations/{id}

```json
{
  "contactAutomation": {
    "contact": "110",
    "seriesid": "2",
    "startid": "0",
    "status": "2",
    "batchid": null,
    "adddate": "2018-09-19T09:44:26-05:00",
    "remdate": "2018-09-19T09:44:26-05:00",
    "timespan": "0",
    "lastblock": "5",
    "lastlogid": "2",
    "lastdate": "2018-09-19T09:44:26-05:00",
    "completedElements": "1",
    "totalElements": "2",
    "completed": 1,
    "completeValue": 100,
    "links": {
      "automation": "https://:account.api-us1.com/api/:version/contactAutomations/2/automation",
      "contact": "https://:account.api-us1.com/api/:version/contactAutomations/2/contact",
      "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/2/contactGoals"
    },
    "id": "2",
    "automation": "2"
  }
}
```

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/contactAutomations/id \
     --header 'accept: application/json'

```

200 OK

```json
{
  "contactAutomation": {
    "contact": "110",
    "seriesid": "2",
    "startid": "0",
    "status": "2",
    "batchid": null,
    "adddate": "2018-09-19T09:44:26-05:00",
    "remdate": "2018-09-19T09:44:26-05:00",
    "timespan": "0",
    "lastblock": "5",
    "lastlogid": "2",
    "lastdate": "2018-09-19T09:44:26-05:00",
    "completedElements": "1",
    "totalElements": "2",
    "completed": 1,
    "completeValue": 100,
    "links": {
      "automation": "https://:account.api-us1.com/api/:version/contactAutomations/2/automation",
      "contact": "https://:account.api-us1.com/api/:version/contactAutomations/2/contact",
      "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/2/contactGoals"
    },
    "id": "2",
    "automation": "2"
  }
}
```

404 Not Found

```json
{
  "message": "No Result found for SubscriberSeries with id 3"
}
```

## Remove a contact from an automation

DELETE https://{youraccountname}.api-us1.com/api/3/contactAutomations/{id}

Path Params
id
int32
required
ID of the contactAutomation to delete

```bash

curl --request DELETE \
     --url https://youraccountname.api-us1.com/api/3/contactAutomations/id \
     --header 'accept: application/json'
```

200 OK

```json
{}
```

403 Forbidden

```json
{
  "message": "No Result found for SubscriberSeries with id 3"
}
```

## List all automations a contact is in

GET https://{youraccountname}.api-us1.com/api/3/contactAutomations

```json
{
  "contactAutomations": [
    {
      "contact": "10003",
      "seriesid": "1",
      "startid": "0",
      "status": "2",
      "batchid": null,
      "adddate": "2018-11-16T02:32:33-06:00",
      "remdate": "2018-11-16T02:32:33-06:00",
      "timespan": "0",
      "lastblock": "1",
      "lastlogid": "0",
      "lastdate": "2018-11-16T02:32:33-06:00",
      "completedElements": "0",
      "totalElements": "1",
      "completed": 1,
      "completeValue": 100,
      "links": {
        "automation": "https://:account.api-us1.com/api/:version/contactAutomations/1/automation",
        "contact": "https://:account.api-us1.com/api/:version/contactAutomations/1/contact",
        "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/1/contactGoals"
      },
      "id": "1",
      "automation": "1"
    }
  ],
  "meta": {
    "total": "1",
    "showcase_stats": []
  }
}
```

```bash
curl --request GET \
     --url https://youraccountname.api-us1.com/api/3/contactAutomations \
     --header 'accept: application/json'

```

200 OK

```json
{
  "contactAutomations": [
    {
      "contact": "10003",
      "seriesid": "1",
      "startid": "0",
      "status": "2",
      "batchid": null,
      "adddate": "2018-11-16T02:32:33-06:00",
      "remdate": "2018-11-16T02:32:33-06:00",
      "timespan": "0",
      "lastblock": "1",
      "lastlogid": "0",
      "lastdate": "2018-11-16T02:32:33-06:00",
      "completedElements": "0",
      "totalElements": "1",
      "completed": 1,
      "completeValue": 100,
      "links": {
        "automation": "https://:account.api-us1.com/api/:version/contactAutomations/1/automation",
        "contact": "https://:account.api-us1.com/api/:version/contactAutomations/1/contact",
        "contactGoals": "https://:account.api-us1.com/api/:version/contactAutomations/1/contactGoals"
      },
      "id": "1",
      "automation": "1"
    }
  ],
  "meta": {
    "total": "1",
    "showcase_stats": []
  }
}
```
