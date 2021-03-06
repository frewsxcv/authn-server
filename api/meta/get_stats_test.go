package meta_test

import (
	"net/http"
	"testing"

	"github.com/keratin/authn-server/api/meta"
	"github.com/keratin/authn-server/api/test"
	"github.com/keratin/authn-server/lib/route"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStats(t *testing.T) {
	app := test.App()
	server := test.Server(app, meta.Routes(app))
	defer server.Close()

	app.Actives.Track(1)

	client := route.NewClient(server.URL).Authenticated(app.Config.AuthUsername, app.Config.AuthPassword)

	res, err := client.Get("/stats")
	require.NoError(t, err)
	body := test.ReadBody(res)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, []string{"application/json"}, res.Header["Content-Type"])
	assert.NotEmpty(t, body)
}
