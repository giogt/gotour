package stringers

import "fmt"

// The IPAddr type describes an IP address
type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}
