package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   uuid.UUID    `json:"id" db:"id"`
	CreatedAt            time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at" db:"updated_at"`
	Email                string       `json:"email" db:"email"`
	PasswordHash         string       `json:"-" db:"password_hash"`
	Password             string       `json:"-" db:"-"`
	PasswordConfirmation string       `json:"-" db:"-"`
	FirstName            string       `json:"first_name" db:"first_name"`
	LastName             string       `json:"last_name" db:"last_name"`
	Phone                string       `json:"phone" db:"phone"`
	Company              nulls.String `json:"company,omitempty" db:"company"`
	IsAgent              bool         `json:"is_agent" db:"is_agent"`
	IsAdmin              bool         `json:"is_admin" db:"is_admin"`
	SecondaryPhone       nulls.String `json:"secondary_phone,omitempty" db:"secondary_phone"`
	PublicEmail          nulls.String `json:"public_email,omitempty" db:"public_email"`
	Address1             nulls.String `json:"address1,omitempty" db:"address1"`
	Address2             nulls.String `json:"address2,omitempty" db:"address2"`
	State                nulls.String `json:"state,omitempty" db:"state"`
	Zipcode              nulls.String `json:"zipcode,omitempty" db:"zipcode"`
	City                 nulls.String `json:"city,omitempty" db:"city"`
	Listings             Listings     `has_many:"listings"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) PrepFields() (*validate.Errors, error) {
	if len(strings.TrimSpace(u.Email)) > 0 {
		u.Email = strings.ToLower(u.Email)
	}
	if len(strings.TrimSpace(u.Password)) > 0 {
		ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return validate.NewErrors(), errors.WithStack(err)
		}
		u.PasswordHash = string(ph)
	}
	return validate.NewErrors(), nil
}

// Update wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Update(tx *pop.Connection) (*validate.Errors, error) {
	if verrs, err := u.PrepFields(); (verrs != nil && verrs.HasAny()) || err != nil {
		return verrs, err
	}
	return tx.ValidateAndSave(u)
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	if verrs, err := u.PrepFields(); verrs.HasAny() || err != nil {
		return verrs, err
	}
	return tx.ValidateAndCreate(u)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate" method.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	//validate based on is agent
	if !u.IsAgent { //is not an agent
		return validate.Validate(
			&validators.StringIsPresent{Field: u.Email, Name: "Email"},
			&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
			&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
			&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
			&validators.StringIsPresent{Field: u.Phone, Name: "Phone"},
			// check to see if the email address is already taken:
			&validators.FuncValidator{
				Field:   u.Email,
				Name:    "Email",
				Message: "%s is already taken",
				Fn: func() bool {
					var b bool
					q := tx.Where("email = ?", u.Email)
					if u.ID != uuid.Nil {
						q = q.Where("id != ?", u.ID)
					}
					b, err = q.Exists(u)
					if err != nil {
						return false
					}
					return !b
				},
			},
		), err

	} else { // trying to save agent
		return validate.Validate(
			&validators.StringIsPresent{Field: u.Email, Name: "Email"},
			&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
			&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
			&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
			&validators.StringIsPresent{Field: u.Phone, Name: "Phone"},
			&validators.StringIsPresent{Field: u.PublicEmail.String, Name: "PublicEmail"},
			&validators.StringIsPresent{Field: u.Company.String, Name: "Company"},
			&validators.StringIsPresent{Field: u.Address1.String, Name: "Address1"},
			&validators.StringIsPresent{Field: u.City.String, Name: "City"},
			&validators.StringIsPresent{Field: u.State.String, Name: "State"},
			&validators.StringIsPresent{Field: u.Zipcode.String, Name: "Zipcode"},
			// check to see if the email address is already taken:
			&validators.FuncValidator{
				Field:   u.Email,
				Name:    "Email",
				Message: "%s is already taken",
				Fn: func() bool {
					var b bool
					q := tx.Where("email = ?", u.Email)
					if u.ID != uuid.Nil {
						q = q.Where("id != ?", u.ID)
					}
					b, err = q.Exists(u)
					if err != nil {
						return false
					}
					return !b
				},
			},
		), err

	}
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
	), err
}
