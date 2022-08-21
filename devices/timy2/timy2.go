package timy2

const (
	RunningTimeFormat      = "15:04:05.0"
	ImpulseTimeFormat      = "15:04:05.0000"
	BaudRate               = 38400
	ImpulseLength          = len("B####bCxxbHH:MM:SS:zhtq")
	ImpulseTimeStampLength = len("HH:MM:SS:zhtq")
	B                      = " "
)

const (
	// All Impulses
	START_IMPULSE = iota
	END_IMPULSE
	OTHER_IMPULSE
)

// TimeFormats
// Channel 0 C0    -    Precision 1/10.000
// Channel 0M C0M  -    Precision 1/100 – manual = keypad
// Channel 1 C1    -    Precision 1/10.000
// Channel 1M C1M  -    Precision 1/100 – manual = keypad
// Channel 2 C2    -    Precision 1/10.000
// Channel 3 C3    -    Precision 1/10.000
// Channel 4 C4    -    Precision 1/10.000
// Channel 5 C5    -    Precision 1/100
// Channel 6 C6    -    Precision 1/100
// Channel 7 C7    -    Precision 1/100
// Channel 8 C8    -    Precision 1/100
// Documentation says something, output example says something else.
// Adjusting below for the output example
var TimeFormatsForChannels map[string]string = map[string]string{
	"C0":  "15:04:05.0000",
	"C0M": "15:04:05.00",
	"C1":  "15:04:05.0000",
	"C1M": "15:04:05.00",
	"C2":  "15:04:05.0000",
	"C3":  "15:04:05.0000",
	"C4":  "15:04:05.0000",
	"C5":  "15:04:05.0000",
	"C6":  "15:04:05.00",
	"C7":  "15:04:05.00",
	"C8":  "15:04:05.00",
}

// Type of Channel
var ChannelType = map[string]int{
	"C0":  START_IMPULSE,
	"C0M": START_IMPULSE,
	"C1":  END_IMPULSE,
	"C1M": END_IMPULSE,
	"C2":  OTHER_IMPULSE,
	"C3":  OTHER_IMPULSE,
	"C4":  OTHER_IMPULSE,
	"C5":  OTHER_IMPULSE,
	"C6":  OTHER_IMPULSE,
	"C7":  OTHER_IMPULSE,
	"C8":  OTHER_IMPULSE,
}
