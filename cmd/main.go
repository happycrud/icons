package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"getsvg/material"
)

func main() {
	material.AddTask("40", "red", "m-1").Render(context.Background(), os.Stdout)
}

var shap = []string{"", "outlined", "round", "sharp", "twotone"}

func mainx() {
	// 遍历
	root := "../material-design-icons/src/"
	dirs, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, category := range dirs {
		if category.IsDir() {
			svgs, _ := os.ReadDir(root + category.Name())
			for _, svg := range svgs {
				f, _ := os.Create("./material/" + svg.Name() + ".templ")
				f.Write([]byte(`package material

					`))
				for _, t := range shap {
					p, err := os.ReadFile(path.Join(root, category.Name(), svg.Name(), "materialicons"+t, "24px.svg"))
					x := bytes.Replace(p, []byte(`height="24"`), []byte(`height={ size } width={ size }`), 1)
					x = bytes.Replace(x, []byte(`width="24"`), []byte(``), 1)
					n := GoCamelCase(svg.Name())
					if isASCIIDigit([]byte(n)[0]) {
						n = "A" + n
					}
					if strings.HasSuffix(n, "Round") || strings.HasSuffix(n, "Outlined") {
						n = n + "2"
					}
					if err == nil {
						c := fmt.Sprintf(`
templ %s(size string,classes ...string){
	%s
}
						`, n+strings.Title(t), string(x))
						f.Write([]byte(c))
					}
				}
				f.Close()
			}
		}
	}
}

func GoCamelCase(s string) string {
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '.' && i+1 < len(s) && isASCIILower(s[i+1]):
			// Skip over '.' in ".{{lowercase}}".
		case c == '.':
			b = append(b, '_') // convert '.' to '_'
		case c == '_' && (i == 0 || s[i-1] == '.'):
			// Convert initial '_' to ensure we start with a capital letter.
			// Do the same for '_' after '.' to match historic behavior.
			b = append(b, 'X') // convert '_' to 'X'
		case c == '_' && i+1 < len(s) && isASCIILower(s[i+1]):
			// Skip over '_' in "_{{lowercase}}".
		case isASCIIDigit(c):
			b = append(b, c)
		default:
			// Assume we have a letter now - if not, it's a bogus identifier.
			// The next word is a sequence of characters that must start upper case.
			if isASCIILower(c) {
				c -= 'a' - 'A' // convert lowercase to uppercase
			}
			b = append(b, c)

			// Accept lower case sequence that follows.
			for ; i+1 < len(s) && isASCIILower(s[i+1]); i++ {
				b = append(b, s[i+1])
			}
		}
	}
	return string(b)
}

// JSONCamelCase converts a snake_case identifier to a camelCase identifier,
// according to the protobuf JSON specification.
func JSONCamelCase(s string) string {
	var b []byte
	var wasUnderscore bool
	for i := 0; i < len(s); i++ { // proto identifiers are always ASCII
		c := s[i]
		if c != '_' {
			if wasUnderscore && isASCIILower(c) {
				c -= 'a' - 'A' // convert to uppercase
			}
			b = append(b, c)
		}
		wasUnderscore = c == '_'
	}
	return string(b)
}

// JSONSnakeCase converts a camelCase identifier to a snake_case identifier,
// according to the protobuf JSON specification.
func JSONSnakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ { // proto identifiers are always ASCII
		c := s[i]
		if isASCIIUpper(c) {
			b = append(b, '_')
			c += 'a' - 'A' // convert to lowercase
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
