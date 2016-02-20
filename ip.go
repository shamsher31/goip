// Tiny utility to deal with IP
package ip

import (
	"errors"
	"net"
	"net/http"
)

// Returns local IP address
func Local() (net.IP, error) {

	var localIp net.IP

	address, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	for _, addrs := range address {

		if ipnet, ok := addrs.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			if ipnet.IP.To4() != nil {
				localIp = ipnet.IP
			}

		}
	}

	return localIp, nil

}

// Returns remote IP address of the requested client
func Remote(req *http.Request) (net.IP, error) {

	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		return nil, err
	}

	remoteIP := net.ParseIP(ip)

	if remoteIP == nil {
		return nil, errors.New("Not a remote IP")
	}

	return remoteIP, nil
}
