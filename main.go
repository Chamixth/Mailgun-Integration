package main

import (
	"fmt"

	"github.com/mailgun/mailgun-go"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@sandbox3cccc85c7f0f43ecbd913dfb623a77ad.mailgun.org>",
		"Hello",
		"Testing some Mailgun awesomeness helloooooo!",
		"chamith.eos@gmail.com",
	)
	_, id, err := mg.Send(m)
	return id, err
}

func main() {
	msg, err := SendSimpleMessage("sandbox3cccc85c7f0f43ecbd913dfb623a77ad.mailgun.org","42e8d00041b5e300875ca0e940a8761c-2b91eb47-7d261d7e")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(msg)
}
