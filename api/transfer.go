package api

import (
	"fmt"
	"net/http"

	db "github.com/ekefan/bank_panda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type transferRequest struct {
	FromAccountID    int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID  int64 `json:"to_account_id" binding:"required,min=1"`
	Amount int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof=currency"`

}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//ctx.JSON sends response to the client ...status code, object.
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return //errorResponse return a key value, so gin can jsonify the error message
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return 
	}
	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}
	args := db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount: req.Amount,
	}

	result, err := server.store.TransferTx(ctx, args)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return false 
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}


	if account.Currency != currency {
		err := fmt.Errorf("account (%d) currency mismatch: %s vs %s", accountId, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}