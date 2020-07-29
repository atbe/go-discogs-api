package discogs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReleasesService struct {
	http *http.Client
}

func newReleasesService(http *http.Client) *ReleasesService {
	return &ReleasesService{
		http: http,
	}
}

type ReleaseModel struct {
	ID            int     `json:"id"`
	Status        string  `json:"status"`
	Year          int     `json:"year"`
	ResourceURL   string  `json:"resource_url"`
	URI           string  `json:"uri"`
	NumberForSale int     `json:"num_for_sale"`
	LowestPrice   float32 `json:"lowest_price"`
	MasterID      int     `json:"master_id"`
	MasterURL     string  `json:"master_url"`
	Title         string  `json:"title"`
	Country       string  `json:"country"`
	ReleasedOn    string  `json:"released"`
}

func (s *ReleasesService) GetRelease(id string) (*ReleaseModel, error) {
	endpoint := fmt.Sprintf("%v/releases/%v", DiscogsApiBaseUrl, id)
	log.Print(fmt.Sprintf("Calling %v", endpoint))

	response, err := s.http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var release ReleaseModel
	err = json.NewDecoder(response.Body).Decode(&release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}
