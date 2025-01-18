package mediator

import "testing"

func TestMediator(t *testing.T) {
	u1 := NewUser("张三")
	u2 := NewUser("李四")

	chatRoom := NewChatRoom("私密聊天室")

	u1.Login(chatRoom)
	u2.Login(chatRoom)

	u1.Talk("你好！")
	u2.Talk("早上好，三哥！")
}
