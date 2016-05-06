package net

import (
	"net"
)

type Addresser interface {
	Address() string
}

type NetAddress struct {
	Host string
	Port string
}

func (na NetAddress) Address() string {
	return net.JoinHostPort(na.Host, na.Port)
}
