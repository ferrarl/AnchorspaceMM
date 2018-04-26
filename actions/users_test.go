package actions

import (
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func NewUser() *models.User {
	return &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		LastName:             "Bates",
		FirstName:            "Mark",
		Phone:                "12344456789",
	}
}

func (as *ActionSuite) Test_Users_New() {
	res := as.HTML("/users/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := NewUser()

	res := as.HTML("/users").Post(u)
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
