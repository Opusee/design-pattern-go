package mediator

import "fmt"

/**
中介模式
*/

type User struct {
	name     string    //名字
	chatRoom *ChatRoom //聊天室引用
}

// NewUser 初始化必须起名字
func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) GetName() string {
	return u.name
}

// Login 用户登陆
func (u *User) Login(chatRoom *ChatRoom) {
	chatRoom.Connect(u)   //调用聊天室连接方法
	u.chatRoom = chatRoom //注入聊天室引用
}

// Talk 用户发言
func (u *User) Talk(msg string) {
	u.chatRoom.SendMsg(u, msg) //给聊天室发消息
}

func (u *User) Listen(fromWhom *User, msg string) {
	fmt.Printf("[%s 的对话框] %s 说：%s\n", u.name, fromWhom.GetName(), msg)
}

type ChatRoom struct {
	name  string  //聊天室命名
	users []*User //聊天室里的用户们
}

// NewChatRoom 初始化必须命名聊天室
func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		name:  name,
		users: make([]*User, 0),
	}
}

func (c *ChatRoom) Connect(user *User) {
	c.users = append(c.users, user) //用户进入聊天室加入列表。
	fmt.Printf("欢迎 [%s] 加入聊天室 [%s]\n", user.GetName(), c.name)
}

func (c *ChatRoom) SendMsg(fromWhom *User, msg string) {
	// 循环所有用户，只发消息给非发送方fromWhom。
	for _, user := range c.users {
		if user.GetName() != fromWhom.GetName() {
			user.Listen(fromWhom, msg) //循环所有用户,发送消息给非发送方fromWhom。
		}
	}
}
