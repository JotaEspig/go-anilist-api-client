# go-anilist-api-client

## Usage:
~~~golang
package main

import ( 
    anilistapi "github.com/JotaEspig/go-anilist-api-client"
)

func main() {
    animes, err := anilistapi.SearchAnime("Fullmetal Alchemist Brotherhood", 0, 0)
    // First argument is the anime that you want to search
    // Second argument is the page you want to get, 0 uses the default value (1)
    // Third argument is how many animes you want to search for each page, 0 uses the default value (10)
    if err != nil {
        return
    }
    animeID := animes.Page.Media[0].ID
    // Get the anime ID of the first anime in the response
    anime, err := anilistapi.GetAnime(animeID)
    // Get the "full" info about the anime
    if err != nil {
        return
    }
}
~~~