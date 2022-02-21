package anilistapi

func SearchCharacter(search string, page int, perPage int) (Character, error) {
	var character Character

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

	err := post(query, variables, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}

func TopCharactersByFavourites(page int, perPage int) (Character, error) {
	var character Character

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

	err := post(query, variables, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}

func GetCharacter(id int) (FullCharacter, error) {
	var character FullCharacter

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

	err := post(query, variables, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}
