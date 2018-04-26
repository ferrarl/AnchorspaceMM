package actions

import (
	"fmt"

	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func NewTerm() *models.Term {
	return &models.Term{Name: "test term", Dmal: 1, Value: 1.0}
}

func (as *ActionSuite) Test_TermsResource_List() {
	as.Login(false, false)
	res := as.HTML("/terms").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TermsResource_Show() {
	as.Login(false, false)
	n := NewTerm()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/terms/" + n.ID.String()).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TermsResource_New() {
	as.Login(true, false)
	res := as.HTML("/terms/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TermsResource_Create() {
	as.Login(true, false)
	n := NewTerm()
	res := as.HTML("/terms").Post(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("test term", n.Name)

	as.Equal(fmt.Sprintf("/terms/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_TermsResource_Edit() {
	as.Login(true, false)
	n := NewTerm()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/terms/" + n.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_TermsResource_Update() {
	as.Login(true, false)
	n := NewTerm()
	as.DB.ValidateAndSave(n)
	n.Name = "The Term"
	res := as.HTML("/terms/" + n.ID.String()).Put(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("The Term", n.Name)

	as.Equal(fmt.Sprintf("/terms/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_TermsResource_Destroy() {
	as.Login(true, false)
	n := NewTerm()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/terms/" + n.ID.String()).Delete()

	err := as.DB.First(n)
	as.Error(err)
	as.Equal("/terms", res.Location())
}
