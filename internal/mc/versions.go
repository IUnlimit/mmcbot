package mc

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
)

type Version interface {
	GetClient() *bot.Client
	GetPlayer() *basic.Player
	SendMessage(text string) error
}
