package swapi

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type Film struct {
	Title        string   `json:"title"`
	EpisodeID    int64    `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

func GetFilm(id int) (Film, error) {
	var f Film
	return f, Get(fmt.Sprintf("/films/%d", id), &f)
}

func (f Film) GetCharacters() (characters []Person, err error) {
	return getPeople(f.Characters)
}

// PrintCrawl prints the opening crawl to os.Stdout
func (f Film) PrintCrawl() {
	scanner := bufio.NewScanner(strings.NewReader(f.OpeningCrawl))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(300 * time.Millisecond)
	}
}
