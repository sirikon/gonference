package public

type AddRatingViewModel struct {
	Stars int `form:"stars"`
	Comment string `form:"comment"`
}
