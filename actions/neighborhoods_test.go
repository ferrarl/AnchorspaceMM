package actions

import (
	"fmt"

	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func NewNeighborhood() *models.Neighborhood {
	return &models.Neighborhood{Name: "test neighborhood", Description: "Good stuff"}
}

func (as *ActionSuite) Test_NeighborhoodsResource_List() {
	res := as.HTML("/neighborhoods").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_NeighborhoodsResource_Show() {
	n := NewNeighborhood()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/neighborhoods/" + n.ID.String()).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_NeighborhoodsResource_New() {
	as.Login(true, false)
	res := as.HTML("/neighborhoods/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_NeighborhoodsResource_Create() {
	as.Login(true, false)
	n := NewNeighborhood()
	res := as.HTML("/neighborhoods").Post(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("test neighborhood", n.Name)

	as.Equal(fmt.Sprintf("/neighborhoods/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_NeighborhoodsResource_Edit() {
	as.Login(true, false)
	n := NewNeighborhood()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/neighborhoods/" + n.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_NeighborhoodsResource_Update() {
	as.Login(true, false)
	n := NewNeighborhood()
	as.DB.ValidateAndSave(n)
	n.Name = "The Neighborhood"
	res := as.HTML("/neighborhoods/" + n.ID.String()).Put(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("The Neighborhood", n.Name)

	as.Equal(fmt.Sprintf("/neighborhoods/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_NeighborhoodsResource_Destroy() {
	as.Login(true, false)
	n := NewNeighborhood()
	as.DB.ValidateAndSave(n)
	res := as.HTML("/neighborhoods/" + n.ID.String()).Delete()

	err := as.DB.First(n)
	as.Error(err)
	as.Equal("/neighborhoods", res.Location())
}
