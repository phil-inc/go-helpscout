package helpscout

import (
	"encoding/json"
	"time"
)

type Conversation struct {
	ID           int           `json:"id"`
	Type         string        `json:"type"`
	Folder       string        `json:"folder"`
	IsDraft      bool          `json:"isDraft"`
	Number       int           `json:"number"`
	Owner        Person        `json:"owner"`
	Mailbox      MailboxRef    `json:"mailbox"`
	Customer     Person        `json:"customer"`
	ThreadCount  int           `json:"threadCount"`
	Status       string        `json:"status"`
	Subject      string        `json:"subject"`
	Preview      string        `json:"preview"`
	CreatedBy    Person        `json:"createdBy"`
	CreatedAt    time.Time     `json:"createdAt"`
	ModifiedAt   time.Time     `json:"modifiedAt"`
	ClosedAt     time.Time     `json:"closedAt"`
	ClosedBy     Person        `json:"closedBy"`
	Source       Source        `json:"source"`
	Cc           []string      `json:"cc"`
	Bcc          []string      `json:"bcc"`
	Tags         []string      `json:"tags"`
	CustomFields []CustomField `json:"customFields"`
	Threads      []Thread      `json:"threads"`
}

type Attachment struct {
	ID       int    `json:"id"`
	MimeType string `json:"mimeType"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	URL      string `json:"url"`
}

type AttachmentData struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

type Source struct {
	Type string `json:"type"`
	Via  string `json:"via"`
}

type Thread struct {
	ID                int          `json:"id"`
	AssignedTo        Person       `json:"assignedTo"`
	Status            string       `json:"status"`
	CreatedAt         time.Time    `json:"createdAt"`
	OpenedAt          string       `json:"openedAt"`
	CreatedBy         Person       `json:"createdBy"`
	Source            Source       `json:"source"`
	ActionType        string       `json:"actionType"`
	ActionSourceID    int          `json:"actionSourceId"`
	Type              string       `json:"type"`
	State             string       `json:"state"`
	Customer          Person       `json:"customer"`
	FromMailbox       MailboxRef   `json:"fromMailbox"`
	Body              string       `json:"body"`
	To                []string     `json:"to"`
	Cc                []string     `json:"cc"`
	Bcc               []string     `json:"bcc"`
	Attachments       []Attachment `json:"attachments"`
	SavedReplyID      int          `json:"savedReplyId"`
	CreatedByCustomer bool         `json:"createdByCustomer"`
}

type CustomField struct {
	ID        int      `json:"id"`
	FieldName string   `json:"fieldName"`
	FieldType string   `json:"fieldType"`
	Required  bool     `json:"required"`
	Order     int      `json:"order"`
	Options   []Option `json:"options"`
}

type Option struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Order int    `json:"order"`
}

type CustomerAddress struct {
	ID         int       `json:"id"`
	Lines      []string  `json:"lines"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postalCode"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type CustomerSocialProfile struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type CustomerEmail struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	Location string `json:"location"`
}

type CustomerChat struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type CustomerPhone struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	Location string `json:"location"`
}

type CustomerWebsite struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Customer struct {
	ID             int                     `json:"id"`
	FirstName      string                  `json:"firstName"`
	LastName       string                  `json:"lastName"`
	PhotoURL       string                  `json:"photoUrl"`
	PhotoType      string                  `json:"photoType"`
	Gender         string                  `json:"gender"`
	Age            string                  `json:"age"`
	Organization   string                  `json:"organization"`
	JobTitle       string                  `json:"jobTitle"`
	Location       string                  `json:"location"`
	Background     string                  `json:"background"`
	CreatedAt      time.Time               `json:"createdAt"`
	ModifiedAt     time.Time               `json:"modifiedAt"`
	Address        CustomerAddress         `json:"address"`
	SocialProfiles []CustomerSocialProfile `json:"socialProfiles"`
	Emails         []CustomerEmail         `json:"emails"`
	Phones         []CustomerPhone         `json:"phones"`
	Chats          []CustomerChat          `json:"chats"`
	Websites       []CustomerWebsite       `json:"websites"`
}

type Mailbox struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Folders    []Folder  `json:"folders"`
}

type Folder struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	UserID      int       `json:"userId"`
	TotalCount  int       `json:"totalCount"`
	ActiveCount int       `json:"activeCount"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

type Person struct {
	ID        int         `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Phone     interface{} `json:"phone"`
	Type      string      `json:"type"`
}

type MailboxRef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Timezone   string    `json:"timezone"`
	PhotoURL   string    `json:"photoUrl"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Type       string    `json:"type"`
}

type ApiError struct {
	Code  int    `json:"id"`
	Error string `json:"error`
}

type Page struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Count int `json:"count"`
	Items json.RawMessage
}

type Item struct {
	Item json.RawMessage
}
