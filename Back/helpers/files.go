package helpers

import (
	"context"
	"fmt"
	"os"
)

//FileExists checks if a file exists and is not a directory
func FileExists(ctx context.Context, filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
