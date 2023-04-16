package pkg

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if strconv.Itoa(randInt(65, 90)) != temp {
			temp = strconv.Itoa(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
