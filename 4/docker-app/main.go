package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var startEnd = r.FormValue("start")

		arr := strings.Split(startEnd, " ")

		fmt.Println(len(arr))

		if len(arr) != 2 {
			http.Error(w, "Input does not match", http.StatusBadRequest)
			return
		}

		start, _ := strconv.ParseInt(arr[0], 10, 64)
		end, _ := strconv.ParseInt(arr[1], 10, 64)
		count := 0

		for i := start; i < end; i++ {
			x := strconv.FormatInt(i, 10)
			lastIndex := len(strconv.FormatInt(i, 10)) - 1
			isPalindrom := true

			for j := 0; j < len(strconv.FormatInt(i, 10)); j++ {
				if string(x[j]) != string(x[lastIndex]) {
					isPalindrom = false
					break
				}

				lastIndex--
			}

			if isPalindrom {
				count++
			}
		}

		countString := strconv.Itoa(count)
		w.Write([]byte(countString))
		return
	}

	http.Error(w, "wrong request, please use method post", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/count", count)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
