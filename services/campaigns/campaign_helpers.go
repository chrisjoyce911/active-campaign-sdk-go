package campaigns

import (
	"fmt"
	"strconv"
)

// StatusInt parses the Campaign.Status string into an int.
// Returns an error if the status is not a valid integer.
func (c *Campaign) StatusInt() (int, error) {
	if c == nil {
		return 0, nil
	}
	if c.Status == "" {
		return 0, fmt.Errorf("campaign status is empty")
	}
	i, err := strconv.Atoi(c.Status)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// StatusEnum returns the CampaignStatus corresponding to the Campaign.Status string.
// If parsing fails or status is unknown, it returns CampaignStatusDraft as a safe default.
func (c *Campaign) StatusEnum() CampaignStatus {
	if c == nil {
		return CampaignStatusDraft
	}
	i, err := c.StatusInt()
	if err != nil {
		return CampaignStatusDraft
	}
	switch i {
	case 0:
		return CampaignStatusDraft
	case 1:
		return CampaignStatusScheduled
	case 2:
		return CampaignStatusSending
	case 3:
		return CampaignStatusPaused
	case 4:
		return CampaignStatusStopped
	case 5:
		return CampaignStatusCompleted
	default:
		return CampaignStatusDraft
	}
}
