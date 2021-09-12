package input

import (
	"path/filepath"
	"strings"
	"unicode"
)

func PathToSnake(path string) string {
	ext := filepath.Ext(path)
	name := strings.TrimSuffix(path, ext)
	snake := StringToSnake(filepath.Base(name))
	full := filepath.Join(filepath.Dir(path), snake+ext)

	return full
}

func StringToSnake(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}
