package helpscout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Request is the request with basic auth
func GetRequest(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("1e37c2bf4f05f2e851f9917c391612d091389e66", "X")

	resp, err := client.Do(req)

	return resp, err
}

// PostRequest is the request with basic auth
func PostRequest(url string, dataStr string) (*http.Response, error) {
	var jsonStr = []byte(dataStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.SetBasicAuth("1e37c2bf4f05f2e851f9917c391612d091389e66", "X")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp, err
}

// CreateConversation in HelpScout
func CreateConversation(mailboxID int, email string, subject string, threadBody string) string {
	url := "https://api.helpscout.net/v1/conversations.json"
	var str = fmt.Sprintf(`
  {
    "type": "chat",
    "customer": {
      "email": "%s"
    },
    "subject": "%s",
    "mailbox": {
      "id": %d
    },
    "threads": [{
      "createdBy": {
        "email": "%s",
        "type": "customer"
      },
      "type": "chat",
      "body": "%s"
    }]
  }`, email, subject, mailboxID, email, threadBody)

	fmt.Println(str)

	resp, _ := PostRequest(url, str)
	defer resp.Body.Close()

	return string(resp.Header["Location"][0])
}

// GetMailboxes returns general Response
func GetMailboxes() Response {
	var res Response

	url := "https://api.helpscout.net/v1/mailboxes.json"

	resp, err := GetRequest(url)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&res)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//htmlData, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(os.Stdout, string(htmlData))
	return res
}

// GetConversations returns general Response
func GetConversations(mailboxID int) Response {
	var res Response

	url := "https://api.helpscout.net/v1/mailboxes/" + strconv.Itoa(mailboxID) + "/conversations.json"
	fmt.Println(url)

	resp, err := GetRequest(url)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&res)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//htmlData, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(os.Stdout, string(htmlData))
	return res
}
