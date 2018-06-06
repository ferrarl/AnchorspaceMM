package actions

import (
	"github.com/gobuffalo/buffalo"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

// AdminPanel default implementation.
func AdminPanel(c buffalo.Context) error {
	current_user := c.Value("current_user").(*models.User)
	if !current_user.IsAdmin && !current_user.IsAgent {
		c.Flash().Add("danger", "You don't have access to that!")
		return c.Render(404, r.HTML("index.html"))
	}
	return c.Render(200, r.HTML("admin/panel.html"))
}
