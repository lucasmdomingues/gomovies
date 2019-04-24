package gomovies

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	IMDbRating string
	IMDbVotes  string
	IMDbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	WebSite    string
	Response   string
	Error      string
}

type Rating struct {
	Source string
	Value  string
}
