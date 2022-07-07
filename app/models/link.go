package models

//easyjson:json
type Link struct {
	OriginalLink string `json:"original_link"`
	ShortLink    string `json:"short_link"`
}
