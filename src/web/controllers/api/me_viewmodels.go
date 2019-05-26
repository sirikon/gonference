package api

// User .
type User struct {
	Username string `json:"username"`
}

type ChangePasswordViewModel struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword string `json:"newPassword"`
	RepeatNewPassword string `json:"repeatNewPassword"`
}
