package campaigns

// CampaignStatus represents the sending status of a campaign as returned by
// the ActiveCampaign API. We use a small typed enum to make code more
// self-documenting when checking or setting campaign status fields.
type CampaignStatus int

const (
	// CampaignStatusDraft indicates the campaign is a draft and not scheduled.
	CampaignStatusDraft CampaignStatus = 0
	// CampaignStatusScheduled indicates the campaign has been scheduled for sending.
	CampaignStatusScheduled CampaignStatus = 1
	// CampaignStatusSending indicates the campaign is currently being sent.
	CampaignStatusSending CampaignStatus = 2
	// CampaignStatusPaused indicates sending has been paused.
	CampaignStatusPaused CampaignStatus = 3
	// CampaignStatusStopped indicates sending was stopped before completion.
	CampaignStatusStopped CampaignStatus = 4
	// CampaignStatusCompleted indicates the campaign has finished sending.
	CampaignStatusCompleted CampaignStatus = 5
)

// String returns a human-friendly name for the CampaignStatus.
func (s CampaignStatus) String() string {
	switch s {
	case CampaignStatusDraft:
		return "Draft"
	case CampaignStatusScheduled:
		return "Scheduled"
	case CampaignStatusSending:
		return "Sending"
	case CampaignStatusPaused:
		return "Paused"
	case CampaignStatusStopped:
		return "Stopped"
	case CampaignStatusCompleted:
		return "Completed"
	default:
		return "Unknown"
	}
}
