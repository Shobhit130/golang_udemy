// Package dog allows us to fully understand dogs
package dog

// Years converts human years to dog years
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

//go doc 
//go doc dog.Years
//go test
//go test -bench .
//godoc http=:8080
//go test -cover
//go test -coverprofile c.out
//go tool cover -html c.out -o coverage.html