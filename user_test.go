package mockinggo

import (
	"encoding/json"
	"fmt"
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

func TestGetAllUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetAllUsers))
	response, err := http.Get(server.URL)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGetUserById(t *testing.T) {

	t.Run("test_with_id_1", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user?id=%v", 1), nil)
		w := httptest.NewRecorder()
		GetUserById(w, r)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, r.URL.Query().Get("id"), "1")
		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("test_with_id_nill", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/user", nil)
		w := httptest.NewRecorder()
		GetUserById(w, r)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	})

}
