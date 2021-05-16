package main

import (
	"fmt"
	"time"
)

func main() {
	// at := access_token.GetNewAccessToken()

	fmt.Println(time.Now())
	// 2021-05-15 01:06:23.560597015 +0530 IST m=+0.000028189

	Expires := time.Now().UTC().Add(24 * time.Hour).Unix()
	fmt.Println(Expires)
	// 1621107383

	now := time.Now().UTC()
	fmt.Println(now)
	// 2021-05-14 19:36:23.56065957 +0000 UTC

	expirationTime := time.Unix(Expires, 0)
	fmt.Println(expirationTime)
	// 2021-05-16 01:06:23 +0530 IST

	fmt.Println(now.After(expirationTime))
	// false
}
