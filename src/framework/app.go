// app
package framework

import (
	"log"
	"network"
	"time"
)

const (
	FPS int = 50
)

type App struct {
	subSystems []ISubSystem
	isQuit     bool
}

func (self *App) Constructor() {
	self.isQuit = false
	self.subSystems = make([]ISubSystem, 0, 10)
}

func (self *App) Destructor() {

}

func (self *App) Initialize(listenPort int) error {
	for _, sys := range self.subSystems {
		if sys.Initialize() != nil {
			log.Printf("sys<%s> initialize failed!", sys.Name())
		}
	}

	go network.ListenAndServer(listenPort)

	return nil
}

func (self *App) Finalize() error {
	for _, sys := range self.subSystems {
		if sys.Finalize() != nil {
			log.Println("sys<%s> finialize failed!", sys.Name())
		}
	}
	return nil
}

func (self *App) Run() {
	if FPS == 0 {
		log.Fatalln("can not start the server because of FPS is zero")
		return
	}

	framDuration := time.Second / time.Duration(FPS)

	now := time.Now()
	var timeToSleep time.Duration = 0

	for self.isQuit == false {
		timepass := time.Now().Sub(now)

		timeToSleep = framDuration - timepass
		if timeToSleep > 0 {
			time.Sleep(timeToSleep)
		} else {
			timeToSleep = 0
		}

		self.handleMsgs()

		self.HeartBeat(float32(timepass+timeToSleep) / float32(time.Second))

		now = time.Now()
	}
}

func (self *App) Quit() {
	self.isQuit = true
}

func (self *App) HeartBeat(timepass float32) {
	for _, sys := range self.subSystems {
		sys.HeartBeat(timepass)
	}
}

func (self *App) RegisterSystem(sys ISubSystem) {
	name := sys.Name()

	for _, v := range self.subSystems {
		if v.Name() == name {
			log.Fatalf("the subsystem<%s> hs been registed\n", name)
			return
		}
	}

	self.subSystems = append(self.subSystems, sys)
}

func (self *App) handleMsgs() {

}
