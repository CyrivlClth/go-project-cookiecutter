// package model
//
// It's an example, you can delete this file.
// @filename: article.go
package model

type Article struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
