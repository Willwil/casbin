// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"log"
	"testing"
)

func testFixAttribute(t *testing.T, s string, res string) {
	myRes := FixAttribute(s)
	log.Printf("%s: %s", s, myRes)

	if myRes != res {
		t.Errorf("%s: %s, supposed to be %s", s, myRes, res)
	}
}

func TestFixAttribute(t *testing.T) {
	testFixAttribute(t, "r.sub.domain", "subAttr(r.sub, \"domain\")")
	testFixAttribute(t, "r.obj.role", "objAttr(r.obj, \"role\")")
	testFixAttribute(t, "r.act.url", "actAttr(r.act, \"url\")")
	testFixAttribute(t, "p.sub.domain", "subAttr(p.sub, \"domain\")")
	testFixAttribute(t, "p.obj.role", "objAttr(p.obj, \"role\")")
	testFixAttribute(t, "p.act.url", "actAttr(p.act, \"url\")")
}

func testArrayEquals(t *testing.T, a []string, b []string, res bool) {
	myRes := ArrayEquals(a, b)
	log.Printf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestArrayEquals(t *testing.T) {
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "b", "c"}, true)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "b"}, false)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "c", "b"}, false)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{}, false)
}

func testArray2DEquals(t *testing.T, a [][]string, b [][]string, res bool) {
	myRes := Array2DEquals(a, b)
	log.Printf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestArray2DEquals(t *testing.T) {
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, true)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}, {"1", "2"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"1", "2", "3"}, {"a", "b", "c"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{}, false)
}

func testSetEquals(t *testing.T, a []string, b []string, res bool) {
	myRes := SetEquals(a, b)
	log.Printf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestSetEquals(t *testing.T) {
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "b", "c"}, true)
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "b"}, false)
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "c", "b"}, true)
	testSetEquals(t, []string{"a", "b", "c"}, []string{}, false)
}
