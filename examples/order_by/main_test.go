// SPDX-FileCopyrightText: 2021 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// +build integration

package main

import "log"

func ExampleDoMain() {
	if err := DoMain(); err != nil {
		log.Printf("Failed to execute example: %v", err)
	}
	// Output:
	//
	// Opening database
	// Creating table 'order_by'
	// Query without cursor
	// Querying values from table without ordering
	// | a          | b                              |
	// | 1          | one                            |
	// | 0          | zero                           |
	// | 4          | four                           |
	// | 3          | three                          |
	// Querying values from table with ordering
	// | a          | b                              |
	// | 0          | zero                           |
	// | 1          | one                            |
	// | 3          | three                          |
	// | 4          | four                           |
	// Query with cursor
	// Querying values from table without ordering
	// | a          | b                              |
	// | 1          | one                            |
	// | 0          | zero                           |
	// | 4          | four                           |
	// | 3          | three                          |
	// Querying values from table with ordering
	// | a          | b                              |
	// | 0          | zero                           |
	// | 1          | one                            |
	// | 3          | three                          |
	// | 4          | four                           |
}
