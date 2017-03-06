package helpscout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	// Helpscout API key
	Key string
	// Helpscout API base. e.g. "https://api.helpscout.net/v1/"
	BaseURL string
	// Requests are transported through this client
	HTTPClient *http.Client
}

func ClientWithKey(key string) *Client {
	return &Client{
		Key:        key,
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.helpscout.net/v1/",
	}
}

func (c *Client) url(uri string) string {
	return c.BaseURL + uri
}

func (c *Client) Get(uri string) []byte {
	req, err := http.NewRequest("GET", c.url(uri) , nil)
	req.SetBasicAuth(c.Key, "X")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (c *Client) Put(uri string, reqBody* bytes.Buffer) http.Header {

	req, err := http.NewRequest("PUT", c.url(uri), reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Key, "X")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if body == nil {
		panic("POST request to helpscout failed")
	}
	return resp.Header
}

func (c *Client) CreateCustomer(customer Customer) {
	uri := fmt.Sprintf("customers.json")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(customer)
	c.Post(uri, b)
}

func getIDFromLocation(location string) string {
	var IDWithExt string
	for i := len(location) - 1; i >= 0; i-- {
		if location[i] == '/' {
			IDWithExt = location[i+1:]
			break
		}
	}
	var ID string
	for i := 0; i<len(IDWithExt) ; i++ {
		if IDWithExt[i] == '.' {
			return IDWithExt[:i]
		} 
	}	
	return ID
}

func getIDFromResponseHeader(respHeader http.Header) string  {
	var ID string
	for k, v := range respHeader {
		if(k == "Location") {
			ID = getIDFromLocation(v[0])
		}
	}
	return ID
}

func (c *Client) Post(uri string, reqBody* bytes.Buffer) http.Header {
	fmt.Println(c.url(uri))
	req, err := http.NewRequest("POST", c.url(uri), reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Key, "X")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if body == nil {
		panic("POST request to helpscout failed")
	}
	
	return resp.Header
}

func (c *Client) CreateConversation(conversation Conversation) string {
	uri := fmt.Sprintf("conversations.json")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(conversation)
	respHeader := c.Post(uri, b)
	conversationID := getIDFromResponseHeader(respHeader)
	return conversationID
}

func (c *Client) CreateConversationThread(conversationThread Thread, conversationID string) {
	uri := fmt.Sprintf("conversations/%s.json", conversationID)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(conversationThread)
	c.Post(uri, b)
}

func (c *Client) UpdateConversationThread(conversationThread Thread, conversationID string, conversationThreadID string) {
    uri := fmt.Sprintf("conversations/%s/threads/%s.json", conversationID, conversationThreadID)
    b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(conversationThread)
	c.Put(uri, b)
}

func (c *Client) getPage(uri string) *Page {
	rawPage := c.Get(uri)
	var p = new(Page)
	err := json.Unmarshal(rawPage, &p)
	if err != nil {
		panic(err)
	}
	return p
}

func (c *Client) getItem(uri string) *Item {
	rawItem := c.Get(uri)
	var i = new(Item)
	err := json.Unmarshal(rawItem, &i)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Client) getConversations(uri string) []Conversation {
	page := c.getPage(uri)
	var coversationList []Conversation
	if err := json.Unmarshal([]byte(page.Items), &coversationList); err != nil {
		panic(err)
	}
	return coversationList
}

func (c *Client) GetConversation(conversationID string) Conversation {
	uri := fmt.Sprintf("conversations/%s.json", conversationID)
	item := c.getItem(uri)
	var conversation Conversation
	if err := json.Unmarshal([]byte(item.Item), &conversation); err != nil {
		panic(err)
	}
	return conversation
}
func (c *Client) GetConversationsFromMailbox(mailboxID string) []Conversation {
	uri := fmt.Sprintf("mailboxes/%s/conversations.json", mailboxID)
	return c.getConversations(uri)
}

func (c *Client) GetConversationsFromMailboxFolder(mailboxID string, folderID string) []Conversation {
	uri := fmt.Sprintf("mailboxes/%s/folders/%s/conversations.json", mailboxID, folderID)
	return c.getConversations(uri)
}

func (c *Client) GetConversationsFromMailboxForCustomer(mailboxID string, customerID string) []Conversation {
	uri := fmt.Sprintf("mailboxes/%s/customer/%s/conversations.json", mailboxID, customerID)
	return c.getConversations(uri)
}

func (c *Client) GetConversationsFromMailboxForUser(mailboxID string, userID string) []Conversation {
	uri := fmt.Sprintf("mailboxes/%s/user/%s/conversations.json", mailboxID, userID)
	return c.getConversations(uri)
}

func (c *Client) getCustomers(uri string) []Customer {
	page := c.getPage(uri)
	var customerList []Customer
	if err := json.Unmarshal([]byte(page.Items), &customerList); err != nil {
		panic(err)
	}
	return customerList
}

func (c *Client) GetCustomers() []Customer {
	uri := fmt.Sprintf("customers.json")
	return c.getCustomers(uri)
}

func (c *Client) GetCustomer(customerID string) Customer {
	uri := fmt.Sprintf("customers/%s.json", customerID)
	item := c.getItem(uri)
	var customer Customer
	if err := json.Unmarshal([]byte(item.Item), &customer); err != nil {
		panic(err)
	}
	return customer
}

func (c *Client) GetCustomersForMailbox(customerID string) []Customer {
	uri := fmt.Sprintf("mailboxes/%s/customers.json", customerID)
	return c.getCustomers(uri)
}

// GetMailboxes returns all mailboxes
func (c *Client) GetMailboxes() []Mailbox {
	uri := fmt.Sprintf("mailboxes.json")
	page := c.getPage(uri)
	var mailboxList []Mailbox
	if err := json.Unmarshal([]byte(page.Items), &mailboxList); err != nil {
		panic(err)
	}
	return mailboxList
}

// GetMailbox returns mailbox for specified mailboxId
func (c *Client) GetMailbox(mailboxID string) Mailbox {
	uri := fmt.Sprintf("mailboxes/%s.json", mailboxID)
	item := c.getItem(uri)
	var mailbox Mailbox
	if err := json.Unmarshal([]byte(item.Item), &mailbox); err != nil {
		panic(err)
	}
	return mailbox
}

func (c *Client) getUsers(url string) []User {
	page := c.getPage(url)
	var userList []User
	if err := json.Unmarshal([]byte(page.Items), &userList); err != nil {
		panic(err)
	}
	return userList
}

func (c *Client) GetUsers() []User {
	uri := fmt.Sprintf("users.json")
	return c.getUsers(uri)
}

func (c *Client) GetUser(userID string) User {
	uri := fmt.Sprintf("users/%s.json", userID)
	item := c.getItem(uri)
	var user User
	if err := json.Unmarshal([]byte(item.Item), &user); err != nil {
		panic(err)
	}
	return user
}