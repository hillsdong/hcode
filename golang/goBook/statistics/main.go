package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", homepage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("server start error", err)
	}
}

func homepage(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(write, pageTop, form)
	if err != nil {
		fmt.Fprintf(write, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			s := getStatis(numbers)
			fmt.Fprint(write, formatStatus(s))
		} else {
			fmt.Fprintf(write, anError, message)
		}
	}
	fmt.Fprint(write, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err == nil {
				numbers = append(numbers, x)
			} else {
				return numbers, "numbers is not valid", false
			}

		}
	}
	if len(numbers) == 0 {
		return numbers, "numbers is empty", false
	}
	return numbers, "", true
}

func formatStatus(s statis) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%v</td></tr>
<tr><td>StdDev</td><td>%f</td></tr>
</table>`, s.numbers, len(s.numbers), s.mean, s.median, s.mode, s.stddev)
}

type statis struct {
	numbers []float64
	mean    float64
	median  float64
	mode    []float64
	stddev  float64
}

func getStatis(numbers []float64) statis {
	Statis := statis{numbers: numbers}
	Statis.mean = sum(numbers) / float64(len(numbers))
	Statis.median = median(numbers)
	Statis.mode = mode(numbers)
	Statis.stddev = stddev(numbers, Statis.mean)
	return Statis
}

func sum(numbers []float64) (s float64) {
	for _, v := range numbers {
		s += v
	}
	return s
}

func median(numbers []float64) (m float64) {
	sort.Float64s(numbers)
	m = numbers[len(numbers)/2]
	if len(numbers)%2 == 0 {
		m = (numbers[len(numbers)/2-1] + m) / 2
	}
	return m
}

func mode(numbers []float64) (m []float64) {
	sort.Float64s(numbers)
	var numpre float64
	var count, maxcount int
	for _, num := range numbers {
		if num != numpre && count == maxcount {
			m = append(m, numpre)
		}

		if num != numpre {
			count = 1
		} else {
			count++
		}
		if count > maxcount {
			maxcount = count
			m = m[:0]
		}

		//fmt.Println(num, numpre, count, maxcount, m)
		numpre = num
	}
	if count == maxcount {
		m = append(m, numpre)
	}
	if len(numbers) == len(m) {
		m = m[:0]
	}
	return m
}

func mode2(numbers []float64) (modes []float64) {
	frequencies := make(map[float64]int, len(numbers))
	highestFrequency := 0
	for _, x := range numbers {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(frequencies) {
		modes = modes[:0] // Or: modes = []float64{}
	}
	sort.Float64s(modes)
	return modes
}

func stddev(numbers []float64, mean float64) (m float64) {
	var total float64
	for _, num := range numbers {
		total += math.Pow(num-mean, 2)
	}
	return math.Sqrt(total / float64(len(numbers)-1))
}

func quadratic(af, bf, cf float64) (complex128, complex128) {
	a, b, c := complex(af, 0), complex(bf, 0), complex(cf, 0)

	base := cmplx.Sqrt(b*b - 4*a*c)
	return (-b + base) / (2 * a), (-b - base) / (2 * a)
}
