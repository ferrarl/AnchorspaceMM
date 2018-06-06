package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

func NewShowRequest(user, listing uuid.UUID) *models.ShowRequest {
	return &models.ShowRequest{RequestedShowtime: "test show", UserID: user, ListingID: listing}
}

func (as *ActionSuite) Test_ShowRequestsResource_List() {
	as.Login(false, true)
	res := as.HTML("/show_requests").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ShowRequestsResource_Show() {
	l := NewListing(as)
	as.DB.ValidateAndSave(l)
	u := as.Login(false, false)
	n := NewShowRequest(u.ID, l.ID)
	as.DB.ValidateAndSave(n)
	res := as.HTML("/show_requests/" + n.ID.String()).Get()
	as.Equal(301, res.Code)
}

func (as *ActionSuite) Test_ShowRequestsResource_New() {
	as.Login(false, false)
	res := as.HTML("/show_requests/new").Get()
	as.Equal(301, res.Code)
}

func (as *ActionSuite) Test_ShowRequestsResource_Create() {
	l := NewListing(as)
	as.DB.ValidateAndSave(l)
	u := as.Login(false, false)
	n := NewShowRequest(u.ID, l.ID)
	res := as.JSON("/show_requests").Post(n)

	as.Equal(201, res.Code)

	err := as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	as.Equal("test show", n.RequestedShowtime)
}

func (as *ActionSuite) Test_ShowRequestsResource_Edit() {
	l := NewListing(as)
	u := as.Login(false, true)
	l.ListerID = u.ID
	as.DB.ValidateAndSave(l)
	n := NewShowRequest(u.ID, l.ID)
	as.DB.ValidateAndSave(n)
	res := as.HTML("/show_requests/" + n.ID.String() + "/edit").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ShowRequestsResource_Update() {
	l := NewListing(as)
	u := as.Login(false, true)
	l.ListerID = u.ID
	as.DB.ValidateAndSave(l)
	na := NewShowRequest(u.ID, l.ID)
	as.DB.ValidateAndSave(na)

	nb := &models.ShowRequest{}
	err := as.DB.Eager().First(nb)
	as.NoError(err)

	now := time.Now().UTC()
	nb.ActualShowtime = nulls.NewTime(now)

	res := as.HTML("/show_requests/" + nb.ID.String()).Put(nb)

	as.Equal(302, res.Code)

	n := &models.ShowRequest{}
	err = as.DB.First(n)
	as.NoError(err)
	as.NotZero(n.ID)

	//check that we saved the time properly
	diff := now.Sub(n.ActualShowtime.Time).Round(time.Second).String()
	as.Equal("0s", diff)

	as.Equal(fmt.Sprintf("/show_requests/%s", n.ID), res.Location())
}

func (as *ActionSuite) Test_ShowRequestsResource_Destroy() {
	l := NewListing(as)
	as.DB.ValidateAndSave(l)
	u := as.Login(true, false)
	n := NewShowRequest(u.ID, l.ID)
	as.DB.ValidateAndSave(n)
	res := as.HTML("/show_requests/" + n.ID.String()).Delete()

	err := as.DB.First(n)
	as.Error(err)
	as.Equal("/show_requests", res.Location())
}
