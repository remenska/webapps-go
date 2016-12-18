package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))
	// Find all matches and save to a slice, n less than 0 means return all matches,
	// indicates length of slice if it's greater than 0.
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll:", all)

	// Find index of first match, start and end position.
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex:", index)

	// Find index of all matches, the n does same job as above.
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")
	// Find first submatch and return array, the first element contains all elements,
	// the second element contains the result of first (), the third element contains
	// the result of second ().
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchIndex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println("SubmatchIndex:", submatchIndex)

	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println("Submatchall:", submatchall)

	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println("SubmatchAllIndex:", submatchallindex)

	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
