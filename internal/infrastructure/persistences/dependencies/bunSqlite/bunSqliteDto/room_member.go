package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

// RoomMember is a struct that represents a chat member entity.
// Technically, it's a many-to-many relationship between Room and User entities.
type RoomMember struct {
	bun.BaseModel

	CreatedAt time.Time `bun:",notnull"`
	DeletedAt time.Time `bun:",notnull"`

	UserID   string        `bun:",pk,notnull"`
	RoomID   string        `bun:",pk,notnull"`
	Messages []RoomMessage `bun:"rel:has-many,join:user_id=user_id,join:room_id=room_id"`
}

func NewRoomMember(member model.RoomMember) RoomMember {
	return RoomMember{
		CreatedAt: member.CreatedAt,
		DeletedAt: member.DeletedAt,

		UserID: member.UserID,
		RoomID: member.RoomID,
	}
}

func (dto RoomMember) ToDomain() model.RoomMember {
	return model.RoomMember{
		CreatedAt: dto.CreatedAt,
		DeletedAt: dto.DeletedAt,

		UserID: dto.UserID,
		RoomID: dto.RoomID,
	}
}
