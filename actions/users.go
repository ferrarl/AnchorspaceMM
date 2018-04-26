package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

type UsersResource struct {
	buffalo.Resource
}

func (v UsersResource) New(c buffalo.Context) error {
	u := models.User{}
	return c.Render(200, r.Auto(c, u))
}

// User Create registers a new user with the application.
func (v UsersResource) Create(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("users/new.html"))
	}

	if uid := c.Session().Get("current_user_id"); uid == nil {
		c.Session().Set("current_user_id", u.ID)
		c.Flash().Add("success", "Welcome to Buffalo!")
		return c.Redirect(302, "/")
	} else {
		return c.Render(200, r.Auto(c, u))
	}

}

// List gets all Users. This function is mapped to the path
// GET /users
func (v UsersResource) List(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin {
		c.Flash().Add("danger", "Could not access that path")
		return c.Redirect(302, "/")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	users := &models.Users{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Users from the DB
	if err := q.All(users); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, users))
}

// Show gets the data for one User. This function is mapped to
// the path GET /users/{user_id}
func (v UsersResource) Show(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin && current_user.ID.String() != c.Param("user_id") {
		c.Flash().Add("danger", "Could not access that path")
		return c.Redirect(302, "/")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, user))
}

// Edit renders a edit form for a User. This function is
// mapped to the path GET /users/{user_id}/edit
func (v UsersResource) Edit(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin && current_user.ID.String() != c.Param("user_id") {
		c.Flash().Add("danger", "Could not access that path")
		return c.Redirect(302, "/")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, user))
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (v UsersResource) Update(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin && current_user.ID.String() != c.Param("user_id") {
		c.Flash().Add("danger", "Could not access that path")
		return c.Redirect(302, "/")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind User to the html form elements
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := user.Update(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, user))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User was updated successfully")

	// and redirect to the users index page
	return c.Render(200, r.Auto(c, user))
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func (v UsersResource) Destroy(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin {
		c.Flash().Add("danger", "Could not access that path")
		return c.Redirect(302, "/")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return errors.WithStack(err)
	}

	if user.ID == current_user.ID {
		c.Session().Clear()
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "User was destroyed successfully")

	// Redirect to the users index page
	return c.Render(200, r.Auto(c, user))
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
