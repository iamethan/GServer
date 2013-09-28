// game_svr
package main

import (
	. "GameServer/PlayerSystem"
	"framework"
)

type GameServer struct {
	framework.App

	playerSystem *PlayerSystem
}

func (self *GameServer) Initialize(listenPort int) error {
	// sub systems
	self.playerSystem = new(PlayerSystem)
	self.RegisterSystem(self.playerSystem)

	return self.App.Initialize(listenPort)
}

func (self *GameServer) PlayerSystem() *PlayerSystem {
	return self.playerSystem
}

func NewGameServer() *GameServer {
	server := new(GameServer)
	server.Constructor()

	return server
}
