package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head>
<title>Soundex</title><body><h3>Soundex</h3>
<p>Compute soundex codes for a list of names.</p>`
	form = `<form action="/" method="POST">
<label for="names">Names (comma or space-separated):</label><br />
<input type="text" name="names" size="30"><br />
<input type="submit" name="compute" value="Compute">
</form>`
	pageBottom = `</body></html>`
	error      = `<p class="error">%s</p>`
)

var digitForLetter = []rune{
	// A  B  C  D  E  F  G  H  I  J  K  L  M
	0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5,
	// N  O  P  Q  R  S  T  U  V  W  X  Y  Z
	5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}

var testCases map[string]string

func main() {
	var ok bool
	if testCases, ok = populateTestCases("soundex-test-data.txt"); !ok {
		log.Fatal("parse file err")
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", testPage)

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("server start err: ", err)
	}
}

func populateTestCases(filename string) (map[string]string, bool) {
	testCases := make(map[string]string)
	if raw, err := ioutil.ReadFile(filename); err == nil {
		for _, line := range strings.Split(string(raw), "\n") {
			if fields := strings.Fields(line); len(fields) == 2 {
				testCases[fields[1]] = fields[0]
			}
		}
		return testCases, true
	}
	return testCases, false
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, error, err)
	}
	fmt.Fprint(writer, pageTop, form)
	if names := processRequest(request); len(names) > 0 {
		soundexes := make([]string, 0, len(names))
		for _, name := range names {
			soundexes = append(soundexes, soundex(name))
		}
		fmt.Fprint(writer, formatResults(names, soundexes))
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (names []string) {
	if slice := request.FormValue("names"); len(slice) > 0 {
		slice = strings.Replace(slice, ",", " ", -1)
		names = strings.Fields(slice)
	}
	return names
}

// "c - 'A'" produces a 0-based index, so 'A' -> 0, 'B' -> 1, etc.
// "'0' + digitForLetter[index]" converts a one digit integer into the
// equivalent Unicode character, i.e., 0 -> "0", 1 -> "1", etc.
func soundex(name string) string {
	nam := []rune(strings.ToUpper(name))
	var sou []rune
	sou = append(sou, nam[0])
	for i := 1; i < len(nam); i++ {
		if j := nam[i] - 'A'; j >= 0 &&
			j < int32(len(digitForLetter)) &&
			digitForLetter[j] != 0 &&
			nam[i] != nam[i-1] {
			sou = append(sou, '0'+digitForLetter[j])
		}
	}
	for len(sou) < 4 {
		sou = append(sou, '0')
	}
	if len(sou) > 4 {
		sou = sou[:4]
	}
	return string(sou)
}

func formatResults(names, soundexes []string) string {
	text := `<table border="1"><tr><th>Name</th><th>Soundex</th></tr>`
	for i := range names {
		text += "<tr><td>" + html.EscapeString(names[i]) + "</td><td>" +
			html.EscapeString(soundexes[i]) + "</td></tr>"
	}
	return text + "</table>"
}

func testPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, `<html><head><title>Soundex Test</title>
<style>.fail{color:#F00;} .pass{color:#0F0;}</style></head><body>
<table border="1"><tr><th>Name</th><th>Soundex</th>
<th>Expected</th><th>Test</th></tr>`)
	var names []string
	for name := range testCases {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		actual := soundex(name)
		expected := testCases[name]
		test := `<span class="fail">FAIL</span>`
		if actual == expected {
			test = `<span class="pass">PASS</span>`
		}
		fmt.Fprintf(writer, `<tr><td>%s</td><td>%s</td><td>%s</td>
<td>%s</td></tr>`, name, actual, expected, test)
	}
	fmt.Fprint(writer, "</table></body></html>")
}
