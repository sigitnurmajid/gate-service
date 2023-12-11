package runner

import (
	"fmt"
	"net"
)

func openDoor1(conn net.Conn) error {
	cmd := []byte{0x02, 0x00, 0x2C, 0xFF, 0x01, 0x00, 0x00}
	checksum := calculateChecksum(cmd)
	cmdd := append(cmd, checksum, 0x03)

	_, err := conn.Write(cmdd)
	if err != nil {
		return fmt.Errorf("open door 1 write error: %v", err)
	}

	return nil
}

func openDoor2(conn net.Conn) error {
	cmd := []byte{0x02, 0x00, 0x2C, 0xFF, 0x02, 0x00, 0x00}
	checksum := calculateChecksum(cmd)
	cmdd := append(cmd, checksum, 0x03)

	_, err := conn.Write(cmdd)
	if err != nil {
		return fmt.Errorf("open door 1 write error: %v", err)
	}

	return nil
}
