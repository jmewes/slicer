// Copyright (C) 2026 Jan Mewes
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func ToKebabCase(s string) string {
	input := []rune(s)
	runes := make([]rune, 0, len(input))
	for i, r := range input {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				// Insert a dash to separate words, but keep consecutive
				// uppercase letters (acronyms) together. A boundary is
				// introduced either when transitioning from a lowercase
				// letter or digit into an uppercase letter, or at the end
				// of an acronym (an uppercase letter followed by a
				// lowercase letter is the start of a new word).
				if i > 0 {
					prev := input[i-1]
					if unicode.IsLower(prev) || unicode.IsDigit(prev) {
						runes = append(runes, '-')
					} else if unicode.IsUpper(prev) && i+1 < len(input) && unicode.IsLower(input[i+1]) {
						runes = append(runes, '-')
					}
				}
				runes = append(runes, unicode.ToLower(r))
			} else {
				runes = append(runes, r)
			}
		} else if unicode.IsDigit(r) {
			runes = append(runes, r)
		} else {
			runes = append(runes, '-')
		}
	}

	result := strings.Trim(string(runes), "-")
	result = regexp.MustCompile(`-+`).ReplaceAllString(result, "-")
	return result
}
