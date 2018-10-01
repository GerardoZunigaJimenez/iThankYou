package service

import (
	"fmt"
	"testing"
)

func TestFetchUserByEmail(t *testing.T) {
	u := FetchUserByEmail("gzuniga@itexico.net")

	fmt.Println(u)
}
