package validators

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

)

func ValidatePassword(password string) (bool, string) {
	var erros []string

	if utf8.RuneCountInString(password) < 8 {
		erros = append(erros, "password must be at least 8 characters long")
	}

	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		erros = append(erros, "password must contain at least one lowercase letter")
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		erros = append(erros, "password must contain at least one capital letter")
	}

	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		erros = append(erros, "password must contain at least one number")
	}

	if matched, _ := regexp.MatchString(`[!@#\$%\^&\*(),.?":{}|<>]`, password); !matched {
		erros = append(erros, "password must contain at least one special character")
	}

	return len(erros) == 0, strings.Join(erros, "\n")
}

func MakeHash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidateHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
