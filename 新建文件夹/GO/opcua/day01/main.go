package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

func main() {
	var (
		//endpoint = flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
		//policy   = flag.String("policy", "", "Security policy: None, Basic128Rsa15, Basic256, Basic256Sha256. Default: auto")
		//mode     = flag.String("mode", "", "Security mode: None, Sign, SignAndEncrypt. Default: auto")
		//certFile = flag.String("cert", "", "Path to cert.pem. Required for security mode/policy != None")
		//keyFile  = flag.String("key", "", "Path to private key.pem. Required for security mode/policy != None")
		nodeID   = flag.String("node", "ns=2;s=Dynamic/RandomFloat", "node id to subscribe to")
		interval = flag.Duration("interval", time.Minute*1, "subscription interval")
	)
	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
	flag.Parse()

	// log.SetFlags(0)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-signalCh
		println()
		cancel()
	}()

	//endpoints, err := opcua.GetEndpoints(ctx, *endpoint)
	//if err != nil {
	//	log.Fatal(err)
	//}

	c := opcua.NewClient("opc.tcp://milo.digitalpetri.com:62541/milo", opcua.SecurityMode(ua.MessageSecurityModeNone))

	//if ep == nil {
	//	log.Fatal("Failed to find suitable endpoint")
	//}
	//opts := []opcua.Option{
	//	opcua.SecurityModeString(*mode),
	//	opcua.CertificateFile(*certFile),
	//	opcua.PrivateKeyFile(*keyFile),
	//	opcua.AuthAnonymous(),
	//	opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
	//}
	//
	//c := opcua.NewClient(ep.EndpointURL, opts...)
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	defer c.CloseWithContext(ctx)

	m, err := monitor.NewNodeMonitor(c)
	if err != nil {
		log.Fatal(err)
	}

	m.SetErrorHandler(func(_ *opcua.Client, sub *monitor.Subscription, err error) {
		log.Printf("error: sub=%d err=%s", sub.SubscriptionID(), err.Error())
	})
	wg := &sync.WaitGroup{}

	// 开始基于频道的订阅
	wg.Add(1)
	go startCallbackSub(ctx, m, *interval, 0, wg, *nodeID)

	// 开始基于频道的订阅
	wg.Add(1)
	go startCallbackSub(ctx, m, *interval, 0, wg, *nodeID)

	//wg.Add(1)
	//go startChanSub(ctx, m, *interval, 0, wg, *nodeID)

	<-ctx.Done()
	wg.Wait()
}

func startCallbackSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, wg *sync.WaitGroup, nodes ...string) {
	sub, err := m.Subscribe(
		ctx,
		&opcua.SubscriptionParameters{
			Interval: interval,
		},
		func(s *monitor.Subscription, msg *monitor.DataChangeMessage) {
			if msg.Error != nil {
				log.Printf("[callback] sub=%d error=%s", s.SubscriptionID(), msg.Error)
			} else {
				log.Printf("[callback] sub=%d ts=%s node=%s value=%v", s.SubscriptionID(), msg.SourceTimestamp.UTC().Format(time.RFC3339), msg.NodeID, msg.Value.Value())
			}
			time.Sleep(lag)
		},
		nodes...)

	if err != nil {
		log.Fatal(err)
	}

	defer cleanup(ctx, sub, wg)

	<-ctx.Done()
}

//func startChanSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, wg *sync.WaitGroup, nodes ...string) {
//	ch := make(chan *monitor.DataChangeMessage, 160)
//	sub, err := m.ChanSubscribe(ctx, &opcua.SubscriptionParameters{Interval: interval}, ch, nodes...)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer cleanup(ctx, sub, wg)
//
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case msg := <-ch:
//			if msg.Error != nil {
//				log.Printf("[channel ] sub=%d error=%s", sub.SubscriptionID(), msg.Error)
//			} else {
//				log.Printf("[channel ] sub=%d ts=%s node=%s value=%v", sub.SubscriptionID(), msg.SourceTimestamp.UTC().Format("2006-01-02 15:04:05"), msg.NodeID, msg.Value.Value())
//			}
//			time.Sleep(lag)
//		}
//	}
//}

func cleanup(ctx context.Context, sub *monitor.Subscription, wg *sync.WaitGroup) {
	log.Printf("stats: sub=%d delivered=%d dropped=%d", sub.SubscriptionID(), sub.Delivered(), sub.Dropped())
	sub.Unsubscribe(ctx)
	wg.Done()
}
