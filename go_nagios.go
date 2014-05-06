package nagios

import (
	"fmt"
	"os"
)

type NagiosStatusVal int

// The values with which a Nagios check can exit
const (
	NAGIOS_OK NagiosStatusVal = iota
	NAGIOS_WARNING
	NAGIOS_CRITICAL
	NAGIOS_UNKNOWN
)

// Maps the NagiosStatusVal entries to output strings
var (
	valMessages = []string{
		"OK:",
		"WARNING:",
		"CRITICAL:",
		"UNKNOWN:",
	}
)

// A type representing a Nagios check status. The Value is a the exit code
// expected for the check and the Message is the specific output string.
type NagiosStatus struct {
	Message string
	Value   NagiosStatusVal
}

// Take a bunch of NagiosStatus pointers and find the highest value, then
// combine all the messages. Things win in the order of highest to lowest.
func (status *NagiosStatus) Aggregate(otherStatuses []*NagiosStatus) {
	for _, s := range otherStatuses {
		if status.Value < s.Value {
			status.Value = s.Value
		}

		status.Message += " - " + s.Message
	}
}

// Exit with an UNKNOWN status and appropriate message
func Unknown(output string) {
	ExitWithStatus(&NagiosStatus{output, NAGIOS_UNKNOWN})
}

// Exit with an CRITICAL status and appropriate message
func Critical(err error) {
	ExitWithStatus(&NagiosStatus{err.Error(), NAGIOS_CRITICAL})
}

// Exit with an WARNING status and appropriate message
func Warning(output string) {
	ExitWithStatus(&NagiosStatus{output, NAGIOS_WARNING})
}

// Exit with an OK status and appropriate message
func Ok(output string) {
	ExitWithStatus(&NagiosStatus{output, NAGIOS_OK})
}

// Exit with a particular NagiosStatus
func ExitWithStatus(status *NagiosStatus) {
	fmt.Fprintln(os.Stdout, valMessages[status.Value], status.Message)
	os.Exit(int(status.Value))
}
