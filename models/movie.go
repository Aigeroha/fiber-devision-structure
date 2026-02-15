package models

type Movie struct {
	ID       int     `json:"id"`
	Name     string  `json:"title"`
	Duration int     `json:"duration"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
}
