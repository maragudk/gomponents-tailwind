package tailwind

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/attr"
	"github.com/maragudk/gomponents/el"
)

// Container restricts the width of the children, centers, and adds some horizontal padding.
func Container(children ...g.Node) g.NodeFunc {
	return el.Div(attr.Class("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"), g.Group(children))
}
