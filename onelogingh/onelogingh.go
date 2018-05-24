package onelogingh

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fatz/ghpubkey-go/ghpubkey"
)

const OneloginGHPathf = "/authorized_keys/%s"

type OneloginGHClient struct {
	Client  *http.Client
	BaseURL string
	Pathf   string
	Rewrite map[string]string
}

// NewOneloginGHClient
func NewOneloginGHClient(baseurl string) (client *OneloginGHClient) {
	client = &OneloginGHClient{
		Client:  &http.Client{},
		BaseURL: baseurl,
		Pathf:   OneloginGHPathf,
	}
	return
}

func (c *OneloginGHClient) createURL(username string) (u *url.URL, err error) {
	u, err = url.Parse(c.BaseURL + fmt.Sprintf(c.Pathf, username))
	if err != nil {
		return nil, err
	}
	return
}

func (c *OneloginGHClient) applyRewrite(oldUsername string) (newUsername string) {
	val, ok := c.Rewrite[oldUsername]
	if ok {
		return val
	}

	return oldUsername
}

// RequestAuthorizedKeys runs a http request against pubkeyd
func (c *OneloginGHClient) RequestAuthorizedKeys(username string) (keys string, err error) {
	username = c.applyRewrite(username)
	u, err := c.createURL(username)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ERROR: Request error %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	authorizedKeys, err := ghpubkey.ParseAuthorizedKeys(data)
	if err != nil {
		return "", fmt.Errorf("ERROR: Could not parse returned authorized keys - %v", err)
	}

	return authorizedKeys.GenAuthFIle(), err
}
