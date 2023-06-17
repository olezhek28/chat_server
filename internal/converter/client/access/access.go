package access

import (
	"github.com/olezhek28/chat_server/internal/model"
	desc "github.com/olezhek28/chat_server/pkg/chat_v1"
)

func ToCreateChatService(req *desc.CreateChatRequest) *model.CreateChat {
	return &model.CreateChat{
		Usernames: req.GetUsernames(),
	}
}
