package rtpproxyhealth

import (
	"bytes"
	"fmt"
	"net"
)

// Check checks the health of the RTPProxy server at the given address
// by sending a token control message and checking the response.
func Check(addr string) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return fmt.Errorf("rtpproxyhealth: failed to connect: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("TOKEN V"))
	if err != nil {
		return fmt.Errorf("rtpproxyhealth: failed to write token control message: %w", err)
	}

	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("rtpproxyhealth: failed to read token control message: %w", err)
	}

	if !bytes.HasPrefix(buf[:n], []byte("TOKEN")) {
		return fmt.Errorf("rtpproxyhealth: invalid token response: %q", string(buf[:n]))
	}
	return nil
}
