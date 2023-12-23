/*
 * Copyright © 2023 pkeorley
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package main

type CredentialsType map[string]string

func Credentials() CredentialsType {
	return CredentialsType{
		"licence": "MIT",
		"author":  "pkeorley",
	}
}
