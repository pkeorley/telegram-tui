/*
 * Copyright © 2023 pkeorley
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package main

import "fmt"

func main() {
	for k, v := range Credentials() {
		fmt.Printf("%s = %s\n", k, v)
	}
}
