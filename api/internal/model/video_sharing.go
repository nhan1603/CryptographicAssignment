package model

type VideoShare struct {
	VideoUrl   string
	UserID     int
	VideoTitle string
}

type ListSharedVideo []SharedVideo

type SharedVideo struct {
	VideoUrl    string
	VideoTitle  string
	SharerEmail string
}

type NewVideoMessage struct {
	SharerEmail string `json:"email"`
}
