package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
)

var pow = [1000]string{}


func main() {
	for i, v := range pow {
		// v = strings.Join(strings.Join(string(123), ","), string(rand.Intn(1000)))
		v = strconv.FormatInt(time.Now().UnixNano(), 10)
		var b strings.Builder
		b.WriteString(v)
		b.WriteString(",")
		b.WriteString(strconv.FormatInt(int64(rand.Intn(1000)), 10))
		fmt.Printf("%d = %s\n", i, b.String())
	}
}