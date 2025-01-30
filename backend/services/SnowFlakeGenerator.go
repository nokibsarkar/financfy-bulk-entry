package services

import (
	"fmt"
	"strconv"

	"github.com/godruoyi/go-snowflake"
)

func GenerateSnowFlake() string {
	// Generate a snowflake id
	id := snowflake.ID()
	fmt.Println(strconv.FormatUint(id, 32))
	// Convert the id to string
	return strconv.FormatUint(id, 32)
}
