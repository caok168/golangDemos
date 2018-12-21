package main

import (
	"fmt"
	"gopkg.in/mikespook/gorbac.v1"
)

func main() {
	permissions := []string{"admin", "guest", "normal"}

	rbac := gorbac.New()
	rbac.Add("rico", permissions[1:], []string{})

	IsGranted := rbac.IsGranted("rico", "guest", nil)

	fmt.Println("grant is ", IsGranted)
	fmt.Println("dump is ", rbac.Dump())

	fmt.Println("permissions are ", rbac.Get("rico").Permissions())
}
