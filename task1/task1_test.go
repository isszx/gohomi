package task1

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	response, err := http.Get("http://127.0.0.1:5050/anything/you?want=learn")
	assert.Nil(t, err, response)
	defer response.Body.Close()
}