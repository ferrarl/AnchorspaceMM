package grifts

import (
	"github.com/gobuffalo/buffalo"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
