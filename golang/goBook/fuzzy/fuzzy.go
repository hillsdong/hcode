package main

import "fmt"

func main() {
	a, _ := New(0)   // Safe to ignore err value when using
	b, _ := New(.25) // known valid values; must check if using
	c, _ := New(.75) // variables though.
	d := c.Copy()
	if err := d.Set(1); err != nil {
		fmt.Println(err)
	}
	process(a, b, c, d)
	s := []*FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *FuzzyBool) {
	fmt.Println("Original:", a, b, c, d)
	fmt.Println("Not:     ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(),
		d.Not().Not())
	fmt.Print("0.And(.25)→", a.And(b), "• .25.And(.75)→", b.And(c),
		"• .75.And(1)→", c.And(d), " • .25.And(.75,1)→", b.And(c, d), "\n")
	fmt.Print("0.Or(.25)→", a.Or(b), "• .25.Or(.75)→", b.Or(c),
		"• .75.Or(1)→", c.Or(d), " • .25.Or(.75,1)→", b.Or(c, d), "\n")
	fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), c.Less(a))
	fmt.Println("Bool:    ", a.Bool(), b.Bool(), c.Bool(), d.Bool())
	fmt.Println("Float:   ", a.Float(), b.Float(), c.Float(), d.Float())
}

// FuzzyBool is a super bool struct
type FuzzyBool struct{ value float32 }

// New return a new FuzzyBool struct
func New(value interface{}) (*FuzzyBool, error) {
	val, err := float32ForValue(value)
	return &FuzzyBool{val}, err
}

func float32ForValue(value interface{}) (fuzzy float32, err error) {
	switch value := value.(type) {
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		err = fmt.Errorf("error %v is not a bool or a int or a float ", value)
	}
	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}
	return fuzzy, err
}

// Set can reset the value
func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = float32ForValue(value)
	return err
}

// Copy return a same new struct
func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

// String for fmt
func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", fuzzy.value*100)
}

// Not return a 1- value
func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

// And return a min value
func (fuzzy *FuzzyBool) And(first *FuzzyBool,
	rest ...*FuzzyBool) *FuzzyBool {
	min := fuzzy.value
	rest = append(rest, first)
	for _, fuz := range rest {
		if fuz.value < min {
			min = fuz.value
		}
	}
	return &FuzzyBool{min}
}

// Or return a max value
func (fuzzy *FuzzyBool) Or(first *FuzzyBool,
	rest ...*FuzzyBool) *FuzzyBool {
	max := fuzzy.value
	rest = append(rest, first)
	for _, fuz := range rest {
		if fuz.value > max {
			max = fuz.value
		}
	}
	return &FuzzyBool{max}
}

// Less can compare with other struct
func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

// Equal can compare with other struct
func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

// Bool return a bool value
func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value >= 0.5
}

// Float 返回一个float64的值
func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}
