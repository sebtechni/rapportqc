package facest

import (
    "net/http"
    "time"
)

const (
    BaseURLV1 = "http://baton-2:8080/Baton/api/tasks/"
)

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{
        BaseURL: BaseURLV1,
        apiKey:  apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }
}