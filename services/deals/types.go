package deals

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Boolish supports API fields that may be returned as either a boolean
// or an integer (0/1). Use Value() to read as a bool with a safe default.
type Boolish struct {
	b *bool
}

func (v *Boolish) UnmarshalJSON(data []byte) error {
	// null -> nil
	if string(data) == "null" {
		v.b = nil
		return nil
	}
	// try boolean
	var bb bool
	if err := json.Unmarshal(data, &bb); err == nil {
		v.b = &bb
		return nil
	}
	// try integer 0/1
	var ii int
	if err := json.Unmarshal(data, &ii); err == nil {
		b := ii != 0
		v.b = &b
		return nil
	}
	return fmt.Errorf("Boolish: unsupported JSON: %s", string(data))
}

// Value returns the boolean value or false when unset.
func (v *Boolish) Value() bool {
	if v != nil && v.b != nil {
		return *v.b
	}
	return false
}

// Intish accepts numbers that may come as a JSON number or a quoted string.
// Value() returns 0 when unset or on parse failure.
type Intish struct {
	n *int
}

func (v *Intish) UnmarshalJSON(data []byte) error {
	// null -> nil
	if string(data) == "null" {
		v.n = nil
		return nil
	}
	// try integer
	var ii int
	if err := json.Unmarshal(data, &ii); err == nil {
		v.n = &ii
		return nil
	}
	// try string then parse
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "" {
			v.n = nil
			return nil
		}
		if parsed, err := strconv.Atoi(s); err == nil {
			v.n = &parsed
			return nil
		}
	}
	return fmt.Errorf("Intish: unsupported JSON: %s", string(data))
}

// Value returns the integer value or 0 if unset.
func (v *Intish) Value() int {
	if v != nil && v.n != nil {
		return *v.n
	}
	return 0
}

// Deal models the deal object returned by the ActiveCampaign V3 API.
// Many fields are strings in the API, even when numerically meaningful.
// Nullable fields are represented as pointers.
type Deal struct {
	Owner               string            `json:"owner,omitempty"`
	Contact             *string           `json:"contact,omitempty"`
	Organization        *string           `json:"organization,omitempty"`
	Group               string            `json:"group,omitempty"`
	Stage               string            `json:"stage,omitempty"`
	Title               string            `json:"title,omitempty"`
	Description         string            `json:"description,omitempty"`
	Percent             string            `json:"percent,omitempty"`
	CDate               string            `json:"cdate,omitempty"`
	MDate               string            `json:"mdate,omitempty"`
	NextDate            *string           `json:"nextdate,omitempty"`
	NextTaskID          *string           `json:"nexttaskid,omitempty"`
	NextTask            *string           `json:"nextTask,omitempty"`
	Value               string            `json:"value,omitempty"`
	Currency            string            `json:"currency,omitempty"`
	WinProbability      *int              `json:"winProbability,omitempty"`
	WinProbabilityMDate string            `json:"winProbabilityMdate,omitempty"`
	Status              string            `json:"status,omitempty"`
	ActivityCount       string            `json:"activitycount,omitempty"`
	NextDealID          string            `json:"nextdealid,omitempty"`
	EDate               string            `json:"edate,omitempty"`
	Links               map[string]string `json:"links,omitempty"`
	ID                  string            `json:"id,omitempty"`
	IsDisabled          *Boolish          `json:"isDisabled,omitempty"`
	Account             *string           `json:"account,omitempty"`
	CustomerAccount     *string           `json:"customerAccount,omitempty"`
	Hash                string            `json:"hash,omitempty"`
}

// DealsListMeta captures pagination or summary metadata.
type DealsListMeta struct {
	Currencies map[string]DealsCurrencyMeta `json:"currencies,omitempty"`
	Total      int                          `json:"total,omitempty"`
}

// DealsCurrencyMeta summarizes totals by currency.
type DealsCurrencyMeta struct {
	Currency string `json:"currency,omitempty"`
	Total    string `json:"total,omitempty"`
	Value    string `json:"value,omitempty"`
}

// ListDealsResponse matches the response shape of GET /api/3/deals.
type ListDealsResponse struct {
	Deals []Deal         `json:"deals,omitempty"`
	Meta  *DealsListMeta `json:"meta,omitempty"`
}

// DealStage models a single deal stage returned by GET /api/3/dealStages.
type DealStage struct {
	CardRegion1 string            `json:"cardRegion1,omitempty"`
	CardRegion2 string            `json:"cardRegion2,omitempty"`
	CardRegion3 string            `json:"cardRegion3,omitempty"`
	CardRegion4 string            `json:"cardRegion4,omitempty"`
	CardRegion5 string            `json:"cardRegion5,omitempty"`
	CDate       string            `json:"cdate,omitempty"`
	Color       string            `json:"color,omitempty"`
	DealOrder   string            `json:"dealOrder,omitempty"`
	Group       string            `json:"group,omitempty"`
	ID          string            `json:"id,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
	Order       string            `json:"order,omitempty"`
	Title       string            `json:"title,omitempty"`
	UDate       string            `json:"udate,omitempty"`
	Width       string            `json:"width,omitempty"`
}

// ListDealStagesResponse wraps the list of deal stages and optional meta.
type ListDealStagesResponse struct {
	DealStages []DealStage `json:"dealStages,omitempty"`
	Meta       struct {
		Total Intish `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}
