package anilistapi

// Gets a response that contains a list of manga from anilist
func SearchManga(search string, page int, perPage int) (*AniMangas, error) {
	var manga = &AniMangas{}

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
			media (id: $id, search: $search, type: MANGA, sort: POPULARITY_DESC) {
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

	err := post(query, variables, manga)

	return manga, err
}

// Gets a response that contains a list of manga
// sorted according to their anilist score
func TopMangaByScore(page int, perPage int) (*AniMangas, error) {
	var manga = &AniMangas{}

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
			media (id: $id, type: MANGA, sort: SCORE_DESC) {
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

	err := post(query, variables, manga)

	return manga, err
}

// Gets a response that contains a list of manga
// sorted according to their anilist popularity
func TopMangaByPopularity(page int, perPage int) (*AniMangas, error) {
	var manga = &AniMangas{}

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
			media (id: $id, type: MANGA, sort: POPULARITY_DESC) {
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

	err := post(query, variables, manga)

	return manga, err
}

// Gets a response that contains a manga from anilist,
// according to its anilist id
func GetManga(id int) (*FullManga, error) {
	var manga *FullManga = &FullManga{}

	query := `
	query ($id: Int) {
		Media (id: $id, type: MANGA) {
			id
			title {
				romaji
				english
				native
			}
			format
			chapters
			volumes
			status
			startDate {
				day
				month
				year
			}
			averageScore
			popularity
			favourites
			source
			genres
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

	err := post(query, variables, manga)

	return manga, err
}
