package campaigns

// DuplicateCampaignResponse represents the API response for copying a campaign.
type DuplicateCampaignResponse struct {
	Succeeded     int    `json:"succeeded,omitempty"`
	Message       string `json:"message,omitempty"`
	NewCampaignID int    `json:"newCampaignId,omitempty"`
}
