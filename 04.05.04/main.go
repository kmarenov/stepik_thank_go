package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	dur := "\""

	hours := int(time.Duration(d).Hours())
	if hours > 0 {
		dur += fmt.Sprintf("%dh", hours)
	}

	minutes := int(time.Duration(d).Minutes()) % 60
	if minutes > 0 {
		dur += fmt.Sprintf("%dm", minutes)
	}

	dur += "\""

	return []byte(dur), nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	rating := "\""

	fill := int(r)
	empty := 5 - fill

	for range fill {
		rating += "★"
	}

	for range empty {
		rating += "☆"
	}

	rating += "\""

	return []byte(rating), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	if indent > 0 {
		ind := ""
		for range indent {
			ind += " "
		}

		j, err := json.MarshalIndent(movies, "", ind)
		if err != nil {
			return "", err
		}

		return string(j), nil
	}

	j, err := json.Marshal(movies)
	if err != nil {
		return "", err
	}

	return string(j), nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
