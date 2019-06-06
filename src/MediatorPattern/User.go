package MediatorPattern

type User struct {
	name string
}

func (u *User) User(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SendMessage(message string) {
	chatRoom:=new(ChatRoom)
	chatRoom.ShowMessage(u,message)
}
