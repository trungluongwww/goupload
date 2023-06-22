package prandom

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomCode(length int) string {
	rs := make([]rune, length)

	for i := range rs {
		rs[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rs)
}

func RandomNameFileFromExtension(ext string) string {
	return fmt.Sprintf("%s_%s.%s", time.Now().Format("2006_01_02_15_04_05"), randomCode(4), ext)
}
