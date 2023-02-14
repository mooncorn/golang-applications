package src

import (
	"context"
	discovery "github.com/libp2p/go-libp2p-discovery"
	host "github.com/libp2p/go-libp2p-host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sirupsen/logrus"
	"time"
)

const service = "DavidPilarski/chatApplication"

type P2P struct {
	Ctx       context.Context
	Host      host.Host
	KadDHT    *dht.IpfsDHT
	Discovery *discovery.RoutingDiscovery
	PubSub    *pubsub.PubSub
}

func CreateP2P() *P2P {
	ctx := context.Background()
	nodehost, kaddht := setupHost(ctx)
	logrus.Debugln("Created the P2P Host and KadDHT is set too")
	bootstrapDHT(ctx, nodehost, kaddht)
	logrus.Debugln("Context layer, host, and DHT have been binded successfully")
	routingDiscovery := discovery.NewRoutingDiscovery(kaddht)
	logrus.Debugln("Discovery service is now ON and active")
	pubsubhandler := setupPubSub(ctx, nodehost, routingDiscovery)
	logrus.Debugln("Publisher and sub-publisher are set accordingly")
	return &P2P{
		Ctx:       ctx,
		Host:      nodehost,
		KadDHT:    kaddht,
		Discovery: routingDiscovery,
		PubSub:    pubsubhandler,
	}
}

func (p2p *P2P) AdvertiseConnect() {
	ttl, err := p2p.Discovery.Advertise(p2p.Ctx, service)
	logrus.Debugln("Advertising the P2P chat is done")
	time.Sleep(time.Second * 5)
	logrus.Debugf("The advertise time-out is %s", ttl)
	peerchan, err := p2p.Discovery.FindPeers(p2p.Ctx, service)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Discovery failed")
	}
	go handlePeerDiscovery(p2p.Host, peerchan)
	logrus.Debugln("Connection is now established")
}

func (p2p *P2P) Announce() {
	cidvalue := generateCID(service)
	logrus.Debugln("Service client ID is %s", cidvalue)
	err := p2p.KadDHT.Provide(p2p.Ctx, cidvalue, true)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Fatalln("Failed to announce connection")
	}
	time.Sleep(time.Second * 5)
	peerchan := p2p.KadDHT.FindProvidersAsync(p2p.Ctx, cidvalue, 0)
	logrus.Traceln("Discovered PeerChat service for peers")
	go handlePeerDiscovery(p2p.Host, peerchan)
}

func (p2p *P2P) Ann() {
	cid := generateCID(service)
	err := p2p.KadDHT.Provide(p2p.Ctx, cid, true)
	if err != nil {
		logrus.Debugln("Failed to announce connection", err.Error())
	}
	peerchan := p2p.KadDHT.FindProvidersAsync(p2p.Ctx, cid, 0)
	go handlePeerDiscovery(p2p.Host, peerchan)
}
