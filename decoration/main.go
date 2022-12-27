package decoration

import (
	"./detool"
	"fmt"
)

type Component interface {
	Operator(string)
}

var sentData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// send data
	sentData = data

}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncriptComponent struct {
	key string
	com Component
}

func (self *EncriptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

func main() {
	sender := &EncriptComponent{key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{}}}

	sender.Operator("Hello World")
	fmt.Print(sentData)
}
