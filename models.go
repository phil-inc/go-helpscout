package helpscout

import (
	"encoding/json"
)

type Conversation struct {
	ID           int           `json:"id,omitempty"`
	Type         string        `json:"type,omitempty"`
	Folder       string        `json:"folder,omitempty"`
	IsDraft      bool          `json:"isDraft,omitempty"`
	Number       int           `json:"number,omitempty"`
	Owner        Person        `json:"owner,omitempty"`
	Mailbox      MailboxRef    `json:"mailbox,omitempty"`
	Customer     Person        `json:"customer,omitempty"`
	ThreadCount  int           `json:"threadCount,omitempty"`
	Status       string        `json:"status,omitempty"`
	Subject      string        `json:"subject,omitempty"`
	Preview      string        `json:"preview,omitempty"`
	CreatedBy    Person        `json:"createdBy,omitempty"`
	CreatedAt    string     `json:"createdAt,omitempty"`
	ModifiedAt   string     `json:"modifiedAt,omitempty"`
	ClosedAt     string     `json:"closedAt,omitempty"`
	ClosedBy     Person        `json:"closedBy,omitempty"`
	Source       Source        `json:"source,omitempty"`
	Cc           []string      `json:"cc,omitempty"`
	Bcc          []string      `json:"bcc,omitempty"`
	Tags         []string      `json:"tags,omitempty"`
	CustomFields []CustomField `json:"customFields,omitempty"`
	Threads      []Thread      `json:"threads,omitempty"`
}

type Attachment struct {
	ID       int    `json:"id,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Filename string `json:"filename,omitempty"`
	Size     int    `json:"size,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	URL      string `json:"url,omitempty"`
}

type AttachmentData struct {
	ID   int    `json:"id,omitempty"`
	Data string `json:"data,omitempty"`
}

type Source struct {
	Type string `json:"type,omitempty"`
	Via  string `json:"via,omitempty"`
}

type Thread struct {
	ID                int          `json:"id,omitempty"`
	AssignedTo        Person       `json:"assignedTo,omitempty"`
	Status            string       `json:"status,omitempty"`
	CreatedAt         string    `json:"createdAt,omitempty"`
	OpenedAt          string       `json:"openedAt,omitempty"`
	CreatedBy         Person       `json:"createdBy,omitempty"`
	Source            Source       `json:"source,omitempty"`
	ActionType        string       `json:"actionType,omitempty"`
	ActionSourceID    int          `json:"actionSourceId,omitempty"`
	Type              string       `json:"type,omitempty"`
	State             string       `json:"state,omitempty"`
	Customer          Person       `json:"customer,omitempty"`
	FromMailbox       MailboxRef   `json:"fromMailbox,omitempty"`
	Body              string       `json:"body,omitempty"`
	To                []string     `json:"to,omitempty"`
	Cc                []string     `json:"cc,omitempty"`
	Bcc               []string     `json:"bcc,omitempty"`
	Attachments       []Attachment `json:"attachments,omitempty"`
	SavedReplyID      int          `json:"savedReplyId,omitempty"`
	CreatedByCustomer bool         `json:"createdByCustomer,omitempty"`
}

type CustomField struct {
	ID        int      `json:"id,omitempty"`
	FieldName string   `json:"fieldName,omitempty"`
	FieldType string   `json:"fieldType,omitempty"`
	Required  bool     `json:"required,omitempty"`
	Order     int      `json:"order,omitempty"`
	Options   []Option `json:"options,omitempty"`
}

type Option struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Order int    `json:"order,omitempty"`
}

type CustomerAddress struct {
	ID         int       `json:"id,omitempty"`
	Lines      []string  `json:"lines,omitempty"`
	City       string    `json:"city,omitempty"`
	State      string    `json:"state,omitempty"`
	PostalCode string    `json:"postalCode,omitempty"`
	Country    string    `json:"country,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
}

type CustomerSocialProfile struct {
	ID    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type CustomerEmail struct {
	ID       int    `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
	Location string `json:"location,omitempty"`
}

type CustomerChat struct {
	ID    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type CustomerPhone struct {
	ID       int    `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
	Location string `json:"location,omitempty"`
}

type CustomerWebsite struct {
	ID    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type Customer struct {
	ID             int                     `json:"id,omitempty"`
	FirstName      string                  `json:"firstName,omitempty"`
	LastName       string                  `json:"lastName,omitempty"`
	PhotoURL       string                  `json:"photoUrl,omitempty"`
	PhotoType      string                  `json:"photoType,omitempty"`
	Gender         string                  `json:"gender,omitempty"`
	Age            string                  `json:"age,omitempty"`
	Organization   string                  `json:"organization,omitempty"`
	JobTitle       string                  `json:"jobTitle,omitempty"`
	Location       string                  `json:"location,omitempty"`
	Background     string                  `json:"background,omitempty"`
	CreatedAt      string               `json:"createdAt,omitempty"`
	ModifiedAt     string               `json:"modifiedAt,omitempty"`
	Address        CustomerAddress         `json:"address,omitempty"`
	SocialProfiles []CustomerSocialProfile `json:"socialProfiles,omitempty"`
	Emails         []CustomerEmail         `json:"emails,omitempty"`
	Phones         []CustomerPhone         `json:"phones,omitempty"`
	Chats          []CustomerChat          `json:"chats,omitempty"`
	Websites       []CustomerWebsite       `json:"websites,omitempty"`
}

type Mailbox struct {
	ID         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Slug       string    `json:"slug,omitempty"`
	Email      string    `json:"email,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	Folders    []Folder  `json:"folders,omitempty"`
}

type Folder struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Type        string    `json:"type,omitempty"`
	UserID      int       `json:"userId,omitempty"`
	TotalCount  int       `json:"totalCount,omitempty"`
	ActiveCount int       `json:"activeCount,omitempty"`
	ModifiedAt  string `json:"modifiedAt,omitempty"`
}

type Person struct {
	ID        int         `json:"id,omitempty"`
	FirstName string      `json:"firstName,omitempty"`
	LastName  string      `json:"lastName,omitempty"`
	Email     string      `json:"email,omitempty"`
	Phone     interface{} `json:"phone,omitempty"`
	Type      string      `json:"type,omitempty"`
}

type MailboxRef struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type User struct {
	ID         int       `json:"id,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Email      string    `json:"email,omitempty"`
	Role       string    `json:"role,omitempty"`
	Timezone   string    `json:"timezone,omitempty"`
	PhotoURL   string    `json:"photoUrl,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	Type       string    `json:"type,omitempty"`
}

type ApiError struct {
	Code  int    `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type Page struct {
	Page  int `json:"page,omitempty"`
	Pages int `json:"pages,omitempty"`
	Count int `json:"count,omitempty"`
	Items json.RawMessage
}

type Item struct {
	Item json.RawMessage
}
