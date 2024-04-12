package mc

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Tnze/go-mc/bot"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
	log "github.com/sirupsen/logrus"
)

// https://go-mc.github.io/tutorial/getting-started/join-game.html
func Connect(address string) {
	status, delay, err := Ping(address)
	if err != nil {
		log.Fatalf("Target server '%s' ping failed, please check !", address)
	}
	log.Infof("[%s] (%d) is active on '%s' with players %d/%d - delay: %s.",
		status.Version.Name, status.Version.Protocol, address, status.Players.Online, status.Players.Max, delay)
	log.Infof("Description: %s", status.Description.String())

	mauth, err := Authorization()
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	version, err := NewVersion_1_16_5(&mauth)
	if err != nil {
		log.Fatalf("Init client error: %v", err)
	}
	client := version.GetClient()

	err = client.JoinServer(address)
	if err != nil {
		log.Fatalf("Join server failed: %v", err)
	}

	log.Info("Login success")

	var packetErr bot.PacketHandlerError
	for {
		err = client.HandleGame()
		if !errors.As(err, &packetErr) {
			log.Fatalf("Unknown error occurred, server exiting with: %v", err)
		}
		log.Warnf("Packet process error: %v", packetErr)
	}
}

func Authorization() (GMMAuth.BotAuth, error) {
	mauth, err := GMMAuth.GetMCcredentials("./auth.cache", "88650e7e-efee-4857-b9a9-cf580a00ef43")
	var invalid GMMAuth.BotAuth
	if err != nil {
		return invalid, err
	}
	log.Print("Authenticated as ", mauth.Name, " (", mauth.UUID, ")")
	return mauth, nil
}

func Ping(address string) (*Status, time.Duration, error) {
	bytes, delay, err := bot.PingAndListTimeout(address, 3*time.Second)
	if err != nil {
		return nil, 0, err
	}

	var status Status
	err = json.Unmarshal(bytes, &status)
	if err != nil {
		return nil, 0, err
	}
	return &status, delay, nil
}
