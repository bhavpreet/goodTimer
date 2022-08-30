module github.com/bhavpreet/goodTimer/timer

go 1.19

require (
	github.com/bhavpreet/goodTimer/devices/timy2/usb v0.0.0-20220830091232-4d116804d37b
	github.com/bhavpreet/goodTimer/parser v0.0.0-20220830091232-4d116804d37b
)

require (
	github.com/bhavpreet/goodTimer/devices/driver v0.0.0-20220830091232-4d116804d37b // indirect
	github.com/bhavpreet/goodTimer/devices/timy2 v0.0.0-20220830091232-4d116804d37b // indirect
	github.com/bhavpreet/goodTimer/devices/timy2/sim v0.0.0-20220830091232-4d116804d37b // indirect
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3 // indirect
	github.com/google/gousb v1.1.2 // indirect
)

replace (
	github.com/bhavpreet/goodTimer/devices/timy2 => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/timy2
	github.com/bhavpreet/goodTimer/devices/timy2/usb => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/timy2/usb
	github.com/bhavpreet/goodTimer/parser => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/parser
	github.com/bhavpreet/goodTimer/devices/driver => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/driver
	github.com/bhavpreet/goodTimer/devices/timy2/sim => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/timy2/sim
)
