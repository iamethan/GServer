// msg
package network

type Msg struct {
	ID  int16
	Len int16
}

func (self *Msg) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (self *Msg) Write(p []byte) (n int, err error) {
	return 0, nil
}
