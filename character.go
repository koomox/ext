package ext

var (
	characterTables = map[string]string{
		"numeric": "0123456789",
		"simple":  "abcdefghijklmnopqrstuvwxyz",
		"special": "!#$%&()*+,-_./:;=?@[]^{}~|",
	}
)

const (
	alphabetA byte = byte('A')
	alphabetZ byte = byte('Z')
	alphabeta byte = byte('a')
	alphabetz byte = byte('z')
	numeric0 byte = byte('0')
	numeric9 byte = byte('9')
)

func IsUpper(ch byte) bool {
	if ch >= alphabetA && ch <= alphabetZ {
		return true
	}

	return false
}

func IsLower(ch byte) bool {
	if ch >= alphabeta && ch <= alphabetz {
		return true
	}

	return false
}

func IsNumber(ch byte) bool {
	if ch >= numeric0 && ch <= numeric9 {
		return true
	}

	return false
}

func IsSymbol(ch rune) bool {
	for _, c := range characterTables["special"] {
		if c == ch {
			return true
		}
	}

	return false
}

func IsAlphabet(ch byte) bool {
	if ch >= alphabetA && ch <= alphabetZ {
		return true
	}
	if ch >= alphabeta && ch <= alphabetz {
		return true
	}

	return false
}