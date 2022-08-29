module github.com/bhavpreet/goodTimer/timer

go 1.19

require (
	github.com/bhavpreet/goodTimer/devices/timy2 v0.0.0-20220829143431-03f1009b633b
	github.com/bhavpreet/goodTimer/devices/timy2/usb v0.0.0-20220829143431-03f1009b633b
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
)

require (
	github.com/bhavpreet/goodTimer/devices/driver v0.0.0-20220829115307-8a6306606ea1 // indirect
	github.com/bhavpreet/goodTimer/devices/timy2/sim v0.0.0-20220829115307-8a6306606ea1 // indirect
	github.com/google/gousb v1.1.2 // indirect
)

replace github.com/bhavpreet/goodTimer/devices/timy2/usb => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/timy2/usb

replace github.com/bhavpreet/goodTimer/devices/timy2 => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/devices/timy2
