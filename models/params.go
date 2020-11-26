package models

// define request parameters

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID int64 `json:"post_id" binding:"required"`
	// agree(1), disagree(-1), cancel_vote(0)
	Direction int8 `json:"direction"  binding:"oneof=1 0 -1"`
}
