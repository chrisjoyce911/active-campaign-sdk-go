package groups

// Models for Groups API

// Group represents a minimal subset of the Group resource used in requests/responses.
type Group struct {
	ID                string            `json:"id,omitempty"`
	Title             string            `json:"title,omitempty"`
	Descript          string            `json:"descript,omitempty"`
	UnsubscribeLink   string            `json:"unsubscribelink,omitempty"`
	OptinConfirm      string            `json:"optinconfirm,omitempty"`
	ReqApproval       string            `json:"reqApproval,omitempty"`
	ReqApproval1st    string            `json:"reqApproval1st,omitempty"`
	ReqApprovalNotify string            `json:"reqApprovalNotify,omitempty"`
	SocialData        string            `json:"socialdata,omitempty"`
	Links             map[string]string `json:"links,omitempty"`
}

type CreateGroupRequest struct {
	Group Group `json:"group"`
}

type CreateGroupResponse struct {
	Group Group `json:"group"`
}

type GetGroupResponse struct {
	Group Group `json:"group"`
}

type UpdateGroupRequest struct {
	Group Group `json:"group"`
}

type UpdateGroupResponse struct {
	Group Group `json:"group"`
}

type ListGroupsResponse struct {
	Groups []Group                `json:"groups"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}
