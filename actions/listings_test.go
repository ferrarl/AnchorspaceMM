package actions

import (
	"fmt"

	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func NewListing(as *ActionSuite) *models.Listing {
	n := NewNeighborhood()
	a := NewUser()
	as.DB.ValidateAndSave(n)
	as.DB.ValidateAndSave(a)
	return &models.Listing{Name: "test listing", NeighborhoodID: n.ID, Address1: "555 Rough Road", City: "Bham", State: "AL", Zipcode: "35216", PictureUrl: "https://google.com", Size: 42, ListerID: a.ID, Population1: 1, Population3: 1, Population5: 1, DailyPricing: 1.0, MonthlyPricing: 1.0, AnnuallyPricing: 1.0, AbleasingPricing: 1.0, CamRate: 1.0, MedianIncome1: 1, MedianIncome3: 1, MedianIncome5: 1, WalkScore: 1, TrafficScore: 1}
}

func (as *ActionSuite) Test_ListingsResource_List() {
	res := as.HTML("/listings").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ListingsResource_Show() {
	n := NewListing(as)
	as.DB.ValidateAndSave(n)
	res := as.HTML("/listings/" + n.ID.String()).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ListingsResource_New() {
	as.Login(false, true)
	res := as.HTML("/listings/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ListingsResource_Create() {
	u := as.Login(false, true)
	n := NewListing(as)
	n.ListerID = u.ID
	res := as.HTML("/listings").Post(n)
	//	fmt.Println("============================")
	//	fmt.Printf("%#v\n", res.Result())
	//	fmt.Println("============================")

	err := as.DB.First(n)

	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("test listing", n.Name)

	as.Equal(fmt.Sprintf("/listings/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_ListingsResource_Edit() {
	u := as.Login(false, true)
	n := NewListing(as)
	n.ListerID = u.ID
	as.DB.ValidateAndSave(n)
	res := as.HTML("/listings/" + n.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ListingsResource_Update() {
	u := as.Login(false, true)
	n := NewListing(as)
	n.ListerID = u.ID
	as.DB.ValidateAndCreate(n)
	n.Name = "The Listing"
	res := as.HTML("/listings/" + n.ID.String()).Put(n)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("The Listing", n.Name)

	as.Equal(fmt.Sprintf("/listings/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_ListingsResource_Destroy() {
	u := as.Login(true, false)
	n := NewListing(as)
	n.ListerID = u.ID
	as.DB.ValidateAndSave(n)
	res := as.HTML("/listings/" + n.ID.String()).Delete()

	err := as.DB.Find(&n, n.ID)
	as.Error(err)
	as.Equal("/listings", res.Location())
}
