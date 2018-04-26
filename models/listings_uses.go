package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type ListingsUse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ListingID uuid.UUID `json:"listing_id" db:"listing_id"`
	Listing   Listing   `belongs_to:"listing"`
	UseID     uuid.UUID `json:"use_id" db:"use_id"`
	Use       Use       `belongs_to:"use"`
}

// String is not required by pop and may be deleted
func (l ListingsUse) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// ListingsUses is not required by pop and may be deleted
type ListingsUses []ListingsUse

// String is not required by pop and may be deleted
func (l ListingsUses) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *ListingsUse) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *ListingsUse) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *ListingsUse) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
