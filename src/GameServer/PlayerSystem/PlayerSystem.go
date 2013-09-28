// PlayerSystem
package PlayerSystem

import (
	"framework"
	//"log"
)

type PlayerSystem struct {
	framework.SubSystem
}

func (self *PlayerSystem) Name() string {
	return "PlayerSystem"
}

func (self *PlayerSystem) HeartBeat(timepass float32) {
	//log.Println("player system heart beat:", timepass)
}
