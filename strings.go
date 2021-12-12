package ext

import (
	"reflect"
	"regexp"
	"strconv"
)

var (
	uuidExpCompile = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
)

func IsValidUUID(uuid string) bool {
	return uuidExpCompile.MatchString(uuid)
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	return i == nil || reflect.ValueOf(i).IsNil()
}

func Atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func DomainSuffix(s string) string {
	i := len(s)
	count := 0
	end := 0
	for ; i != 0; i-- {
		if s[i-1] == '.' {
			if (i - 1) == 0 {
				return s
			}
			count += 1
			switch count {
			case 1:
				end = i - 1
			case 2:
				switch s[i:end] {
				case "com", "co", "gov", "edu", "org", "net":
					end = i - 1
				default:
					return s[i:]
				}
			default:
				return s[i:]
			}
		}
	}
	if count == 0 {
		return ""
	}

	return s
}

func DomainKeyword(s string) string {
	i := len(s)
	count := 0
	end := i
	for ; i != 0; i-- {
		if s[i-1] == '.' {
			if (i - 1) == 0 {
				return s[:end]
			}
			count += 1
			switch count {
			case 1:
				end = i - 1
			case 2:
				switch s[i:end] {
				case "com", "co", "gov", "edu", "org", "net":
					end = i - 1
				default:
					return s[i:end]
				}
			default:
				return s[i:end]
			}
		}
	}
	if count == 0 {
		return ""
	}

	return s[i:end]
}

func DomainCountry(s string) string {
	i := len(s)
	for ; i != 0; i-- {
		if s[i-1] == '.' {
			break
		}
	}
	if i == 0 {
		return ""
	}
	return s[i:]
}

func Filextension(s string) string {
	i := len(s)
	for ; i != 0; i-- {
		if s[i-1] == '.' {
			break
		}
	}
	if i == 0 {
		return ""
	}
	return s[i:]
}

func Filename(s string) string {
	i := len(s)
	d := i
	for ; i != 0; i-- {
		if s[i-1] == '?' {
			d = i - 1
			break
		}
	}
	if i == 0 {
		i = d
	}
	for ; i != 0; i-- {
		if s[i-1] == '/' {
			break
		}
	}
	return s[i:d]
}