package node

import (
	"regexp"
	"time"
)

type Node struct {
	Name       string
	Timestamps []time.Time
	Schedule   *time.Duration
}

// Returns duration since last checkin showing only largest time increment
func (n *Node) Since() string {
	last := n.Timestamps[len(n.Timestamps)-1]
	s := time.Since(last).String()
	r, err := regexp.Compile("[0-9]*(\\.[0-9]*)?[a-z]")
	if err != nil {
		return "-1"
	}
	return r.FindString(s)
}

// Returns if last checked in time is less than scheduled checkin interval.
func (n *Node) Health() string {
	// Schedule not set
	if n.Schedule == nil {
		return "?"
	}
	last := n.Timestamps[len(n.Timestamps)-1]
	// Checkin is over scheduled time.
	if time.Since(last) > *n.Schedule {
		return "Bad"
	}
	return "Good"
}
