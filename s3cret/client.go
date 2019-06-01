package s3cret

import "github.com/aws/aws-sdk-go/aws/session"

type Client struct {
	session session.Session
}

// NewClient creates a new s3cret client
func NewClient(session session.Session) *Client {
	return &Client{
		session: session,
	}
}
