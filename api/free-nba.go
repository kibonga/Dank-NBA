package api

import (
	"Dank-NBA/datatypes"
	"fmt"
	"io"
	"net/http"
)

const playerApiUrl = "https://free-nba.p.rapidapi.com/players/%d"

func Get(id int) (*datatypes.Player, error) {
	url := fmt.Sprintf(playerApiUrl, id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "b5a5ba6804msh7f8952479a3d44cp1b96b4jsn8f15a3f4f68d")
	req.Header.Add("X-RapidAPI-Host", "free-nba.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retreive data, status code = %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data")
	}
	json := string(body)
	fmt.Println(json)
	return &datatypes.Player{
		FirstName:    "Lebron",
		LastName:     "James",
		HeightFeet:   6,
		HeightInches: 6,
		Position:     "PF",
	}, nil
}
