package main

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func isDoublePrime(value int, numberOfDigits int) bool {
	origValue := value
	var isDPrime bool
	switch value {
	case 0, 1:
		// If it is 0 or 1. it is not a prime. assign false
		isDPrime = false
	default:
		// Else assign the initial value to true
		isDPrime = true
	}

	for value > 1 && isDPrime {
		// Used to extract digits from right
		// For e.g. 17.
		// iteration 1, 17
		// iteration 2, 1. Then exits loop
		// Since the loop will only be called if the prev value is true, no need to 'AND'
		isDPrime = isPrime(value)
		value = value / 10
	}
	for origValue > 10 && isDPrime {
		// Used to extract digits from left
		// For e.g. 17.
		// iteration 1, 7
		// iteration 2, exit
		numberOfDigits = numberOfDigits - 1
		origValue = origValue % int(math.Pow10(numberOfDigits))
		isDPrime = isPrime(origValue)
	}
	return isDPrime
}

func home(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	val, exists := pathParams["number"]
	var respStatus int
	var message string
	if exists {
		num, err := strconv.Atoi(val)
		if err != nil {
			respStatus = http.StatusBadRequest
			message = `{"message": "need a number"}`
		} else {
			respStatus = http.StatusOK
			message = strconv.FormatBool(isDoublePrime(num, len(val)))
		}
	} else {
		respStatus = http.StatusNotFound
		message = `{"message": "Couldn't locate the endpoint"}`
	}

	w.WriteHeader(respStatus)
	w.Write([]byte(message))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/isDoublePrime/{number}", home).Methods(http.MethodGet)
	const port int64 = 8010
	portStr := strconv.FormatInt(port, 10)
	log.Println("Server is listening on port " + portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, r))
}
