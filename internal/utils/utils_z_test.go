// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package utils_test

import (
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/test/gtest"
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_ReadCloser(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			r    []byte
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			r    []byte
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, true)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
	})
}

func Test_RemoveSymbols(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`), `abac101`)
	})
}

func Test_OriginValueAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := utils.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := utils.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := utils.OriginValueAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := utils.OriginValueAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}

func Test_OriginTypeAndKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := utils.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.String)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s = "s"
		out := utils.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.String)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := utils.OriginTypeAndKind(s)
		t.Assert(out.InputKind, reflect.Slice)
		t.Assert(out.OriginKind, reflect.Slice)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []int
		out := utils.OriginTypeAndKind(&s)
		t.Assert(out.InputKind, reflect.Ptr)
		t.Assert(out.OriginKind, reflect.Slice)
	})
}
