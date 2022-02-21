package anilistapi

import (
	"context"

	"github.com/machinebox/graphql"
)

var url string = "https://graphql.anilist.co"

//Make a request to the API
func post(query string, variables map[string]interface{}, target interface{}) error {
	client := graphql.NewClient(url)
	req := graphql.NewRequest(query)
	for key, value := range variables {
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
