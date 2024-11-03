package network

import (
	"encoding/json"
	"errors"
	"log"
	"net"

	"github.com/WeAreInSpace/dotio"
)

func HandleConn(address string) (conn net.Conn, ib dotio.Inbound, og dotio.Outgoing) {
	conn, dialE := net.Dial("tcp", address)
	if dialE != nil {
		log.Fatal(dialE)
	}
	ib = dotio.Inbound{
		Conn: conn,
	}
	og = dotio.Outgoing{
		Conn: conn,
	}
	return conn, ib, og
}

type PacketManager struct {
	Conn net.Conn
	Ib   *dotio.Inbound
	Og   *dotio.Outgoing
}

func (pm *PacketManager) GetMOTD() error {
	fristPacket := pm.Og.Write()
	sentFristPacketE := fristPacket.Sent(dotio.WriteInt32(0))
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

		log.Println(string1)
		log.Println(string2)
	}

	return nil
}

type PlayerHandshake struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func (pm *PacketManager) Handshake(phs PlayerHandshake, playerHandshakeEvent chan string) error {
	reqLogin := pm.Og.Write()
	reqLoginE := reqLogin.Sent(dotio.WriteInt32(1))
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

	sentLoginE := login.Sent(dotio.WriteInt32(1))
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

func (pm *PacketManager) FollowPlayer(playerX, playerY float64) error {
	playId := pm.Og.Write()
	sentPlayIdE := playId.Sent(dotio.WriteInt32(1))
	if sentPlayIdE != nil {
		return sentPlayIdE
	}

	playerPos := pm.Og.Write()

	playerPos.WriteFloat64(playerX)
	playerPos.WriteFloat64(playerY)

	sentPlayerPos := playerPos.Sent(dotio.WriteInt32(2))
	if sentPlayerPos != nil {
		return sentPlayerPos
	}
	return nil
}
