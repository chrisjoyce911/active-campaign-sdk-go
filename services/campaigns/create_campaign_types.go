package campaigns

// CreateCampaignRequest represents the minimal payload accepted by POST /campaigns
type CreateCampaignRequest struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	CanSplitContent *bool  `json:"canSplitContent,omitempty"`
}
