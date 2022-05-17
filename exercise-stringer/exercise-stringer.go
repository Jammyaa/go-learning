/*
Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	host := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"GoogleDNS": {8, 8, 8, 8},
	}

	fmt.Println(host)

	for name, ip := range host {
		fmt.Println(name, ip)
	}
}
