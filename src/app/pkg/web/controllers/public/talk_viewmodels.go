package public

type AddRatingViewModel struct {
	Stars int `form:"stars"`
	Comment string `form:"comment"`
}

type AddQuestionViewModel struct {
	Question string `form:"question"`
}
