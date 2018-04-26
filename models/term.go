package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Term struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Dmal      int       `json:"dmal" db:"dmal"`
	Value     float32   `json:"value" db:"value"`
	Listings  Listings  `many_to_many:"listings_terms"`
}

// String is not required by pop and may be deleted
func (t Term) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

func TermTypeOptions() map[string]int {
	return map[string]int{"Daily": 1, "Monthly": 2, "Annually": 3, "A/B Leasing": 4}
}

func (t Term) TypeLabel() string {
	switch t.Dmal {
	case 1:
		return "Daily"
	case 2:
		return "Monthly"
	case 3:
		return "Annually"
	case 4:
		return "A/B Leasing"
	default:
		return "<UNKNOWN>"
	}
}

// Terms is not required by pop and may be deleted
type Terms []Term

// String is not required by pop and may be deleted
func (t Terms) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Term) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.IntIsPresent{Field: t.Dmal, Name: "Dmal"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Term) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Term) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
