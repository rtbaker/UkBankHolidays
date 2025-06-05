package getbankholidays

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Event struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Notes   string `json:"notes"`
	Bunting bool   `json:"bunting"`
}

func (e Event) getYear() int {
	t, err := time.Parse(
		"2006-01-02",
		e.Date,
	)
	if err != nil {
		return 0
	}

	return t.Year()
}

func (e Event) String() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Date)
}

type CountryData struct {
	Division string  `json:"division"`
	Events   []Event `json:"events"`
}

type APIResponse map[string]CountryData

func FetchFromServer() (APIResponse, error) {
	res, err := http.Get("https://www.gov.uk/bank-holidays.json")

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var response APIResponse
	json.Unmarshal(body, &response)

	return response, nil
}
