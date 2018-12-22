package grifts

import (
	"github.com/dvidr/blog/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
