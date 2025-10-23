package campaigns

// Campaign represents a campaign object as returned by the ActiveCampaign API.
// Many API fields are string-encoded (including numeric values) and some
// fields can be null; to keep unmarshalling robust we represent most fields
// as strings or pointer-to-string for nullable values. Add or tighten types
// later when a stricter mapping is desired.
type Campaign struct {
	Type                  string            `json:"type,omitempty"`
	UserID                string            `json:"userid,omitempty"`
	SegmentID             string            `json:"segmentid,omitempty"`
	BounceID              string            `json:"bounceid,omitempty"`
	RealCID               string            `json:"realcid,omitempty"`
	SendID                string            `json:"sendid,omitempty"`
	ThreadID              string            `json:"threadid,omitempty"`
	SeriesID              string            `json:"seriesid,omitempty"`
	FormID                string            `json:"formid,omitempty"`
	BaseTemplateID        string            `json:"basetemplateid,omitempty"`
	BaseMessageID         string            `json:"basemessageid,omitempty"`
	AddressID             string            `json:"addressid,omitempty"`
	Source                string            `json:"source,omitempty"`
	Name                  string            `json:"name,omitempty"`
	CDate                 string            `json:"cdate,omitempty"`
	MDate                 string            `json:"mdate,omitempty"`
	SDate                 *string           `json:"sdate,omitempty"`
	LDate                 *string           `json:"ldate,omitempty"`
	SendAmt               string            `json:"send_amt,omitempty"`
	TotalAmt              string            `json:"total_amt,omitempty"`
	Opens                 string            `json:"opens,omitempty"`
	UniqueOpens           string            `json:"uniqueopens,omitempty"`
	LinkClicks            string            `json:"linkclicks,omitempty"`
	UniqueLinkClicks      string            `json:"uniquelinkclicks,omitempty"`
	SubscriberClicks      string            `json:"subscriberclicks,omitempty"`
	Forwards              string            `json:"forwards,omitempty"`
	UniqueForwards        string            `json:"uniqueforwards,omitempty"`
	HardBounces           string            `json:"hardbounces,omitempty"`
	SoftBounces           string            `json:"softbounces,omitempty"`
	Unsubscribes          string            `json:"unsubscribes,omitempty"`
	UnsubReasons          string            `json:"unsubreasons,omitempty"`
	Updates               string            `json:"updates,omitempty"`
	SocialShares          string            `json:"socialshares,omitempty"`
	Replies               string            `json:"replies,omitempty"`
	UniqueReplies         string            `json:"uniquereplies,omitempty"`
	Status                string            `json:"status,omitempty"`
	Public                string            `json:"public,omitempty"`
	MailTransfer          string            `json:"mail_transfer,omitempty"`
	MailSend              string            `json:"mail_send,omitempty"`
	MailCleanup           string            `json:"mail_cleanup,omitempty"`
	MailerLogFile         string            `json:"mailer_log_file,omitempty"`
	TrackLinks            string            `json:"tracklinks,omitempty"`
	TrackLinksAnalytics   string            `json:"tracklinksanalytics,omitempty"`
	TrackReads            string            `json:"trackreads,omitempty"`
	TrackReadsAnalytics   string            `json:"trackreadsanalytics,omitempty"`
	AnalyticsCampaignName string            `json:"analytics_campaign_name,omitempty"`
	Tweet                 string            `json:"tweet,omitempty"`
	Facebook              string            `json:"facebook,omitempty"`
	Survey                string            `json:"survey,omitempty"`
	EmbedImages           string            `json:"embed_images,omitempty"`
	HTMLUnsub             string            `json:"htmlunsub,omitempty"`
	TextUnsub             string            `json:"textunsub,omitempty"`
	HTMLUnsubData         *string           `json:"htmlunsubdata,omitempty"`
	TextUnsubData         *string           `json:"textunsubdata,omitempty"`
	Recurring             string            `json:"recurring,omitempty"`
	WillRecur             string            `json:"willrecur,omitempty"`
	SplitType             string            `json:"split_type,omitempty"`
	SplitContent          string            `json:"split_content,omitempty"`
	SplitOffset           string            `json:"split_offset,omitempty"`
	SplitOffsetType       string            `json:"split_offset_type,omitempty"`
	SplitWinnerMessageID  string            `json:"split_winner_messageid,omitempty"`
	SplitWinnerAwaiting   string            `json:"split_winner_awaiting,omitempty"`
	ResponderOffset       string            `json:"responder_offset,omitempty"`
	ResponderType         string            `json:"responder_type,omitempty"`
	ResponderExisting     string            `json:"responder_existing,omitempty"`
	ReminderField         string            `json:"reminder_field,omitempty"`
	ReminderFormat        *string           `json:"reminder_format,omitempty"`
	ReminderType          string            `json:"reminder_type,omitempty"`
	ReminderOffset        string            `json:"reminder_offset,omitempty"`
	ReminderOffsetType    string            `json:"reminder_offset_type,omitempty"`
	ReminderOffsetSign    string            `json:"reminder_offset_sign,omitempty"`
	ReminderLastCronRun   *string           `json:"reminder_last_cron_run,omitempty"`
	ActiveRSSInterval     string            `json:"activerss_interval,omitempty"`
	ActiveRSSURL          *string           `json:"activerss_url,omitempty"`
	ActiveRSSItems        string            `json:"activerss_items,omitempty"`
	IP4                   string            `json:"ip4,omitempty"`
	LastStep              string            `json:"laststep,omitempty"`
	ManageText            string            `json:"managetext,omitempty"`
	Schedule              string            `json:"schedule,omitempty"`
	ScheduledDate         *string           `json:"scheduleddate,omitempty"`
	WaitPreview           string            `json:"waitpreview,omitempty"`
	DeleteStamp           *string           `json:"deletestamp,omitempty"`
	ReplySys              string            `json:"replysys,omitempty"`
	CreatedTimestamp      string            `json:"created_timestamp,omitempty"`
	UpdatedTimestamp      string            `json:"updated_timestamp,omitempty"`
	CreatedBy             *string           `json:"created_by,omitempty"`
	UpdatedBy             *string           `json:"updated_by,omitempty"`
	IP                    string            `json:"ip,omitempty"`
	SeriesSendLockTime    *string           `json:"series_send_lock_time,omitempty"`
	CanSkipApproval       string            `json:"can_skip_approval,omitempty"`
	UseQuartzScheduler    string            `json:"use_quartz_scheduler,omitempty"`
	VerifiedOpens         string            `json:"verified_opens,omitempty"`
	VerifiedUniqueOpens   string            `json:"verified_unique_opens,omitempty"`
	SegmentName           string            `json:"segmentname,omitempty"`
	HasPredictiveContent  string            `json:"has_predictive_content,omitempty"`
	MessageID             string            `json:"message_id,omitempty"`
	Screenshot            string            `json:"screenshot,omitempty"`
	CampaignMessageID     string            `json:"campaign_message_id,omitempty"`
	EdVersion             string            `json:"ed_version,omitempty"`
	Links                 map[string]string `json:"links,omitempty"`
	ID                    string            `json:"id,omitempty"`
	User                  string            `json:"user,omitempty"`
	Automation            interface{}       `json:"automation,omitempty"`
}

// ListCampaignsResponse is a simple wrapper matching the V3 campaigns list
// response which contains an array of campaigns and optional meta info.
type ListCampaignsResponse struct {
	Campaigns []Campaign `json:"campaigns,omitempty"`
	Meta      struct {
		Total string `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}

// CampaignLink represents an individual link associated with a campaign.
type CampaignLink struct {
	CampaignID string            `json:"campaignid,omitempty"`
	MessageID  *string           `json:"messageid,omitempty"`
	Link       string            `json:"link,omitempty"`
	Name       string            `json:"name,omitempty"`
	Ref        string            `json:"ref,omitempty"`
	Tracked    string            `json:"tracked,omitempty"`
	Links      map[string]string `json:"links,omitempty"`
	ID         string            `json:"id,omitempty"`
	Campaign   string            `json:"campaign,omitempty"`
	Message    *string           `json:"message,omitempty"`
}

// CampaignLinksResponse represents the response for GET /campaigns/{id}/links
type CampaignLinksResponse struct {
	Links []CampaignLink `json:"links,omitempty"`
}
