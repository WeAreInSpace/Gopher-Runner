package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/WeAreInSpace/Gopher-Runner/packet"
)

func HandleConn(address string) (conn net.Conn, ib packet.Inbound, og packet.Outgoing) {
	conn, dialE := net.Dial("tcp", address)
	if dialE != nil {
		log.Fatal(dialE)
	}
	ib = packet.Inbound{
		Conn: conn,
	}
	og = packet.Outgoing{
		Conn: conn,
	}
	return conn, ib, og
}

type PacketManager struct {
	Conn net.Conn
	Ib   *packet.Inbound
	Og   *packet.Outgoing
}

func (pm *PacketManager) GetMOTD() error {
	fristPacket := pm.Og.Write()
	sentFristPacketE := fristPacket.Sent(packet.WriteInt32(0))
	if sentFristPacketE != nil {
		return sentFristPacketE
	}

	motdId, motd, readMotdE := pm.Ib.Read()
	if readMotdE != nil {
		return readMotdE
	}

	if motdId == 2 {
		string1 := motd.ReadString()
		string2 := motd.ReadString()

		fmt.Println(string1)
		fmt.Println(string2)
	}

	return nil
}

type PlayerHandshake struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func (pm *PacketManager) Handshake(phs PlayerHandshake, playerHandshakeEvent chan string) error {
	reqLogin := pm.Og.Write()
	reqLoginE := reqLogin.Sent(packet.WriteInt32(1))
	if reqLoginE != nil {
		playerHandshakeEvent <- "ERR"
		return reqLoginE
	}

	login := pm.Og.Write()
	playerData, jsonPlayerDataE := json.Marshal(phs)
	if jsonPlayerDataE != nil {
		playerHandshakeEvent <- "ERR"
		return jsonPlayerDataE
	}

	login.WriteString(string(playerData))

	sentLoginE := login.Sent(packet.WriteInt32(1))
	if sentLoginE != nil {
		playerHandshakeEvent <- "ERR"
		return sentLoginE
	}

	loginResId, loginRes, loginResE := pm.Ib.Read()
	if loginResE != nil {
		playerHandshakeEvent <- "ERR"
		return loginResE
	}
	switch loginResId {
	case 0:
		playerHandshakeEvent <- "OK"
		return nil
	case 1:
		loginResMessage := loginRes.ReadString()
		log.Printf("ERROR: %s\n", loginResMessage)
		playerHandshakeEvent <- "ERR"
		return errors.New(loginResMessage)
	}

	playerHandshakeEvent <- "OK"

	return nil
}

func (pm *PacketManager) FollowPlayer() error {
	fristPacket := pm.Og.Write()
	sentFristPacketE := fristPacket.Sent(packet.WriteInt32(0))
	if sentFristPacketE != nil {
		return sentFristPacketE
	}

	motdId, motd, readMotdE := pm.Ib.Read()
	if readMotdE != nil {
		return readMotdE
	}

	if motdId == 2 {
		string1 := motd.ReadString()
		string2 := motd.ReadString()

		fmt.Println(string1)
		fmt.Println(string2)
	}
	return nil
}
