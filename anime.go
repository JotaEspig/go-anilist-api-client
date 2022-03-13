package anilistapi

// Gets a response that contains a list of anime from anilist
func SearchAnime(search string, page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, search: $search, type: ANIME, sort: POPULARITY_DESC) {
				id
				title {
					romaji
					english
				}
				format
			}
		}
	}
	`

	variables := map[string]interface{}{
		"search":  search,
		"page":    page,
		"perPage": perPage,
	}

	err := post(query, variables, &anime)

	return anime, err
}

// Gets a response that contains a list of anime
// sorted according to their anilist score
func TopAnimeByScore(page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, type: ANIME, sort: SCORE_DESC) {
				id
				title {
					romaji
					english
				}
				format
				averageScore
			}
		}
	}
	`

	variables := map[string]interface{}{
		"perPage": perPage,
		"page":    page,
	}

	err := post(query, variables, &anime)

	return anime, err
}

// Gets a response that contains a list of anime
// sorted according to their anilist popularity
func TopAnimeByPopularity(page int, perPage int) (AniManga, error) {
	var anime AniManga

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
			media (id: $id, type: ANIME, sort: POPULARITY_DESC) {
				id
				title {
					romaji
				}
				format
				popularity
			}
		}
	}
	`

	variables := map[string]interface{}{
		"perPage": perPage,
		"page":    page,
	}

	err := post(query, variables, &anime)

	return anime, err
}

// Gets a response that contains an anime from anilist,
// according to its anilist id
func GetAnime(id int) (FullAnime, error) {
	var anime FullAnime

	query := `
	query ($id: Int) {
		Media (id: $id, type: ANIME) {
			id
			title {
				romaji
				english
				native
			}
			format
			episodes
			status
			season
			seasonYear
			averageScore
			popularity
			favourites
			source
			genres
			characters (role: MAIN) {
				edges {
					node {
						id
						name {
							full
						}
					}
					role
				}
			}
			description
			coverImage {
				extraLarge
			}
		}
	}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	err := post(query, variables, &anime)

	return anime, err
}
