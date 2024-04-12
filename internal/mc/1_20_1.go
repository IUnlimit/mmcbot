package mc

// import (
// 	"github.com/Tnze/go-mc/bot"
// 	"github.com/Tnze/go-mc/bot/basic"
// 	"github.com/Tnze/go-mc/bot/msg"
// 	"github.com/Tnze/go-mc/bot/playerlist"
// 	"github.com/Tnze/go-mc/chat"
// 	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
// 	log "github.com/sirupsen/logrus"
// )

// type Version_1_20_1 struct {
// 	client      *bot.Client
// 	player      *basic.Player
// 	chatHandler *msg.Manager
// }

// func (v *Version_1_20_1) GetClient() *bot.Client {
// 	return v.client
// }

// func (v *Version_1_20_1) GetPlayer() *basic.Player {
// 	return v.player
// }

// func (v *Version_1_20_1) SendMessage(text string) error {
// 	v.chatHandler.SendMessage(text)
// 	return nil
// }

// func NewVersion_1_20_1(mauth *GMMAuth.BotAuth) (Version, error) {
// 	// client 相当于打开的游戏程序窗口, 承载了网络连接与客户端设置
// 	client := bot.NewClient()
// 	client.Auth = bot.Auth{
// 		Name: mauth.Name,
// 		// UUID: mauth.UUID,
// 		// AsTk: mauth.AsTk,
// 	}

// 	// player 相当于加入服务器的一个游戏会话, 承载了游戏过程中角色的状态数据
// 	player := basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{
// 		// GameStart:    onGameStart,
// 		// Disconnect:   onDisconnect,
// 		// HealthChange: onHealthChange,
// 		// Death:        onDeath,
// 	})

// 	playerList := playerlist.New(client)
// 	chatHandler := msg.New(client, player, playerList, msg.EventsHandler{
// 		SystemChat:        onSystemMsg,
// 		PlayerChatMessage: onPlayerMsg,
// 		DisguisedChat:     onDisguisedMsg,
// 	})
// 	// worldManager := world.NewWorld(client, player, world.EventsListener{
// 	// 	LoadChunk:   onChunkLoad,
// 	// 	UnloadChunk: onChunkUnload,
// 	// })

// 	// screenManager = screen.NewManager(client, screen.EventsListener{
// 	// 	Open:    nil,
// 	// 	SetSlot: onScreenSlotChange,
// 	// 	Close:   nil,
// 	// })
// 	version := &Version_1_20_1{
// 		client:      client,
// 		player:      player,
// 		chatHandler: chatHandler,
// 	}

// 	return version, nil
// }

// // 系统消息（System Chat）：系统发送的提示，如“Tnze加入了游戏”，分为显示在消息栏和屏幕中央两种。
// func onSystemMsg(c chat.Message, overlay bool) error {
// 	log.Printf("System: %v, Overlay: %v", c, overlay)
// 	return nil
// }

// // 玩家消息（Player Chat）：玩家发送的消息，可以被签名确保是由玩家本人发送的。
// func onPlayerMsg(msg chat.Message, validated bool) error {
// 	var prefix string
// 	if !validated {
// 		prefix = "[Not Secure] "
// 	}
// 	log.Printf("%sPlayer: %v", prefix, msg)
// 	return nil
// }

// // 伪消息（Disguised Chat）：在服务器后台用“/say”命令发送的消息
// func onDisguisedMsg(msg chat.Message) error {
// 	log.Printf("Disguised: %v", msg)
// 	return nil
// }
