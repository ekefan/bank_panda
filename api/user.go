package api



// TODO: write test for creating user api
import (
	"net/http"
	"time"

	db "github.com/ekefan/bank_panda/db/sqlc"
	"github.com/ekefan/bank_panda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//ctx.JSON sends response to the client ...status code, object.
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return //errorResponse return a key value, so gin can jsonify the error message
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := struct {
		Username          string    `json:"username"`
		Fullname          string    `json:"full_name"`
		Email             string    `json:"email"`
		PasswordChangedAt time.Time `json:"password_changed_at"`
		CreatedAt         time.Time `json:"created_at"`
	}{
		Username: user.Username,
		Fullname: user.FullName,
		Email:    user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
 
	ctx.JSON(http.StatusOK, resp)
}
