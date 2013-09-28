// service
package network

import (
	"log"
	"net"
	"strconv"
)

func ListenAndServer(port int) error {
	log.Printf("start listen port<%d>...\n", port)
	svr := new(Server)
	return svr.RunAndListen(port)
}

// Server accept incoming client connections on the listener l,
// create a new conn goroutine for each. The service goroutines
// read the msg, save all int buff and wait for app to process.
type Server struct {
}

func (self *Server) RunAndListen(port int) error {
	addr := ":" + strconv.Itoa(port)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		return e
	}

	return self.Serve(l)
}

func (self *Server) Serve(l net.Listener) error {
	defer l.Close()

	for {
		rw, e := l.Accept()
		if e != nil {
			log.Printf("network service error:%v\n", e)
			continue
		}

		c, err := self.newConn(rw)
		if err != nil {
			continue
		}

		go c.server()
	}
}

func (self *Server) newConn(rwc net.Conn) (c *conn, err error) {
	c = new(conn)
	c.rwc = rwc

	return c, nil
}
