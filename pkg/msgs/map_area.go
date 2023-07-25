//autogenerated:yes
//nolint:revive,lll
package msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
)

type MapArea struct {
	msg.Package `ros:"xbot_msgs"`
	Name        string
	Area        geometry_msgs.Polygon
	Obstacles   []geometry_msgs.Polygon
}
