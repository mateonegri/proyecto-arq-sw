package dto

type InsertImageDto struct {
	Url    string `json:"url"`
}

type InsertImagesDto struct {
	Images []InsertImageDto `json:"images"`
}