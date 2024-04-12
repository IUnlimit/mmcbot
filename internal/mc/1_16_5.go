package mc

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
	log "github.com/sirupsen/logrus"
)

type Version_1_16_5 struct {
	client *bot.Client
	player *basic.Player
}

func (v *Version_1_16_5) GetClient() *bot.Client {
	return v.client
}

func (v *Version_1_16_5) GetPlayer() *basic.Player {
	return v.player
}

func (v *Version_1_16_5) SendMessage(text string) error {

	return nil
}

func NewVersion_1_16_5(mauth *GMMAuth.BotAuth) (Version, error) {
	// client 相当于打开的游戏程序窗口, 承载了网络连接与客户端设置
	client := bot.NewClient()
	// player 相当于加入服务器的一个游戏会话, 承载了游戏过程中角色的状态数据
	player := basic.NewPlayer(client, basic.DefaultSettings)
	basic.EventsListener{
		// GameStart:  onGameStart,
		ChatMsg: onChatMsg,
		// Disconnect: onDisconnect,
		// Death:      onDeath,
	}.Attach(client)

	client.Auth = bot.Auth{
		Name: mauth.Name,
		UUID: mauth.UUID,
		AsTk: mauth.AsTk,
	}
	return &Version_1_16_5{
		client: client,
		player: player,
	}, nil
}

func onChatMsg(c chat.Message, _ byte, _ uuid.UUID) error {
	log.Println("Chat:", c.String()) // output chat message without any format code (like color or bold)
	return nil
}
