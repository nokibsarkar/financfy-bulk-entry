package handler

import (
	"net/http"

	"github.com/adibaruet/financfy-backend/core"
	"github.com/adibaruet/financfy-backend/repository"
	"github.com/adibaruet/financfy-backend/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Server(client *pgxpool.Pool) {
	AccountRepositoryDatabse := repository.NewAccountRepository(client)
	accountHandler := AccountHandler{service: service.NewAccountService(AccountRepositoryDatabse)}

	apiRouter := chi.NewRouter()

	apiRouter.Post(core.ApiBasePath+"/auth/account", accountHandler.createAccount)
	apiRouter.Post(core.ApiBasePath+"/auth/login", accountHandler.login)

	http.ListenAndServe(":8080", apiRouter)
}
