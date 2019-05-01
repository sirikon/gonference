package auth

import "net/http"

const cookieName = "gonference.session"

func createSessionCookie() *http.Cookie {
	return &http.Cookie{
		Name:  cookieName,
		Value: "asdf",
	}
}

// EnsureCookie .
func EnsureCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		cookie = createSessionCookie()
		http.SetCookie(w, cookie)
	}
	return cookie
}
