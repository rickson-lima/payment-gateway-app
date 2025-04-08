package dto

import (
	"time"

	"github.com/rickson-lima/payment-gateway-app/go-gateway-api/internal/domain"
)

type CreateAccountInput struct {
	Name  string `json:"name"`  // Name string
	Email string `json:"email"` // Email string
}

type AccountOutput struct {
	ID        string    `json:"id"`                // ID string
	Name      string    `json:"name"`              // Name string
	Email     string    `json:"email"`             // Email string
	Balance   float64   `json:"balance"`           // Balance string
	APIKey    string    `json:"api_key,omitempty"` // APIKey string
	CreatedAt time.Time `json:"created_at"`        // CreatedAt string
	UpdatedAt time.Time `json:"update_at"`         // UpdateAt string
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
