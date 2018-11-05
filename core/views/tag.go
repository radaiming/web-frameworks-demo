package views

type CreateTagRequest struct {
	Message string `json:"message"`
	Object  string `json:"object"`
	Tag     string `json:"tag"`
	Tagger  struct {
		Date  string `json:"date"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"tagger"`
	Type string `json:"type"`
}

type CreateTagResponse struct {
	Message string `json:"message"`
	NodeID  string `json:"node_id"`
	Object  struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"object"`
	Sha    string `json:"sha"`
	Tag    string `json:"tag"`
	Tagger struct {
		Date  string `json:"date"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"tagger"`
	URL          string `json:"url"`
	Verification struct {
		Payload   interface{} `json:"payload"`
		Reason    string      `json:"reason"`
		Signature interface{} `json:"signature"`
		Verified  bool        `json:"verified"`
	} `json:"verification"`
}
