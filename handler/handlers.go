package handler

import (
	"fmt"
	"net/http"

	"github.com/shoet/expense-manage-api-mock/store"
	"github.com/shoet/expense-manage-api-mock/util"
)

type ListExpenses struct {
	Repo *store.Repository
}

func (l *ListExpenses) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e, err := l.Repo.ListExpenses()
	if err != nil {
		fmt.Printf("failed ListExpenses: %v", err)
		errResp := util.ErrorResponse{
			Message: util.ErrorMessageInternalServerError,
		}
		util.RespondJSON(w, http.StatusInternalServerError, errResp)
		return
	}
	util.RespondJSON(w, http.StatusOK, e)
	return
}
