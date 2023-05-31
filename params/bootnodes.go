// Copyright 2015 The go-popcateum Authors
// This file is part of the go-popcateum library.
//
// The go-popcateum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-popcateum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-popcateum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/popcateum/go-popcateum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Popcateum network.
var MainnetBootnodes = []string{
	
}

// LongcatBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Longcat test network.
var LongcatBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	
}

// YoloV3Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv3 ephemeral test network.
// TODO: Set Yolov3 bootnodes
var YoloV3Bootnodes = []string{
	
}

var V5Bootnodes = []string{
	// gcp-usa-1
	"enr:-KC4QOO_7oktDmCcwsy_mGObg6_hc8lJhHzUfntdbdE3tCKEVPbY15-_wD_4wCi_RBKWBjKyXTd1OL9LlPIKrVrrrOwDg2V0aMrJhKNYcKeDPQkAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQKimbNi1UJXr6FcSV3O10Q-1XMz-gDxDAnNsD8cXwR-doRzbmFwwIN0Y3CC7L6DdWRwguy-",

	// gcp-eu-1
	"enr:-KC4QHlvzF7qNt7Z8PU1ia3CKBclZ9p_8jEmcNZgm3hJ_XCIcXo5OJh9IrxHk41U2f-hNNMfdUECA--4SNMKsc9kCoECg2V0aMrJhKNYcKeDPQkAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQMgTTFoAdxQ_jBfHlTEICzJkMpw3LSM6vTgpdiMqtsLvYRzbmFwwIN0Y3CC7L6DdWRwguy-",

	// gcp-au-1
	"enr:-KC4QO_bmh0nYQUHFuv6la001hau2Mh6311yzizZSciHLAe9SWFzrGoy1TIz5EZtV9iRYeq9RPA7CULJtiVof2JO-IkGg2V0aMrJhKNYcKeDPQkAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQK-bR0yWS1Kjv3YDpeSJkI2SNBvjHMKJ0uLRmxep7umM4RzbmFwwIN0Y3CC7L6DdWRwguy-",

	// gcp-ko-1
	"enr:-KC4QGy0-EGaYy100TynH7pIeq3811WXlAsy5l3IGlUs3VlYYmvMIvBgowNbHWtJbja9DPmx1eM2qUtQ6KrlM36d0hYCg2V0aMrJhKNYcKeDPQkAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQOQPZxjaaI-mRAnpcqcp9wlwPntKYDcpDsISeE9fNs6HYRzbmFwwIN0Y3CC7L6DdWRwguy-",

}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/popcateum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case LongcatGenesisHash:
		net = "longcat"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".popcateum.org"
}
