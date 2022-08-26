module github.com/bhavpreet/goodTimer/timer

go 1.19

require (
	github.com/bhavpreet/goodTimer/devices/timy2 v0.0.0-20220826051240-7d2a8ec54a92
	github.com/bhavpreet/goodTimer/timer/serial v0.0.0-20220826051240-7d2a8ec54a92
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
)

require (
	github.com/bhavpreet/goodTimer/devices/timy2/sim v0.0.0-20220826051240-7d2a8ec54a92 // indirect
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
)

replace github.com/bhavpreet/goodTimer/timer/serial => /Users/bhav/go/src/github.com/bhavpreet/goodTimer/timer/serial
