package handler

import (
	"encoding/json"
	"net/http"
	"time"

	exception "github.com/adibaruet/financfy-backend/Exception"
	"github.com/adibaruet/financfy-backend/core"
	"github.com/adibaruet/financfy-backend/dto"
	"github.com/adibaruet/financfy-backend/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var account dto.CreateAccountDTO

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unexpected fields in the JSON payload

	err := decoder.Decode(&account)
	if err != nil {
		core.ResponseWriter(w, http.StatusBadRequest, exception.Failure{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body: " + err.Error(),
		})
		return
	}

	newAccount, failure := a.service.Store(account)
	if failure != nil {
		core.ResponseWriter(w, failure.StatusCode, failure)
	} else {
		core.ResponseWriter(w, http.StatusCreated, newAccount)
	}
}

func (a AccountHandler) login(w http.ResponseWriter, r *http.Request) {
	var login dto.LoginDTO
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unexpected fields in the JSON payload

	err := decoder.Decode(&login)
	if err != nil {
		core.ResponseWriter(w, http.StatusBadRequest, exception.Failure{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body: " + err.Error(),
		})
		return
	}

	// Call the service to validate login
	account, failure := a.service.Login(login)

	if failure != nil {
		// Handle failure and respond
		core.ResponseWriter(w, failure.StatusCode, failure)
		return
	}

	// If login succeeds, set the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",                      // Set the cookie name to "user_id"
		Value:    account.ID.String(),            // Use the user_id from the AccountResponseDTO
		Expires:  time.Now().Add(24 * time.Hour), // Set expiration for 24 hours
		HttpOnly: true,                           // Prevent JavaScript access
		Secure:   false,                          // Ensure it's only sent over HTTPS
		SameSite: http.SameSiteStrictMode,        // Prevent CSRF attacks
	})

	// Respond with the account details
	core.ResponseWriter(w, http.StatusOK, account)
}
