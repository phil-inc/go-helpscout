package helpscout

type Mailbox struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
