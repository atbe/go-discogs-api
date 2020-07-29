package discogs

import (
	"net/http"

	"github.com/dghubble/sling"
)

const DiscogsApiBaseUrl = "https://api.discogs.com"

type Client struct {
	sling           *sling.Sling
	ReleasesService *ReleasesService
}

func NewClient(client *http.Client) *Client {
	base := sling.New().Client(client).Base(DiscogsApiBaseUrl)
	base = base.Set("Content-Type", "application/json")
	return &Client{
		sling:           base,
		ReleasesService: newReleasesService(client),
	}
}
