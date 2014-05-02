package go_nagios

import (
	"fmt"
	"os"
)

type NagiosStatusVal int

const (
	NAGIOS_OK NagiosStatusVal = iota
	NAGIOS_WARNING
	NAGIOS_CRITICAL
	NAGIOS_UNKNOWN
)

type NagiosStatus struct {
	Message string
	Value NagiosStatusVal
}

func (status *NagiosStatus) Aggregate(otherStatuses []*NagiosStatus) {
	for _, s := range(otherStatuses) {
		if status.Value < s.Value {
			status.Value = s.Value
		}

		status.Message += " - " + s.Message
	}
}

func Unknown(output string) {
	fmt.Fprint(os.Stdout, "UNKNOWN:", output)
	os.Exit(3)
}

func Critical(err error) {
	fmt.Fprint(os.Stdout, "CRITICAL:", err.Error())
	os.Exit(2)
}

func Warning(output string) {
	fmt.Fprint(os.Stdout, "WARNING:", output)
	os.Exit(1)
}

func Ok(output string) {
	fmt.Fprint(os.Stdout, "OK:", output)
	os.Exit(0)
}

func ExitWithNagiosStatus(status *NagiosStatus) {
	switch {
		case status.Value == NAGIOS_UNKNOWN:
			println("UNKNOWN:", status.Message)
			os.Exit(3)
		case status.Value == NAGIOS_CRITICAL:
			println("CRITICAL:", status.Message)
			os.Exit(2)
		case status.Value == NAGIOS_WARNING:
			println("WARNING:", status.Message)
			os.Exit(1)
		case status.Value == NAGIOS_OK:
			println("OK:", status.Message)
			os.Exit(0)
	}
}
