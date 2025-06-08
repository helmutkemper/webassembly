package utils

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"syscall/js"
	"time"
)

// VerifyName validates and processes the provided name.
func VerifyName(name string) (string, error) {
	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")

	// Replace special characters with underscores
	regex := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	name = regex.ReplaceAllString(name, "_")

	// Check if the name is empty
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	// Check if the name is too long
	if len(name) > 50 {
		return "", errors.New("name is too long")
	}

	// Check if the name is too short
	if len(name) < 3 {
		return "", errors.New("name is too short")
	}

	// Check if the name is not unique
	if !IsNameUnique(name) {
		return "", errors.New("name is not unique")
	}

	// Return the valid name
	return name, nil
}

// VerifyUniqueName validates, processes, and checks the uniqueness of the provided name.
func VerifyUniqueName(name string) (string, error) {
	name, err := VerifyName(name)
	if err != nil {
		return "", err
	}

	// Check if the name is not unique
	if !IsNameUnique(name) {
		return "", errors.New("name is not unique")
	}

	return name, nil
}

// VerifyUniqueId validates and processes the provided ID.
func VerifyUniqueId(ID string) (string, error) {
	ID, err := VerifyName(ID)
	if err != nil {
		return "", err
	}

	// Check if the ID is not unique
	if !IsIdUnique(ID) {
		return "", errors.New("ID is not unique")
	}

	return ID, nil
}

// IsNameUnique checks if a name is unique.
func IsNameUnique(name string) (unique bool) {
	names := GetAllNames()
	for _, n := range names {
		if n == name {
			return false
		}
	}
	return true
}

// IsIdUnique checks if an ID is unique.
func IsIdUnique(ID string) (unique bool) {
	ids := GetAllIds()
	for _, id := range ids {
		if id == ID {
			return false
		}
	}
	return true
}

// GetAllNames retrieves all names from a data source (stubbed here).
func GetAllNames() (names []string) {
	names = make([]string, 0)

	elements := js.Global().Get("document").Call("querySelectorAll", "*")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i).Call("getAttribute", "name")
		if element.IsNull() {
			continue
		}
		names = append(names, element.String())
	}

	return
}

// GetAllIds retrieves all IDs from a data source (stubbed here).
func GetAllIds() (ids []string) {
	ids = make([]string, 0)

	elements := js.Global().Get("document").Call("querySelectorAll", "*")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i).Call("getAttribute", "id")
		if element.IsNull() {
			continue
		}
		ids = append(ids, element.String())
	}

	return
}

// GetRandomNumericId generates a random numeric ID.
func GetRandomNumericId() string {
	rand.Seed(time.Now().UnixNano())
	return randomString(13, "0123456789")
}

// GetRandomId generates a random ID and ensures uniqueness.
func GetRandomId() string {
	for {
		ID := GetRandomString(10)
		if IsIdUnique(ID) {
			return ID
		}
	}
}

// GetRandomString generates a random string of the given length.
func GetRandomString(length int) string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	return randomString(length, characters)
}

// GetRandomName generates a random name.
func GetRandomName() string {
	return GetRandomString(10)
}

// GetRandomColor generates a random color in a hexadecimal format.
func GetRandomColor() string {
	rand.Seed(time.Now().UnixNano())
	return "#" + randomString(6, "0123456789ABCDEF")
}

// GetRandomNumber generates a random number between min and max (inclusive).
func GetRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GetRandomBoolean generates a random boolean value.
func GetRandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}

// randomString generates a random string of length n from characters.
func randomString(n int, characters string) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n)
	charSetLength := len(characters)
	for i := range result {
		result[i] = characters[rand.Intn(charSetLength)]
	}
	return string(result)
}
