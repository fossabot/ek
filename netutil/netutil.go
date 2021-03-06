// Package netutil with network utils
package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"net"
	"os"
	"strings"

	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with routes info in procfs
var procRouteFile = "/proc/net/route"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetIP return main IPv4 address
func GetIP() string {
	return getMainIP(false)
}

// GetIP6 return main IPv6 address
func GetIP6() string {
	return getMainIP(true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMainIP(v6 bool) string {
	interfaces, err := net.Interfaces()

	if err != nil {
		return ""
	}

	defaultInterface := getDefaultRouteInterface()

	for i := len(interfaces) - 1; i >= 0; i-- {
		// Ignore TUN/TAP interfaces
		if strings.HasPrefix(interfaces[i].Name, "t") {
			continue
		}

		if defaultInterface != "" && interfaces[i].Name != defaultInterface {
			continue
		}

		addr, err := interfaces[i].Addrs()

		if err != nil || len(addr) == 0 {
			continue
		}

		for _, a := range addr {
			ipnet, ok := a.(*net.IPNet)

			if ok && strings.Contains(ipnet.IP.String(), "::") == v6 {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func getDefaultRouteInterface() string {
	fd, err := os.OpenFile(procRouteFile, os.O_RDONLY, 0)

	if err != nil {
		return ""
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	var header bool

	for s.Scan() {
		if !header {
			header = true
			continue
		}

		if strutil.ReadField(s.Text(), 1, true) == "00000000" {
			return strutil.ReadField(s.Text(), 0, true)
		}
	}

	return ""
}
