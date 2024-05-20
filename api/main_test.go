package api

import (
	"os"
	"testing"
	"time"

	"github.com/ekefan/bank_panda/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/ekefan/bank_panda/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey: utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(store, config)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
