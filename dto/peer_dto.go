package dto

import "github.com/libp2p/go-libp2p-core/peer"

//节点信息包括country，city，经纬度
type PeerInfo struct {
	Country   string
	City      string
	Latitude  float64
	Longitude float64
	Peer      peer.AddrInfo
}
