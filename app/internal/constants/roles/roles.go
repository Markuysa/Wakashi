package roles

import "strings"

// Below are the constants that
// are necessary to work correctly
// with adding new users and granting permissions
const (
	Administrator = iota + 1
	Shogun
	Daimyo
	Samurai
	Collector
	Card
)

// GetRoleID converts role string into internal integer roleID
func GetRoleID(role string) int {
	role = strings.ToLower(role)
	switch role {
	case "administrator":
		return Administrator
	case "shogun":
		return Shogun
	case "daimyo":
		return Daimyo
	case "samurai":
		return Samurai
	case "collector":
		return Collector
	case "card":
		return Card
	}
	return -1
}

// GetRoleString converts internal roleID into string role to
// show it to user
func GetRoleString(role int) string {
	switch role {
	case 1:
		return "administrator"
	case 2:
		return "shogun"
	case 3:
		return "daimyo"
	case 4:
		return "samurai"
	case 5:
		return "collector"
	case 6:
		return "card"
	}
	return ""
}
