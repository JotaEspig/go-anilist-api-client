package anilistapi

// Top level structs

// Contains a page struct that contains animes or mangas
type AniManga struct {
	Page page `json:"Page"`
}

//Represents a page struct that contains characters
type Character struct {
	Page characterPage `json:"Page"`
}

// Represents a full data anime struct
type FullAnime struct {
	Media fullAnimeMedia `json:"Media"`
}

// Represents a full data manga struct
type FullManga struct {
	Media fullMangaMedia `json:"Media"`
}

// Represents a full data character struct
type FullCharacter struct {
	Character fullCharacterMedia `json:"Character"`
}

// Page structs

// Represents a page containing either anime or manga structs
type page struct {
	PageInfo pageInfo `json:"pageInfo"`
	Media    []media  `json:"media"`
}

// Represents a page containing character structs
type characterPage struct {
	PageInfo   pageInfo         `json:"pageInfo"`
	Characters []characterMedia `json:"characters"`
}

// Representes the information about a page,
// indifferent if it's an animanga page or a character page
type pageInfo struct {
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
	PerPage     int `json:"perPage"`
}

// Media structs

// Represents data from either an anime or a manga (limited)
type media struct {
	Id           int    `json:"id"`
	Title        title  `json:"title"`
	Format       string `json:"format"`
	AverageScore int    `json:"averageScore"`
	Popularity   int    `json:"popularity"`
}

// Represents data from a character (limited)
type characterMedia struct {
	Id         int  `json:"id"`
	Name       name `json:"name"`
	Favourites int  `json:"favourites"`
}

// Represents data from an anime (contains full data)
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

// Represents data from a manga (contains full data)
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

// Represents data from a character (contains full data)
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

// Represents a title that constains Romaji and English variations
type title struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}

// Represents a title that contains Romaji, English and/or Native variations
type fullTitle struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
	Native  string `json:"native"`
}

// Represents a name that contains no variations (contains full name only)
type name struct {
	Full string `json:"full"`
}

// Represents a name that contains Full and Native variations
type fullName struct {
	Full   string `json:"full"`
	Native string `json:"native"`
}

// Other structs

// Represents a date that a manga started
type startDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// Represents a date of birth
type dateOfBirth struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}

// Contains an extra large image link
type coverImage struct {
	ExtraLarge string `json:"extraLarge"`
}

// Contains a large image link
type image struct {
	Large string `json:"large"`
}
