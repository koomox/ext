package ext

var (
	characterTables = map[string]string{
		"numeric": "0123456789",
		"simple":  "abcdefghijklmnopqrstuvwxyz",
		"special": "!#$%&()*+,-_./:;=?@[]^{}~|",
	}
)

func IsUpper(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}

	return false
}

func IsLower(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}

	return false
}

func IsNumber(ch interface{}) bool {
	switch ch.(type) {
	case rune:
		if ch.(rune) >=  && ch.(rune) <= '9' {
			return true
		}
	case byte:
		if rune(ch.(byte)) >= '0' && rune(ch.(byte)) <= '9' {
			return true
		}
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

func IsAlphabet(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= 'a' && ch <= 'z' {
		return true
	}

	return false
}