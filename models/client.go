package models

import (
	"fmt"
	"strings"

	oauth2 "github.com/Bifang-Bird/goOauth2"
)

// Client client model
type Client struct {
	ID        string
	Secret    string
	Domain    string
	Public    bool
	UserID    string
	Password  string
	Account   string
	GrantType oauth2.GrantType
}

// GetID client id
func (c *Client) GetID() string {
	return c.ID
}

// GetSecret client secret
func (c *Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *Client) GetDomain() string {
	return c.Domain
}

// IsPublic public
func (c *Client) IsPublic() bool {
	return c.Public
}

// GetUserID user id
func (c *Client) GetUserID() string {
	return c.UserID
}

// GetUserID user id
func (c *Client) GetAccount() string {
	return c.Account
}

func (c *Client) GetGrantType() oauth2.GrantType {
	return c.GrantType
}

func (c *Client) VerifyPassword(password string) bool {
	booll := strings.EqualFold(c.GetSecret(), password)
	fmt.Printf("VerifyPassword c.password = %v , password = %v , bool= %v\n", c.Password, password, booll)
	return booll

}
