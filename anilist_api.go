package anilistapi

import (
	"context"

	"github.com/machinebox/graphql"
)

const url string = "https://graphql.anilist.co"

type validStruct interface {
	*AniMangas | *Characters | *FullAnime | *FullManga | *FullCharacter
}

// Make a post request to the API, and the result goes to target variable.
//	target must be a pointer to one of the response structs
func post[pT validStruct](query string, vars map[string]interface{}, target pT) error {
	client := graphql.NewClient(url)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	ctx := context.Background()
	err := client.Run(ctx, req, target)
	if err != nil {
		return err
	}

	return nil
}

// Check the values of "page" and "perPage", and changes their values if it's needed
func checkPageValues(page *int, perPage *int) {
	if *page < 1 {
		*page = 1
	}
	if *perPage > 20 || *perPage < 1 {
		*perPage = 10
	}
}
