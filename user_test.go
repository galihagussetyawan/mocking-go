package mockinggo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

	type User struct {
		Name    string
		Address string
	}

	server := httptest.NewServer(http.HandlerFunc(GetUser))
	response, err := http.Get(server.URL)
	if err != nil {
		t.Error(err.Error())
	}

	defer response.Body.Close()

	var j User

	json.NewDecoder(response.Body).Decode(&j)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "galih agus setyawan", j.Name)
	assert.Equal(t, "trenggalek", j.Address)
}
