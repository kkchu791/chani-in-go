package main

type Client struct {
	UserID string
	Stream chan string
}

func NewClient(uID string) *Client {
	return &Client{
		UserID: uID,
		Stream: make(chan string),
	}
}
