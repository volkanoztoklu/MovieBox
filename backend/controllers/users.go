package controllers

import (
	_ "fmt"
	"net/http"

	"github.com/C305DatabaseProject/database-project/backend/database"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	// Check if user with given id exists
	sql := `SELECT id, displayname, email, dateofbirth, avatar, bio, location, social_twitter, social_instagram, type
		FROM users WHERE id = ?`
	database.DB.QueryRow(sql, id).Scan(&user.ID, &user.Displayname, &user.Email, &user.DateOfBirth, &user.Avatar, &user.Bio, &user.Location, &user.Twitter, &user.Instagram, &user.Type)
	if user.ID == 0 {
		// User not found
		c.JSON(http.StatusNotFound, ErrorMessage("User not found."))
		return
	}
	// If it exists, retrieve user data
	c.JSON(http.StatusOK, OkMessage(user))
}

func Watchlist(c *gin.Context) {
	// TODO: Arda buraya bakicak
}

func Watched(c *gin.Context) {
	id := c.Param("id")
	// Check watched table with given user id
	sql := `SELECT id, title, release_date, poster, rating
		FROM movies LEFT JOIN user_watched
		ON movies.id = user_watched.movie_id
		WHERE user_watched.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {

	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, OkMessage(movies))
}

func Favorites(c *gin.Context) {
	id := c.Param("id")
	// Check Favorites table with given user id
	sql := `SELECT id, title, release_date, poster, rating
		FROM movies LEFT JOIN user_favorites
		ON movies.id = user_favorites.movie_id
		WHERE user_favorites.user_id = ?;`
	rows, err := database.DB.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(err.Error()))
		return
	}
	// Return movie details with the movie ids
	var movies []Movie
	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Poster, &movie.Rating)
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   movies,
	})
}

func Comments(c *gin.Context) {
	id := c.Param("id")
	// Check comments with user id
	sql := `SELECT users.id, users.username, movie_id , comment, comment_date 
	FROM project.user_comments LEFT JOIN project.users
	ON project.user_comments.user_id = project.users.id
	WHERE project.users.id = ?;`
	rows, _ := database.DB.Query(sql, id)
	// Return Comments details with the User ids
	var comments []Comment
	var comment Comment
	for rows.Next() {
		rows.Scan(&comment.UserID, &comment.Username, &comment.MovieID, &comment.Comment, &comment.CommentDate)
		comments = append(comments, comment)
	}
	if comments == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   "User not have any comment.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   comments,
	})
}