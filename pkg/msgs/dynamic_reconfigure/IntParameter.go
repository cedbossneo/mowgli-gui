//autogenerated:yes
//nolint:revive,lll
package dynamic_reconfigure

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type IntParameter struct {
	msg.Package `ros:"dynamic_reconfigure"`
	Name        string
	Value       int32
}
