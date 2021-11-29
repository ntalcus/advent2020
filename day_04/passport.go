package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("file not read")
	}
	s := bufio.NewScanner(f)
	count := 0
	p := &passport{}
	for s.Scan() {
		t := s.Text()
		if len(t) == 0 {
			if p.IsValid() {
				count++
			}
			p = &passport{}
			continue
		}
		sections := strings.Split(t, " ")
		for _, section := range sections {
			field := strings.Split(section, ":")
			switch name := field[0]; name {
			case "byr":
				p.byr = field[1]
			case "iyr":
				p.iyr = field[1]
			case "eyr":
				p.eyr = field[1]
			case "hgt":
				p.hgt = field[1]
			case "hcl":
				p.hcl = field[1]
			case "ecl":
				p.ecl = field[1]
			case "pid":
				p.pid = field[1]
			case "cid":
				p.cid = field[1]
			default:
				continue
			}
		}
	}
	if p.IsValid() {
		count++
	} else {
		fmt.Printf("not valid: %+v \n", p)
	}
	fmt.Printf("num valid: %d\n", count)
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) IsValid() bool {
	return validYear(p.byr, 1920, 2002) &&
		validYear(p.iyr, 2010, 2020) &&
		validYear(p.eyr, 2020, 2030) &&
		validHeight(p.hgt) &&
		validHairColor(p.hcl) &&
		validEyeColor(p.ecl) &&
		validPID(p.pid)
	// cid != ""

}

func validYear(data string, lower int, upper int) bool {
	val, err := strconv.ParseInt(data, 0, 64)

	return err == nil && int(val) >= lower && int(val) <= upper
}

var lowerboundCM, upperboundCM int = 150, 193
var lowerboundIN, upperboundIN int = 59, 76

func validHeight(data string) bool {
	if len(data) == 0 {
		return false
	}
	valid := false
	if heightCM := strings.TrimSuffix(data, "cm"); len(heightCM) != len(data) {
		val, err := strconv.ParseInt(heightCM, 0, 64)
		valid = err == nil && int(val) >= lowerboundCM && int(val) <= upperboundCM
	} else if heightIN := strings.TrimSuffix(data, "in"); len(heightIN) != len(data) {
		val, err := strconv.ParseInt(heightIN, 0, 64)
		valid = err == nil && int(val) >= lowerboundIN && int(val) <= upperboundIN
	}
	return valid
}

func validHairColor(data string) bool {
	if len(data) != 7 || data[:1] != "#" {
		return false
	}
	for _, char := range data[1:] {
		if !validHexChar(char) {
			return false
		}
	}
	return true
}

func validHexChar(char rune) bool {
	return (char >= rune(48) && char <= rune(57)) ||
		(char >= rune(97) && char <= rune(122))
}

var colors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validEyeColor(data string) bool {
	valid := colors[data]
	return valid
}

func validPID(data string) bool {
	trimmed := strings.TrimLeftFunc(data, func(r rune) bool {
		return r == rune(48)
	})
	_, err := strconv.ParseInt(trimmed, 0, 64)
	return err == nil && len(data) == 9
}
