package helpscout

type Response struct {
	Page  int                      `json:"page"`
	Pages int                      `json:"pages"`
	Count int                      `json:"count"`
	Items []map[string]interface{} `json:"items"`
}
