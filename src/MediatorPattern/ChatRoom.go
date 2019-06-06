package MediatorPattern

import (
	"fmt"
	"time"
)

type ChatRoom struct {
}

func (c *ChatRoom) ShowMessage(user *User, message string) {
	fmt.Println(time.Now(), "User", user.GetName(), ":", message)
}
