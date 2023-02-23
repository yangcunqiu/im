package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumberToString(num int) string {
	s := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
