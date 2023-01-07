package auth

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenerateAccessCookie(value string) *fiber.Cookie {
	accessCookie := &fiber.Cookie{
		Name:     "access-token",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Minute * 15),
		Value:    value,
		Path:     "/",
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}

	return accessCookie
}

func GenerateRefreshCookie(value string) *fiber.Cookie {
	refreshCookie := &fiber.Cookie{
		Name:     "refresh-token",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		Value:    value,
		Path:     "/",
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}

	return refreshCookie
}

func GenerateFakeAccessCookie() *fiber.Cookie {
	accessCookie := &fiber.Cookie{
		Name:     "access-token",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Minute),
		Value:    "",
		Path:     "/",
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}

	return accessCookie

}
func GenerateFakeRefreshCookie() *fiber.Cookie {
	refreshCookie := &fiber.Cookie{
		Name:     "refresh-token",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Minute),
		Value:    "",
		Path:     "/",
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}

	return refreshCookie

}

func GetCookieValue(request *http.Request, cookieString string) (string, error) {
	cookie, err := request.Cookie(cookieString)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
