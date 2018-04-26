package actions

import (
	"fmt"

	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func (as *ActionSuite) Test_UsesResource_List() {
	as.Login(false, false)
	res := as.HTML("/uses").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_UsesResource_Show() {
	as.Login(false, false)
	n := &models.Use{Name: "test use", Description: "Good stuff"}
	as.DB.ValidateAndSave(n)
	res := as.HTML("/uses/" + n.ID.String()).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_UsesResource_New() {
	as.Login(true, false)
	res := as.HTML("/uses/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_UsesResource_Create() {
	as.Login(true, false)
	n := &models.Use{Name: "test use", Description: "Good stuff"}
	res := as.HTML("/uses").Post(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("test use", n.Name)

	as.Equal(fmt.Sprintf("/uses/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_UsesResource_Edit() {
	as.Login(true, false)
	n := &models.Use{Name: "test use", Description: "Good stuff"}
	as.DB.ValidateAndSave(n)
	res := as.HTML("/uses/" + n.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_UsesResource_Update() {
	as.Login(true, false)
	n := &models.Use{Name: "test use", Description: "Good stuff"}
	as.DB.ValidateAndSave(n)
	n.Name = "The Use"
	res := as.HTML("/uses/" + n.ID.String()).Put(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("The Use", n.Name)

	as.Equal(fmt.Sprintf("/uses/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_UsesResource_Destroy() {
	as.Login(true, false)
	n := &models.Use{Name: "test use", Description: "Good stuff"}
	as.DB.ValidateAndSave(n)
	res := as.HTML("/uses/" + n.ID.String()).Delete()

	err := as.DB.First(n)
	as.Error(err)
	as.Equal("/uses", res.Location())
}
