package api

import (
	"os"
	"testing"
	"time"

	db "github.com/MirzaKarabulut/simplebank/db/sqlc"
	"github.com/MirzaKarabulut/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config {
		SymmetricTokenKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}