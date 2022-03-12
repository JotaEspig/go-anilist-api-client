# go-anilist-api-client

## Usage:
~~~golang
package main

import "github.com/JotaEspig/go-anilist-api-client"

func main() {
    anime, err := anilistapi.SearchAnime("Fullmetal Alchemist", 0, 0)
    // First argument is the anime that you want to search
    // Second argument is the page you want to get
    // Third argument is how many animes you want to search for each page
    if err != nil {
        return
    }
}
~~~