/*
 * makefile_test.go
 *
 * Copyright (c) 2016-2018 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

import (
	"strings"
	"testing"
)

// TestMakefile tests generated Makefile contains a name given as a parameter.
func TestMakefile(t *testing.T) {

	param := Makefile{
		Dest: "test",
	}

	data, err := param.Generate()
	if err != nil {
		t.Fatal("Generate returned an error:", err)
	}

	if !strings.Contains(string(data), "-d=test") {
		t.Errorf("Generated Makefile was wrong.\n%s", string(data))
	}

}
