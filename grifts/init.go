package grifts

import (
  "github.com/gobuffalo/buffalo"
	"test_project/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
