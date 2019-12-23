package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile(filepath string) []string {
	f, err := os.Open(filepath)
	defer f.Close()

	var content []string

	CheckErr(err)
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		CheckErr(err)

		content = append(content, strings.Trim(string(line), " "))
	}

	return content

	// inputStrs := strings.Split(, ",")
	// res := []int{}
	// for _, i := range inputStrs {
	// 	smth, err := strconv.Atoi(i)
	// 	CheckErr(err)

	// 	res = append(res, smth)
	// }

	// return res
}
