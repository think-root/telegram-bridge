package repository

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func Socialify(usernameRepo string) error {
	log.Println("Starting Socialify image parsing")

	patterns := []string{"Diagonal Stripes", "Charlie Brown", "Brick Wall", "Circuit Board", "Formal Invitation", "Signal"}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPattern := patterns[rng.Intn(len(patterns))]
	socialifyUrl := fmt.Sprintf("https://socialify.git.ci/%s/png?description=0&font=Jost&forks=1&issues=1&language=1&name=1&owner=1&pattern=%s&pulls=1&stargazers=1&theme=Light", usernameRepo, randomPattern)

	response, err := http.Get(socialifyUrl)
	if err != nil {
		log.Println(err)
		return err
	}
	defer response.Body.Close()

	file, err := os.Create("./tmp/gh_project_img/image.png")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("error code: %v", response.StatusCode)
	}

	log.Println("Socialify image parsing finished")

	return nil
}
