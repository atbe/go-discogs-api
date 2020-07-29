package discogs

import (
	"log"
	"net/http"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	httpClient := http.Client{}
	apiClient := NewClient(&httpClient)

	release, err := apiClient.ReleasesService.GetRelease("249504")

	if err != nil {
		t.Error(err)
	}

	log.Print(release.ID)
	log.Print(release.LowestPrice)
}
