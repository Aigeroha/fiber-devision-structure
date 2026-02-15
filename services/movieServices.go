package services

import (
	"errors"
	"log"
	"project-two/models"
)

var movies = []models.Movie{
	{ID: 1, Name: "Зверополис", Duration: 108, Genre: "анимация, комедия", Rating: 8.1}, // 0
	{ID: 2, Name: "28 лет спустя", Duration: 115, Genre: "ужасы, триллер", Rating: 7.4},
	{ID: 3, Name: "Вечность", Duration: 102, Genre: "драма, фантастика", Rating: 6.8},
	{ID: 4, Name: "Аватар", Duration: 162, Genre: "фантастика, приключения", Rating: 8.0},
	{ID: 5, Name: "Начало", Duration: 148, Genre: "фантастика, триллер", Rating: 8.8},
}

func GetMovies() []models.Movie {
	return movies
}

func UpdateMovie(id int, data map[string]interface{}) (models.Movie, error) {
	for i, movie := range movies {

		if movie.ID == id {
			log.Println(movie.Name, movie.Duration, movie.Genre, movie.Rating)
			if name, ok := data["title"].(string); ok {
				movie.Name = name
				log.Println("title", name)
			}
			if duration, ok := data["duration"].(float64); ok {
				movie.Duration = int(duration)
				log.Println("duration", duration)
			}
			if genre, ok := data["genre"].(string); ok {
				movie.Genre = genre
				log.Println("genre", genre)
			}
			if rating, ok := data["rating"].(float64); ok {
				movie.Rating = float64(rating)
			}
			movies[i] = movie
			return movie, nil
		}
	}
	return models.Movie{}, errors.New("movie not found")
}

func GetMovieByID(id int) (models.Movie, error) {

	for _, movie := range movies {
		if movie.ID == id {
			return movie, nil
		}
	}

	return models.Movie{}, errors.New("movie not found")
}

func PostMovie(movie models.Movie) models.Movie {

	movie.ID = len(movies) + 1
	movies = append(movies, movie)

	return movie
}

func DeleteMovie(id int) (string, error) {

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return "Movie deleted", nil
		}
	}
	return "", errors.New("movie not found")
}

func PutMovie(id int, updatedMovie models.Movie) (models.Movie, error) {

	for i, movie := range movies {
		if movie.ID == id {
			updatedMovie.ID = id
			movies[i] = updatedMovie
			return updatedMovie, nil
		}
	}
	return models.Movie{}, errors.New("movie not found")
}
