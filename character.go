package anilistapi

// Gets a response that contains a list of characters from anilist
func SearchCharacter(search string, page int, perPage int) (*Characters, error) {
	var char = &Characters{}

	checkPageValues(&page, &perPage)

	query := `
	query ($id: Int, $page: Int, $perPage: Int, $search: String) {
		Page (page: $page, perPage: $perPage) {
			pageInfo {
				total
				currentPage
				lastPage
				perPage
			}
			characters (id: $id, search: $search, sort: FAVOURITES_DESC) {
				id
				name {
					full
				}
			}
		}
	}
	`

	variables := map[string]interface{}{
		"search":  search,
		"page":    page,
		"perPage": perPage,
	}

	err := post(query, variables, char)

	return char, err
}

// Gets a response that contains a list of characters
// sorted according to their amount of favourites on anilist
func TopCharactersByFavourites(page int, perPage int) (*Characters, error) {
	var char = &Characters{}

	checkPageValues(&page, &perPage)

	query := `
	query ($id: Int, $page: Int, $perPage: Int) {
		Page (page: $page, perPage: $perPage) {
			pageInfo {
				total
				currentPage
				lastPage
				perPage
			}
			characters (id: $id, sort: FAVOURITES_DESC) {
				id
				name {
					full
				}
				favourites
			}
		}
	}
	`

	variables := map[string]interface{}{
		"page":    page,
		"perPage": perPage,
	}

	err := post(query, variables, char)

	return char, err
}

// Gets a response that contains a character from anilist,
// according to its anilist id
func GetCharacter(id int) (*FullCharacter, error) {
	var char = &FullCharacter{}

	query := `
	query ($id: Int) {
		Character (id: $id) {
			id
			name {
				full
				native
			}
			gender
			age
			dateOfBirth {
				day
				month
			}
			favourites
			media (sort: POPULARITY_DESC) {
				nodes {
					id
					title {
						romaji
					}
					format
				}
			}
			description
			image {
				large
			}
		}
	}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	err := post(query, variables, char)

	return char, err
}
