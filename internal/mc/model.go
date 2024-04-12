package mc

import (
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

type Status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			Id   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	// Favicon, ForgeData
}
