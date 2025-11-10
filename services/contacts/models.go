package contacts

// Models for contacts API - these are placeholders and should match ActiveCampaign's API

// CreateContactRequest is the payload when creating a contact.
// TODO: expand fields per https://developers.activecampaign.com/reference#create-contact
type CreateContactRequest struct {
	Contact *Contact `json:"contact"`
}

// Contact represents the contact object.
// TODO: expand fields and tags, fieldValues, links, etc.
type Contact struct {
	ID                  string        `json:"id,omitempty"`
	Email               string        `json:"email,omitempty"`
	Phone               string        `json:"phone,omitempty"`
	FirstName           string        `json:"firstName,omitempty"`
	LastName            string        `json:"lastName,omitempty"`
	OrgID               string        `json:"orgid,omitempty"`
	OrgName             string        `json:"orgname,omitempty"`
	SegmentIOID         string        `json:"segmentio_id,omitempty"`
	BouncedHard         string        `json:"bounced_hard,omitempty"`
	BouncedSoft         string        `json:"bounced_soft,omitempty"`
	BouncedDate         *string       `json:"bounced_date,omitempty"`
	IP                  string        `json:"ip,omitempty"`
	UA                  string        `json:"ua,omitempty"`
	Hash                string        `json:"hash,omitempty"`
	SocialDataLastCheck *string       `json:"socialdata_lastcheck,omitempty"`
	EmailLocal          string        `json:"email_local,omitempty"`
	EmailDomain         string        `json:"email_domain,omitempty"`
	SentCnt             string        `json:"sentcnt,omitempty"`
	RatingTstamp        *string       `json:"rating_tstamp,omitempty"`
	Gravatar            string        `json:"gravatar,omitempty"`
	Deleted             string        `json:"deleted,omitempty"`
	Anonymized          string        `json:"anonymized,omitempty"`
	ADate               string        `json:"adate,omitempty"`
	UDate               string        `json:"udate,omitempty"`
	EDate               string        `json:"edate,omitempty"`
	DeletedAt           *string       `json:"deleted_at,omitempty"`
	CreatedUTCTimestamp string        `json:"created_utc_timestamp,omitempty"`
	UpdatedUTCTimestamp string        `json:"updated_utc_timestamp,omitempty"`
	CreatedTimestamp    string        `json:"created_timestamp,omitempty"`
	UpdatedTimestamp    string        `json:"updated_timestamp,omitempty"`
	CreatedBy           string        `json:"created_by,omitempty"`
	UpdatedBy           string        `json:"updated_by,omitempty"`
	MPPTracking         string        `json:"mpp_tracking,omitempty"`
	LastClickDate       *string       `json:"last_click_date,omitempty"`
	LastOpenDate        *string       `json:"last_open_date,omitempty"`
	LastMPPOpenDate     string        `json:"last_mpp_open_date,omitempty"`
	BestSendHour        string        `json:"best_send_hour,omitempty"`
	ScoreValues         []string      `json:"scoreValues,omitempty"`
	AccountContacts     []interface{} `json:"accountContacts,omitempty"`
	// FieldValues allows creating/updating custom field values inline when
	// creating or updating a contact. Each FieldValue refers to a previously
	// created custom field by id.
	FieldValues  *[]FieldValue     `json:"fieldValues,omitempty"`
	Links        map[string]string `json:"links,omitempty"`
	Organization interface{}       `json:"organization,omitempty"`
}

// CreateContactResponse is returned when a contact is created.
// TODO: align with real response fields.
type CreateContactResponse struct {
	Contact            *Contact             `json:"contact"`
	ContactAutomations *[]ContactAutomation `json:"contactAutomations,omitempty"`
	ContactData        *[]ContactData       `json:"contactData,omitempty"`
	ContactLists       *[]ContactList       `json:"contactLists,omitempty"`
	Deals              *[]interface{}       `json:"deals,omitempty"`
	FieldValues        *[]FieldValue        `json:"fieldValues,omitempty"`
	GeoAddresses       *[]GeoAddress        `json:"geoAddresses,omitempty"`
	GeoIps             *[]GeoIp             `json:"geoIps,omitempty"`
	AccountContacts    *[]interface{}       `json:"accountContacts,omitempty"`
	ScoreValues        *[]ScoreValue        `json:"scoreValues,omitempty"`
}

// ContactAutomation represents a contact's automation entry
type ContactAutomation struct {
	Contact           string            `json:"contact,omitempty"`
	SeriesID          string            `json:"seriesid,omitempty"`
	StartID           string            `json:"startid,omitempty"`
	Status            string            `json:"status,omitempty"`
	BatchID           string            `json:"batchid,omitempty"`
	AddDate           string            `json:"adddate,omitempty"`
	RemDate           *string           `json:"remdate,omitempty"`
	TimeSpan          string            `json:"timespan,omitempty"`
	LastBlock         string            `json:"lastblock,omitempty"`
	LastLogID         string            `json:"lastlogid,omitempty"`
	LastDate          string            `json:"lastdate,omitempty"`
	InALS             string            `json:"in_als,omitempty"`
	CompletedElements int               `json:"completedElements,omitempty"`
	TotalElements     int               `json:"totalElements,omitempty"`
	Completed         int               `json:"completed,omitempty"`
	CompleteValue     int               `json:"completeValue,omitempty"`
	Links             map[string]string `json:"links,omitempty"`
	ID                string            `json:"id,omitempty"`
	Automation        string            `json:"automation,omitempty"`
}

// ContactData represents geo/contact data
type ContactData struct {
	Contact          string        `json:"contact,omitempty"`
	TStamp           string        `json:"tstamp,omitempty"`
	GeoTStamp        *string       `json:"geoTstamp,omitempty"`
	GeoIp4           string        `json:"geoIp4,omitempty"`
	GeoCountry2      string        `json:"geoCountry2,omitempty"`
	GeoCountry       string        `json:"geo_country,omitempty"`
	GeoState         string        `json:"geoState,omitempty"`
	GeoCity          string        `json:"geoCity,omitempty"`
	GeoZip           string        `json:"geoZip,omitempty"`
	GeoArea          string        `json:"geoArea,omitempty"`
	GeoLat           string        `json:"geoLat,omitempty"`
	GeoLon           string        `json:"geoLon,omitempty"`
	GeoTz            string        `json:"geoTz,omitempty"`
	GeoTzOffset      string        `json:"geoTzOffset,omitempty"`
	GAFirstVisit     *string       `json:"ga_first_visit,omitempty"`
	GATimesVisited   string        `json:"ga_times_visited,omitempty"`
	FBID             string        `json:"fb_id,omitempty"`
	FBName           string        `json:"fb_name,omitempty"`
	TWID             string        `json:"tw_id,omitempty"`
	CreatedTimestamp string        `json:"created_timestamp,omitempty"`
	UpdatedTimestamp string        `json:"updated_timestamp,omitempty"`
	CreatedBy        string        `json:"created_by,omitempty"`
	UpdatedBy        string        `json:"updated_by,omitempty"`
	Links            []interface{} `json:"links,omitempty"`
	ID               string        `json:"id,omitempty"`
}

// ContactList represents membership info for a contact on a list
type ContactList struct {
	Contact   string  `json:"contact,omitempty"`
	List      string  `json:"list,omitempty"`
	Form      *string `json:"form,omitempty"`
	SeriesID  string  `json:"seriesid,omitempty"`
	SDate     string  `json:"sdate,omitempty"`
	UDate     *string `json:"udate,omitempty"`
	Status    string  `json:"status,omitempty"`
	Responder string  `json:"responder,omitempty"`
	Sync      string  `json:"sync,omitempty"`
	// Additional fields commonly returned by the API on contact list membership
	UnsubReason           string            `json:"unsubreason,omitempty"`
	Campaign              *string           `json:"campaign,omitempty"`
	Message               *string           `json:"message,omitempty"`
	FirstName             string            `json:"first_name,omitempty"`
	LastName              string            `json:"last_name,omitempty"`
	Ip4Sub                string            `json:"ip4Sub,omitempty"`
	SourceID              string            `json:"sourceid,omitempty"`
	AutoSyncLog           *string           `json:"autosyncLog,omitempty"`
	Ip4Last               string            `json:"ip4_last,omitempty"`
	Ip4Unsub              string            `json:"ip4Unsub,omitempty"`
	CreatedTimestamp      string            `json:"created_timestamp,omitempty"`
	UpdatedTimestamp      string            `json:"updated_timestamp,omitempty"`
	CreatedBy             *string           `json:"created_by,omitempty"`
	UpdatedBy             *string           `json:"updated_by,omitempty"`
	UnsubscribeAutomation *string           `json:"unsubscribeAutomation,omitempty"`
	Links                 map[string]string `json:"links,omitempty"`
	ID                    string            `json:"id,omitempty"`
	Automation            *string           `json:"automation,omitempty"`
}

// ContactListsResponse is returned from GET /contacts/{id}/contactLists
type ContactListsResponse struct {
	ContactLists *[]ContactList `json:"contactLists"`
}

// Accessor helper to avoid nil checks
func (r *ContactListsResponse) ContactListsOrEmpty() []ContactList {
	if r == nil || r.ContactLists == nil {
		return []ContactList{}
	}
	return *r.ContactLists
}

// FieldValue represents a custom field value for a contact
type FieldValue struct {
	Contact   string            `json:"contact,omitempty"`
	Field     string            `json:"field,omitempty"`
	Value     string            `json:"value,omitempty"`
	CDate     string            `json:"cdate,omitempty"`
	UDate     string            `json:"udate,omitempty"`
	CreatedBy string            `json:"created_by,omitempty"`
	UpdatedBy string            `json:"updated_by,omitempty"`
	Links     map[string]string `json:"links,omitempty"`
	ID        string            `json:"id,omitempty"`
	Owner     string            `json:"owner,omitempty"`
}

// GeoAddress represents a geo address record
type GeoAddress struct {
	IP4      string        `json:"ip4,omitempty"`
	Country2 string        `json:"country2,omitempty"`
	Country  string        `json:"country,omitempty"`
	State    string        `json:"state,omitempty"`
	City     string        `json:"city,omitempty"`
	Zip      string        `json:"zip,omitempty"`
	Area     string        `json:"area,omitempty"`
	Lat      string        `json:"lat,omitempty"`
	Lon      string        `json:"lon,omitempty"`
	TZ       string        `json:"tz,omitempty"`
	TStamp   string        `json:"tstamp,omitempty"`
	Links    []interface{} `json:"links,omitempty"`
	ID       string        `json:"id,omitempty"`
}

// GeoIp represents a geo IP entry for a contact
type GeoIp struct {
	Contact    string            `json:"contact,omitempty"`
	CampaignID string            `json:"campaignid,omitempty"`
	MessageID  string            `json:"messageid,omitempty"`
	GeoAddrID  string            `json:"geoaddrid,omitempty"`
	IP4        string            `json:"ip4,omitempty"`
	TStamp     string            `json:"tstamp,omitempty"`
	Links      map[string]string `json:"links,omitempty"`
	ID         string            `json:"id,omitempty"`
	GeoAddress string            `json:"geoAddress,omitempty"`
}

// ContactSearchResponse for searching contacts by email.
// TODO: align with real response.
type ContactSearchResponse struct {
	ScoreValues *[]ScoreValue          `json:"scoreValues,omitempty"`
	Contacts    []Contact              `json:"contacts"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

// ScoreValue represents a contact scoring record
type ScoreValue struct {
	Score      string            `json:"score,omitempty"`
	Contact    string            `json:"contact,omitempty"`
	Deal       interface{}       `json:"deal,omitempty"`
	CDate      string            `json:"cdate,omitempty"`
	MDate      string            `json:"mdate,omitempty"`
	ScoreValue string            `json:"scoreValue,omitempty"`
	Links      map[string]string `json:"links,omitempty"`
	ID         string            `json:"id,omitempty"`
}

// UpdateListStatusForContactRequest - TODO
type UpdateListStatusForContactRequest struct {
	ContactList *ContactList `json:"contactList"`
}

// UpdateContactListStatusResponse - TODO: align with real response
type UpdateContactListStatusResponse struct {
	// ...existing code...
}

// ContactTag represents the mapping of a contact to a tag.
type ContactTagLinks struct {
	Tag     string `json:"tag,omitempty"`
	Contact string `json:"contact,omitempty"`
}

type ContactTag struct {
	Contact          string          `json:"contact,omitempty"`
	Tag              string          `json:"tag,omitempty"`
	CDate            string          `json:"cdate,omitempty"`
	CreatedTimestamp string          `json:"created_timestamp,omitempty"`
	UpdatedTimestamp string          `json:"updated_timestamp,omitempty"`
	CreatedBy        interface{}     `json:"created_by,omitempty"`
	UpdatedBy        interface{}     `json:"updated_by,omitempty"`
	Links            ContactTagLinks `json:"links,omitempty"`
	ID               string          `json:"id,omitempty"`
}

// ContactTagsResponse is returned from GET /contacts/{id}/contactTags
type ContactTagsResponse struct {
	ContactTags *[]ContactTag `json:"contactTags"`
}

// Accessor for ContactTagsResponse
func (c *ContactTagsResponse) ContactTagsOrEmpty() []ContactTag {
	if c == nil || c.ContactTags == nil {
		return []ContactTag{}
	}
	return *c.ContactTags
}

// Accessors for CreateContactResponse to avoid nil checks in callers
func (c *CreateContactResponse) ContactAutomationsOrEmpty() []ContactAutomation {
	if c == nil || c.ContactAutomations == nil {
		return []ContactAutomation{}
	}
	return *c.ContactAutomations
}

func (c *CreateContactResponse) ContactDataOrEmpty() []ContactData {
	if c == nil || c.ContactData == nil {
		return []ContactData{}
	}
	return *c.ContactData
}

func (c *CreateContactResponse) ContactListsOrEmpty() []ContactList {
	if c == nil || c.ContactLists == nil {
		return []ContactList{}
	}
	return *c.ContactLists
}

func (c *CreateContactResponse) FieldValuesOrEmpty() []FieldValue {
	if c == nil || c.FieldValues == nil {
		return []FieldValue{}
	}
	return *c.FieldValues
}

func (c *CreateContactResponse) GeoAddressesOrEmpty() []GeoAddress {
	if c == nil || c.GeoAddresses == nil {
		return []GeoAddress{}
	}
	return *c.GeoAddresses
}

func (c *CreateContactResponse) GeoIpsOrEmpty() []GeoIp {
	if c == nil || c.GeoIps == nil {
		return []GeoIp{}
	}
	return *c.GeoIps
}

func (c *CreateContactResponse) ScoreValuesOrEmpty() []ScoreValue {
	if c == nil || c.ScoreValues == nil {
		return []ScoreValue{}
	}
	return *c.ScoreValues
}

// Accessor for ContactSearchResponse score values
func (c *ContactSearchResponse) ScoreValuesOrEmpty() []ScoreValue {
	if c == nil || c.ScoreValues == nil {
		return []ScoreValue{}
	}
	return *c.ScoreValues
}
