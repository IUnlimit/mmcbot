package mc

import (
	"github.com/Tnze/go-mc/bot"
	"log"
)

// https://go-mc.github.io/tutorial/getting-started/join-game.html
func Demo() {
	client := bot.NewClient()
	client.Auth.Name = "IllTamer"
	client.Auth = bot.Auth{
		Name: "IllTamer",
		UUID: "58f6356e-b30c-4811-8bfc-d72a9ee99e73",
		AsTk: "*******************",
	}

	err := client.JoinServer("192.168.3.7:25565")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Login success")

	if err = client.HandleGame(); err != nil {
		log.Fatal(err)
	}
}
