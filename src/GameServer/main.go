// main
package main

func main() {

	svr := NewGameServer()

	svr.Initialize(1000)

	svr.Run()

	svr.Finalize()
}
