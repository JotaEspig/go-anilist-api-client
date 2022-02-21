package anilistapi

// Top level structs

type AniManga struct {
	Page page `json:"Page"`
}

type Character struct {
	Page characterPage `json:"Page"`
}

type FullAnime struct {
	Media fullAnimeMedia `json:"Media"`
}

type FullManga struct {
	Media fullMangaMedia `json:"Media"`
}

type FullCharacter struct {
	Character fullCharacterMedia `json:"Character"`
}

// Page structs

type page struct {
	PageInfo pageInfo `json:"pageInfo"`
	Media    []media  `json:"media"`
}

type characterPage struct {
	PageInfo   pageInfo         `json:"pageInfo"`
	Characters []characterMedia `json:"characters"`
}

type pageInfo struct {
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
	PerPage     int `json:"perPage"`
}

// Media structs

type media struct {
	Id           int    `json:"id"`
	Title        title  `json:"title"`
	Format       string `json:"format"`
	AverageScore int    `json:"averageScore"`
	Popularity   int    `json:"popularity"`
}

type characterMedia struct {
	Id         int  `json:"id"`
	Name       name `json:"name"`
	Favourites int  `json:"favourites"`
}

type fullAnimeMedia struct {
	Id           int        `json:"id"`
	Title        fullTitle  `json:"title"`
	Format       string     `json:"format"`
	Episodes     int        `json:"episodes"`
	Status       string     `json:"status"`
	Season       string     `json:"season"`
	SeasonYear   int        `json:"seasonYear"`
	AverageScore int        `json:"averageScore"`
	Popularity   int        `json:"popularity"`
	Favourites   int        `json:"favourites"`
	Source       string     `json:"source"`
	Genres       []string   `json:"genres"`
	Description  string     `json:"description"`
	CoverImage   coverImage `json:"coverImage"`
}

type fullMangaMedia struct {
	Id           int        `json:"id"`
	Title        fullTitle  `json:"title"`
	Format       string     `json:"format"`
	Chapters     int        `json:"chapters"`
	Volumes      int        `json:"volumes"`
	Status       string     `json:"status"`
	StartDate    startDate  `json:"startDate"`
	AverageScore int        `json:"averageScore"`
	Popularity   int        `json:"popularity"`
	Favourites   int        `json:"favourites"`
	Source       string     `json:"source"`
	Genres       []string   `json:"genres"`
	Description  string     `json:"description"`
	CoverImage   coverImage `json:"coverImage"`
}

type fullCharacterMedia struct {
	Id          int         `json:"id"`
	Name        fullName    `json:"name"`
	Gender      string      `json:"gender"`
	Age         string      `json:"age"`
	DateOfBirth dateOfBirth `json:"dateOfBirth"`
	Favourites  int         `json:"favourites"`
	Media       struct {
		Nodes []media `json:"nodes"`
	} `json:"media"`
	Description string `json:"description"`
	Image       image  `json:"image"`
}

// Title structs

type title struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}

type fullTitle struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
	Native  string `json:"native"`
}

type name struct {
	Full string `json:"full"`
}

type fullName struct {
	Full   string `json:"full"`
	Native string `json:"native"`
}

// Other structs

type startDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type dateOfBirth struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}

type coverImage struct {
	ExtraLarge string `json:"extraLarge"`
}

type image struct {
	Large string `json:"large"`
}
