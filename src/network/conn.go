// conn
package network

import (
	//"io"
	"log"
	"net"
)

type conn struct {
	rwc    net.Conn
	maxCap int
	msgs   chan interface{}
}

func (self *conn) maxMsgsCap(maxCap int) {
	self.maxCap = maxCap
}

func (self *conn) server() {
	defer self.close()

	if self.maxCap == 0 {
		self.maxCap = 256
	}

	self.msgs = make(chan interface{}, self.maxCap)

	for {
		buf := make([]byte, 1024)
		len, e := self.rwc.Read(buf)
		if e != nil {
			log.Println("read error:", e)
			break
		}

		str := string(buf[:len])

		log.Printf("recv from:%s len:%d, content:%s", self.rwc.RemoteAddr().String(), len, str)

		self.rwc.Write(buf[:len])
	}
}

func (self *conn) close() {
	self.Flush()

	if cap(self.msgs) > 0 {
		close(self.msgs)
	}

	if self.rwc != nil {
		self.rwc.Close()
		self.rwc = nil
	}

	log.Println("the connection has been closed!")
}

func (self *conn) Send(msg []byte) {

}

func (self *conn) Flush() {

}
