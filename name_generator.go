package go90s

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/theRemix/go90s/lib"
)

type rhash struct {
	Aidx int64 // if -1 then don't use
	Nidx int64
	Vidx int64 // if -1 then don't use
}

func GetRandomName() string {
	rand.Seed(time.Now().UnixNano())
	return GetName(strconv.Itoa(rand.Intn(999999999)))
}

// hash the input to get a nineties name
func GetName(input string) string {
	hash := getHash(input)
	words := []string{}

	if hash.Aidx >= 0 {
		rand.Seed(hash.Aidx)
		words = append(words, names.Adjectives[rand.Intn(len(names.Adjectives))])
	}
	if hash.Nidx >= 0 {
		rand.Seed(hash.Nidx)
		words = append(words, names.Names[rand.Intn(len(names.Names))])
	}
	if hash.Vidx >= 0 {
		rand.Seed(hash.Vidx)
		words = append(words, names.Verbs[rand.Intn(len(names.Verbs))])
	}

	return strings.Join(words, " ")
}

// input is 5-10 digits
// if the first number is 1, use verb
// if the second and third number is also 1, also use adjective
// split the digits, in 2 or 3 parts, concat them, then mod
func getHash(input string) *rhash {
	h := fnv.New32a()
	h.Write([]byte(input))
	digits := strconv.Itoa(int(h.Sum32()))

	splitDigits := []int{}
	aDigits := []int{-1}
	nDigits := []int{-1}
	vDigits := []int{-1}
	var splitidx int

	for _, i := range digits {
		j, err := strconv.Atoi(string(i))
		if err != nil {
			panic(err)
		}
		splitDigits = append(splitDigits, j)
	}

	if splitDigits[0] == 1 {

		if splitDigits[1] == 1 && splitDigits[2] == 1 {
			// adjective + name + verb
			splitidx = len(digits) / 3
			aDigits = splitDigits[0:splitidx]
			nDigits = splitDigits[splitidx : splitidx*2]
			vDigits = splitDigits[splitidx*2:]
		} else {
			// name + verb
			splitidx = len(digits) / 2
			nDigits = splitDigits[0:splitidx]
			vDigits = splitDigits[splitidx:]
		}
	} else {
		// adjective + name
		splitidx = len(digits) / 2
		aDigits = splitDigits[0:splitidx]
		nDigits = splitDigits[splitidx:]
	}

	return &rhash{
		Aidx: intSliceToInt(aDigits),
		Nidx: intSliceToInt(nDigits),
		Vidx: intSliceToInt(vDigits),
	}
}

func intSliceToInt(digits []int) int64 {
	outs := []string{}
	for i := range digits {
		number := digits[i]
		text := strconv.Itoa(number)
		outs = append(outs, text)
	}

	joinedInt, err := strconv.Atoi(strings.Join(outs, ""))
	if err != nil {
		fmt.Println(err)
		joinedInt = -1
	}

	return int64(joinedInt)
}
