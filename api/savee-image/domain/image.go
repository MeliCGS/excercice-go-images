package image_domain

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Image struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Autor string `json:"autor"`
	Alt   string `json:"alt"`
	Size  *Size  `json:"size"`
}
