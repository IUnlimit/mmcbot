package mmcbot

import "github.com/IUnlimit/mmcbot/internal/mc"

func Run() {
	mc.Connect("127.0.0.1:25565")
}
