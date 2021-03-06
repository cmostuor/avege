// +build freebsd dragonfly

package redir

import (
	"log"
	"net"

	"common"
	"inbound"
)

func handleInbound(conn *net.TCPConn, outboundHander common.OutboundHandler) {
	log.Printf("redir connect from %s, FreeBSD/DragonflyBSD is detected, use ipfw now.\n",
		conn.RemoteAddr().String())
}

func GetInboundHandler(inbound *inbound.InBound) inbound.InBoundHander {
	return handleInbound
}
