package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var (
	flagRtpproxyAddr = flag.String("rtpproxy", "127.0.0.1:7722", "Rtpproxy address")
	flagAddr         = flag.String("addr", ":8080", "Address the server will listen on")
)

func main() {
	flag.Parse()

	log.Info("Rtpproxy addr is: ", *flagRtpproxyAddr)
	handler := NewAppHandler(*flagRtpproxyAddr)

	log.Info("Server listening on: ", *flagAddr)
	log.Fatal(handler.Serve(*flagAddr))
}
