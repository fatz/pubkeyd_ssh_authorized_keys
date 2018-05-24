package onelogingh

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOneloginGHClient(t *testing.T) {
	client := NewOneloginGHClient("authorizedkeys.test.example.com")

	assert.Equal(t, http.DefaultClient, client.Client)
	assert.Equal(t, "authorizedkeys.test.example.com", client.BaseURL)

}
