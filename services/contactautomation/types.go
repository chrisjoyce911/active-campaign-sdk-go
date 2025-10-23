package contactautomation

// ContactAutomationPayload represents a contact-automation relationship.
type ContactAutomationPayload struct {
	ID         string `json:"id,omitempty"`
	Contact    string `json:"contact,omitempty"`
	Automation string `json:"automation,omitempty"`
	Status     string `json:"status,omitempty"`
}

// CreateContactAutomationRequest is the request envelope for adding a contact to an automation.
type CreateContactAutomationRequest struct {
	ContactAutomation ContactAutomationPayload `json:"contactAutomation"`
}

// ContactAutomationResponse is the envelope returned for single contact-automation endpoints.
type ContactAutomationResponse struct {
	ContactAutomation ContactAutomationPayload `json:"contactAutomation"`
}

// ListContactAutomationsResponse represents a list response for contact automations.
type ListContactAutomationsResponse struct {
	ContactAutomations *[]ContactAutomationPayload `json:"contactAutomations"`
}

// AutomationCount represents a simple automation count.
type AutomationCount struct {
	Automation string `json:"automation"`
	Count      int    `json:"count"`
}

// AutomationCountsResponse is the envelope for counts endpoint.
type AutomationCountsResponse struct {
	Counts *[]AutomationCount `json:"counts"`
}
