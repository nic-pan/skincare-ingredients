package models

type Ingredient struct {
	ID          int    `json:id`
	Name        string `json:name`
	Effect      string `json:effect`
	ForSkinType string `json:forSkinType`
}
