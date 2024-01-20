package main

import (
	"flag"

	"github.com/mauri870/rtpproxy-http-monitor/internal/rtpproxyhealth"
	log "github.com/sirupsen/logrus"
)

var (
	flagRtpproxyAddr = flag.String("rtpproxy", "127.0.0.1:7722", "Rtpproxy address")
	flagAddr         = flag.String("addr", "", "Address the server will listen on")
)

func main() {
	flag.Parse()

	// if no addr is provided assume command line mode
	if *flagAddr == "" {
		err := rtpproxyhealth.Check(*flagRtpproxyAddr)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	log.Info("Rtpproxy addr is: ", *flagRtpproxyAddr)
	handler := NewAppHandler(*flagRtpproxyAddr)

	log.Info("Server listening on: ", *flagAddr)
	log.Fatal(handler.Serve(*flagAddr))
}
