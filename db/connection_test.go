package db

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	c := ConnectDB()
	fmt.Println(c)
}
