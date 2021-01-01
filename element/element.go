package element

import (
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var allElements = make(map[string][]Element)

// Element struct to represent a chemical element.
type Element struct {
	Symbol  string
	Name    string
	AtomNum int
	AtomWt  float64
	Group   string
}

// Initialize a new Element.
func initElement(symbol, name, group string, atomNum int, atomWt float64) Element {
	return Element{symbol, name, atomNum, atomWt, group}
}

// Import elements from periodic-table.xlsx.
// Returns map of elements sorted alphabetically.
func ImportElements() map[string][]Element {
	f, err := excelize.OpenFile("periodic-table.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows := f.GetRows("PT")
	for _, row := range rows {
		var symbol, name, group string
		var atomNum int
		var atomWt float64
		for i, colCell := range row {
			switch i {
			case 0:
				atomNum, err = strconv.Atoi(colCell)
				if err != nil {
					log.Fatal(err)
				}
			case 1:
				symbol = colCell
			case 2:
				name = colCell
			case 3:
				atomWt, err = strconv.ParseFloat(colCell, 64)
				if err != nil {
					log.Fatal(err)
				}
			case 4:
				group = colCell
			}
		}
		elem := initElement(symbol, name, group, atomNum, atomWt)
		allElements = populateAllElements(elem)
	}
	return allElements
}

// Populate a map with all elements sorted alphabetically.
func populateAllElements(elem Element) map[string][]Element {
	key := elem.Symbol[:1]
	allElements[key] = append(allElements[key], elem)
	return allElements
}

// Find a spelling for a word.
func Spell(input string, spelling []Element) []Element {
	if len(input) > 0 {
		start := input[:1]
		elements := allElements[strings.ToUpper(start)]
		for i, elem := range elements {
			if strings.HasPrefix(strings.ToLower(input), strings.ToLower(elem.Symbol)) {
				spelling = append(spelling[:], elem)
				spelling = Spell(input[len(elem.Symbol):], spelling[:])
				break
			} else if i == len(elements)-1 {
				spelling = nil
			}
		}
	}
	return spelling
}
