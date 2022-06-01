// @author: JeffreyDin
// @license:
// @contact: dingjianfeng15@gmail.com
// @software:
// @file: say.go
// @ide: GoLand
// @date: 2022/6/1
// @desc:

package test2

import "fmt"

// SayHi say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
