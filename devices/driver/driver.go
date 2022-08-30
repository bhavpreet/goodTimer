package driver

type Reader interface {
	Initialize(interface{}) error
	SubscribeToImpulses() (chan string, func(), error) // func = close() => defer close() 
}
