package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ShowRequest struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	UserID            uuid.UUID  `json:"user_id" db:"user_id"`
	User              User       `belongs_to:"user"`
	ListingID         uuid.UUID  `json:"listing_id" db:"listing_id"`
	Listing           Listing    `belongs_to:"listing"`
	RequestedShowtime string     `json:"requested_showtime" db:"requested_showtime"`
	ActualShowtime    nulls.Time `json:"actual_showtime" db:"actual_showtime"`
}

// String is not required by pop and may be deleted
func (s ShowRequest) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// ShowRequests is not required by pop and may be deleted
type ShowRequests []ShowRequest

// String is not required by pop and may be deleted
func (s ShowRequests) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *ShowRequest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: s.ListingID, Name: "Listing ID"},
		&validators.UUIDIsPresent{Field: s.UserID, Name: "Requester ID"},
		&validators.StringIsPresent{Field: s.RequestedShowtime, Name: "RequestedShowtime"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *ShowRequest) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *ShowRequest) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
