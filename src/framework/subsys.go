// subsys
package framework

type ISubSystem interface {
	Name() string
	Initialize() error
	Finalize() error
	HeartBeat(timepass float32)
}

type SubSystem struct {
}

func (self *SubSystem) Name() string {
	return ""
}

func (self *SubSystem) Initialize() error {
	return nil
}

func (self *SubSystem) Finalize() error {
	return nil
}

func (self *SubSystem) HeartBeat(timepass float32) {

}
