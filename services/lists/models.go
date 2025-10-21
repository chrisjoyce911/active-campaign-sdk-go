package lists

import (
	"encoding/json"
	"strconv"
)

// Typed models for Lists API requests/responses.

// StringOrInt is a helper that accepts either a JSON number or string and
// canonicalizes it to a string value.
type StringOrInt string

func (s *StringOrInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*s = ""
		return nil
	}
	// If quoted string
	if b[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*s = StringOrInt(str)
		return nil
	}
	// Otherwise treat as number
	var num json.Number
	if err := json.Unmarshal(b, &num); err != nil {
		return err
	}
	*s = StringOrInt(num.String())
	return nil
}

// IntOrString accepts either a JSON number or string and canonicalizes to int.
type IntOrString int

func (i *IntOrString) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*i = 0
		return nil
	}
	if b[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		v, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*i = IntOrString(v)
		return nil
	}
	// number
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		// try via json.Number to preserve large ints
		var num json.Number
		if err2 := json.Unmarshal(b, &num); err2 != nil {
			return err
		}
		iv, err3 := strconv.Atoi(num.String())
		if err3 != nil {
			return err3
		}
		*i = IntOrString(iv)
		return nil
	}
	*i = IntOrString(v)
	return nil
}

// List represents the fields of an ActiveCampaign list returned by the API.
type List struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	StringID          string `json:"stringid,omitempty"`
	Channel           string `json:"channel,omitempty"`
	SenderURL         string `json:"sender_url,omitempty"`
	SenderReminder    string `json:"sender_reminder,omitempty"`
	SenderName        string `json:"sender_name,omitempty"`
	SenderAddr1       string `json:"sender_addr1,omitempty"`
	SenderAddr2       string `json:"sender_addr2,omitempty"`
	SenderCity        string `json:"sender_city,omitempty"`
	SenderState       string `json:"sender_state,omitempty"`
	SenderZip         string `json:"sender_zip,omitempty"`
	SenderCountry     string `json:"sender_country,omitempty"`
	SenderPhone       string `json:"sender_phone,omitempty"`
	RequireName       string `json:"require_name,omitempty"`
	OptInOptOut       string `json:"optinoptout,omitempty"`
	PUseTracking      string `json:"p_use_tracking,omitempty"`
	PUseAnalyticsRead string `json:"p_use_analytics_read,omitempty"`
	PUseAnalyticsLink string `json:"p_use_analytics_link,omitempty"`
	PUseTwitter       string `json:"p_use_twitter,omitempty"`
	PUseFacebook      string `json:"p_use_facebook,omitempty"`
	EmbedImage        string `json:"p_embed_image,omitempty"`
	UseCaptcha        string `json:"p_use_captcha,omitempty"`
	// send_last_broadcast may be returned as a number or a string by the API;
	// use StringOrInt to canonicalize to string.
	SendLastBroadcast     StringOrInt       `json:"send_last_broadcast,omitempty"`
	Private               string            `json:"private,omitempty"`
	AnalyticsDomains      string            `json:"analytics_domains,omitempty"`
	AnalyticsSource       string            `json:"analytics_source,omitempty"`
	AnalyticsUA           string            `json:"analytics_ua,omitempty"`
	Description           string            `json:"description,omitempty"`
	CDate                 string            `json:"cdate,omitempty"`
	UDate                 string            `json:"udate,omitempty"`
	CreatedTimestamp      string            `json:"created_timestamp,omitempty"`
	UpdatedTimestamp      string            `json:"updated_timestamp,omitempty"`
	NonDeletedSubscribers IntOrString       `json:"non_deleted_subscribers,omitempty"`
	ActiveSubscribers     IntOrString       `json:"active_subscribers,omitempty"`
	Links                 map[string]string `json:"links,omitempty"`
	User                  StringOrInt       `json:"user,omitempty"`
	// Additional fields observed in API responses
	UserID               StringOrInt `json:"userid,omitempty"`
	FullAddress          string      `json:"fulladdress,omitempty"`
	OptinMessageID       string      `json:"optinmessageid,omitempty"`
	OptOutConf           string      `json:"optoutconf,omitempty"`
	DeleteStamp          StringOrInt `json:"deletestamp,omitempty"`
	CreatedBy            StringOrInt `json:"created_by,omitempty"`
	UpdatedBy            StringOrInt `json:"updated_by,omitempty"`
	CarbonCopy           string      `json:"carboncopy,omitempty"`
	SubscriptionNotify   string      `json:"subscription_notify,omitempty"`
	UnsubscriptionNotify string      `json:"unsubscription_notify,omitempty"`
	ToName               string      `json:"to_name,omitempty"`
	TwitterToken         string      `json:"twitter_token,omitempty"`
	TwitterTokenSecret   string      `json:"twitter_token_secret,omitempty"`
	FacebookSession      string      `json:"facebook_session,omitempty"`
}

// CreateListRequest wraps a List for POST /lists
type CreateListRequest struct {
	List List `json:"list"`
}

// CreateListResponse represents the server response for creating a list.
type CreateListResponse struct {
	List List `json:"list"`
}

// GetListResponse represents the server response for retrieving a single list.
type GetListResponse struct {
	List List `json:"list"`
}

// ListsResponse represents the response for listing multiple lists.
type ListsResponse struct {
	Lists []List                 `json:"lists"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}

// ListGroup models for POST /listGroups
type ListGroup struct {
	ListID  int `json:"listid"`
	GroupID int `json:"groupid"`
}

type CreateListGroupRequest struct {
	ListGroup ListGroup `json:"listGroup"`
}

type CreateListGroupResponse struct {
	ListGroup ListGroup `json:"listGroup"`
}
