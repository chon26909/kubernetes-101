package resty_client

import (
	"github.com/go-resty/resty/v2"
)

// NewRestyClient คืนค่า Resty Client พร้อมการตั้งค่า
func NewRestyClient() *resty.Client {
	client := resty.New()
	client.SetRetryCount(3)
	client.SetHeader("Content-Type", "application/json")
	return client
}
