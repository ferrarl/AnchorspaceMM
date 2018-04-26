package actions

import (
	"log"

	"github.com/gobuffalo/pop/nulls"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func (as *ActionSuite) Login(is_admin, is_agent bool) *models.User {
	u := NewUser()
	u.IsAdmin = is_admin
	u.IsAgent = is_agent

	if is_agent {
		u.PublicEmail = nulls.NewString("test@email.com")
		u.City = nulls.NewString("Bham")
		u.State = nulls.NewString("AL")
		u.Zipcode = nulls.NewString("35216")
		u.Company = nulls.NewString("Nerf")
		u.Address1 = nulls.NewString("555 Rough Road")
	}

	verrs, err := u.Create(as.DB)
	//	fmt.Printf("============================\n%v\n%v\n=================================\n", err, verrs)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
	return u
}

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/signin").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		LastName:             "Bates",
		FirstName:            "Mark",
		Phone:                "12344456789",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	u := &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		LastName:             "Bates",
		FirstName:            "Mark",
		Phone:                "12344456789",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	log.Printf("========================\n%v\n", verrs)
	as.False(verrs.HasAny())

	u.Password = "bad"
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
