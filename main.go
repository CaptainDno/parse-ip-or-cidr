package parse_ip_or_cidr

import (
	"errors"
	"net"
	"net/netip"
	"strings"
)

func ParseIPOrCIDR(raw string) (*net.IPNet, error) {
	if raw == "" {
		return nil, errors.New("address can not be empty")
	}

	if strings.Contains(raw, "/") {
		_, ipNet, err := net.ParseCIDR(raw)
		if err != nil {
			return nil, err
		}
		return ipNet, nil

	} else {
		addr, err := netip.ParseAddr(raw)
		if err != nil {
			return nil, err
		}

		ipNet := &net.IPNet{
			IP: addr.AsSlice(),
		}
		if addr.Is4() {
			ipNet.Mask = net.CIDRMask(32, 32)
		} else if addr.Is6() {
			ipNet.Mask = net.CIDRMask(128, 128)
		}
		return ipNet, nil
	}
}
