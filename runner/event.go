package runner

import (
	"encoding/binary"
	"fmt"
	"gate-service/service"
	"gate-service/zoo"
	"log"
	"net"
)

type Event struct {
	api     *zoo.ZooService
	msg     service.Message
	session *service.SessionStore
}

func CreateEvent(api *zoo.ZooService, msg service.Message, session *service.SessionStore) *Event {
	return &Event{
		api:     api,
		msg:     msg,
		session: session,
	}
}

func (r *Event) Run() {
	//Checksum validating
	isChecksumValid := verifyChecksum(r.msg.Payload)
	if !isChecksumValid {
		log.Println("Checksum is not valid.")
		return
	}

	command := r.msg.Payload[2]

	switch command {
	// hearbit command
	case 0x56:
		r.onStatusHandler()
	// card event command
	case 0x53:
		r.onEventCardHandler()
	// alarm event command
	case 0x54:
		r.onEventAlarmEvent()
	}
}

func (r *Event) onStatusHandler() {
	response := []byte{0x02, 0x00, 0x56, 0xFF, 0x00, 0x08, 0x00, 0x5B, 0xA0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x58, 0x03}
	_, err := r.msg.Conn.Write(response)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	r.session.Set(r.msg.Conn.RemoteAddr().String(), string(r.msg.Payload[28:34]))
}

func (r *Event) onEventCardHandler() {
	card := binary.LittleEndian.Uint32(r.msg.Payload[7:11])
	cardstr := fmt.Sprint(card)

	log.Println("Card scanned", cardstr, r.msg.Conn.RemoteAddr())

	err := eventResponse(r.msg.Conn, 0x53, r.msg.Payload[20])
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	serial, ok := r.session.Get(r.msg.Conn.RemoteAddr().String())
	if !ok {
		fmt.Println("session get error")
	}

	var state string
	switch r.msg.Payload[18] {
	case 1:
		state = "IN"
	case 2:
		state = "OUT"
	default:
		state = "Invalid"
	}

	// Webhook function here
	resp, err := r.api.PostGateAuth(zoo.PostGateAuthBody{
		ID:        serial,
		NFCID:     cardstr,
		Direction: state,
	})
	if err != nil {
		fmt.Println("Post auth get error:", err)
		return
	}

	if resp.Data.AccessGranted {
		switch resp.Data.Action {
		case "IN":
			err := openDoor1(r.msg.Conn)
			if err != nil {
				fmt.Println("command error:", err)
				return
			}
		case "OUT":
			err := openDoor2(r.msg.Conn)
			if err != nil {
				fmt.Println("command error:", err)
				return
			}
		}
	}
}

func (r *Event) onEventAlarmEvent() {

}

func eventResponse(conn net.Conn, event byte, numberOfRecord byte) error {
	res := []byte{0x02, 0x00, event, 0xFF, 0x00, 0x01, 0x00, numberOfRecord}
	checksum := calculateChecksum(res)
	resd := append(res, checksum, 0x03)

	_, err := conn.Write(resd)
	if err != nil {
		return err
	}

	return nil
}
