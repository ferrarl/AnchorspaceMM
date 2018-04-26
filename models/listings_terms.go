package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type ListingsTerm struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ListingID uuid.UUID `json:"listing_id" db:"listing_id"`
	Listing   Listing   `belongs_to:"listing"`
	TermID    uuid.UUID `json:"term_id" db:"term_id"`
	Term      Term      `belongs_to:"term"`
}

// String is not required by pop and may be deleted
func (l ListingsTerm) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// ListingsTerms is not required by pop and may be deleted
type ListingsTerms []ListingsTerm

// String is not required by pop and may be deleted
func (l ListingsTerms) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *ListingsTerm) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *ListingsTerm) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *ListingsTerm) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
