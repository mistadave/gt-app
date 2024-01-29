package form

type GameCreate struct {
	Name        string `json:"name" binding:"required"`
	Desc        string `json:"desc" binding:"required"`
	Image       string `json:"image" binding:"optional"`
	Genre       string `json:"genre" binding:"required"`
	ReleaseDate string `json:"releaseDate" binding:"optional"`
	//Links       []Link `json:"links" binding:"optional"`
}
