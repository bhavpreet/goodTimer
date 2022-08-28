package driver

type Reader interface {
	Initialize(interface{}) error
	SubscribeToImpulses(done chan bool) (chan string, error)
}
