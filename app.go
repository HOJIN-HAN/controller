package controller

import "github.com/revel/revel"

// App ..
type App struct {
	GorpController
}

// Index ...
func (c App) Index() revel.Result {
	return c.Render()
}
