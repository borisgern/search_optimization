package main

import (
	"bytes"
	"fmt"
	jlexer "github.com/mailru/easyjson/jlexer"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var _ *jlexer.Lexer

var seenBrowsers = make(map[string]bool)
var uniqueBrowsers int

func easyjsonE6b4cdeDecodeCourseraOrgHw3Js(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "company":
			out.Company = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE6b4cdeDecodeCourseraOrgHw3Js(&r, v)
	return r.Error()
}

type User struct {
	Browsers []string `json:"browsers"`
	Company  string   `json:"company"`
	Country  string   `json:"country"`
	Email    string   `json:"email"`
	Job      string   `json:"job"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}

func FastSearch(out io.Writer) {
	const filePath string = "./data/users.txt"
	var foundUsers string
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(fileContents, []byte("\n"))
	for i, line := range lines {
		var user User
		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false
		withAndroid := false
		withMSIE := false
		if len(user.Browsers) == 0 {
			continue
		}
		for _, browserRaw := range user.Browsers {
			if withAndroid = strings.Contains(browserRaw, "Android"); withAndroid && !isAndroid {
				isAndroid = true
			}
			if withMSIE = strings.Contains(browserRaw, "MSIE"); withMSIE && !isMSIE {
				isMSIE = true
			}

			if withAndroid || withMSIE {
				if _, ok := seenBrowsers[browserRaw]; !ok {
					seenBrowsers[browserRaw] = true
				}
			}
		}
		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		str := strconv.Itoa(i)
		foundUsers += "[" + str + "] " + user.Name + " <" + email + ">\n"

	}
	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

func main() {

}
