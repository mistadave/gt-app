package form

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"optional"`
	Email    string `json:"email" binding:"required"`
	BirthDay string `json:"birthday" binding:"required"`
	Gender   string `json:"gender" binding:"optional"`
	PhotoURL string `json:"photo_url" binding:"optional"`
	Time     int    `json:"time" binding:"optional"`
}

type UserUpdate struct {
	Name     string `json:"name" binding:"optional"`
	Email    string `json:"email" binding:"optional"`
	BirthDay string `json:"birthday" binding:"optional"`
	Gender   string `json:"gender" binding:"optional"`
	PhotoURL string `json:"photo_url" binding:"optional"`
	Time     int    `json:"time" binding:"optional"`
}

type UserChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
