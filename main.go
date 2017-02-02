package main

import (
	"fmt"

	"github.com/phil-inc/go-helpscout/helpscout"
)

func main() {
	mailboxID := 93154
	// res := helpscout.GetMailboxes()

	// email := "chawei.hsu@gmail.com"
	// subject := "Request from Go"

	// threadBody := "Test"
	// res := helpscout.CreateConversation(mailboxID, email, subject, threadBody)

	res := helpscout.GetConversations(mailboxID)

	fmt.Println(res.Items[0]["customer"])
}
