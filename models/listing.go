package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Listing struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
	Name             string        `json:"name" db:"name"`
	NeighborhoodID   uuid.UUID     `json:"neighborhood_id" db:"neighborhood_id"`
	Neighborhood     Neighborhood  `belongs_to:"neighborhood"`
	Address1         string        `json:"address1" db:"address1"`
	Address2         string        `json:"address2" db:"address2"`
	City             string        `json:"city" db:"city"`
	State            string        `json:"state" db:"state"`
	Zipcode          string        `json:"zipcode" db:"zipcode"`
	PictureUrl       string        `json:"picture_url" db:"picture_url"`
	VideoUrl         string        `json:"video_url" db:"video_url"`
	ThreeDUrl        string        `json:"3d_url" db:"threed_url"`
	Size             int           `json:"size" db:"size"`
	ListerID         uuid.UUID     `json:"lister_id" db:"user_id"`
	Lister           User          `belongs_to:"user"`
	DailyPricing     float64       `json:"daily_pricing" db:"daily_pricing"`
	MonthlyPricing   float64       `json:"monthly_pricing" db:"monthly_pricing"`
	AnnuallyPricing  float64       `json:"annually_pricing" db:"annually_pricing"`
	AbleasingPricing float64       `json:"ableasing_pricing" db:"ableasing_pricing"`
	CamRate          float64       `json:"cam_rate" db:"cam_rate"`
	Attributes       slices.String `json:"attributes" db:"attributes"`
	Population1      int           `json:"population_1" db:"population_1"`
	Population3      int           `json:"population_3" db:"population_3"`
	Population5      int           `json:"population_5" db:"population_5"`
	MedianIncome1    int           `json:"median_income_1" db:"median_income_1"`
	MedianIncome3    int           `json:"median_income_3" db:"median_income_3"`
	MedianIncome5    int           `json:"median_income_5" db:"median_income_5"`
	WalkScore        int           `json:"walk_score" db:"walk_score"`
	TrafficScore     int           `json:"traffic_score" db:"traffic_score"`
	Uses             Uses          `many_to_many:"listings_uses"`
	Terms            Terms         `many_to_many:"listings_terms"`
	ShowRequests     ShowRequests  `has_many:"show_requests"`
}

//Uses and Terms required as well

// String is not required by pop and may be deleted
func (l Listing) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Listings is not required by pop and may be deleted
type Listings []Listing

// String is not required by pop and may be deleted
func (l Listings) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Listing) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.Name, Name: "Name"},
		&validators.StringIsPresent{Field: l.Address1, Name: "Address1"},
		&validators.StringIsPresent{Field: l.City, Name: "City"},
		&validators.StringIsPresent{Field: l.State, Name: "State"},
		&validators.StringIsPresent{Field: l.Zipcode, Name: "Zipcode"},
		&validators.StringIsPresent{Field: l.PictureUrl, Name: "PictureUrl"},
		&validators.IntIsPresent{Field: l.Size, Name: "Size"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Listing) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Listing) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
