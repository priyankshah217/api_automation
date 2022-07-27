package article

import "time"

type Article struct {
	TagList     []string `json:"tagList"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
}

type Request struct {
	Article Article `json:"article"`
}

type Response struct {
	Article struct {
		Slug        string    `json:"slug"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Body        string    `json:"body"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		TagList     []string  `json:"tagList"`
		Author      struct {
			Username string      `json:"username"`
			Bio      interface{} `json:"bio"`
			Image    string      `json:"image"`
		} `json:"author"`
		FavoritedBy []interface{} `json:"favoritedBy"`
		Count       struct {
			FavoritedBy int `json:"favoritedBy"`
		} `json:"_count"`
		FavoritesCount int  `json:"favoritesCount"`
		Favorited      bool `json:"favorited"`
	} `json:"article"`
}
