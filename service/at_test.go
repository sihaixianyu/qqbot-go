package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/solywsh/chatgpt"
)

func TestChatWithContext(t *testing.T) {
	chat := chatgpt.New(apiKey, "user_id(not required)", 10*time.Second)
	defer chat.Close()
	//select {
	//case <-chat.GetDoneChan():
	//	fmt.Println("time out")
	//}
	question := "请你记住我的名字是CRT"
	answer, err := chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)
	fmt.Println("------------------------------------------------")

	question = "我的名字是什么"
	answer, err = chat.ChatWithContext(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("A: %s\n", answer)
}
