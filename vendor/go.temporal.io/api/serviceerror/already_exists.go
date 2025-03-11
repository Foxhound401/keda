// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serviceerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// AlreadyExists represents general AlreadyExists gRPC error.
	AlreadyExists struct {
		Message string
		st      *status.Status
	}
)

// NewAlreadyExist returns new AlreadyExists error.
func NewAlreadyExist(message string) error {
	return &AlreadyExists{
		Message: message,
	}
}

// Error returns string message.
func (e *AlreadyExists) Error() string {
	return e.Message
}

func (e *AlreadyExists) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	return status.New(codes.AlreadyExists, e.Message)
}

func newAlreadyExists(st *status.Status) error {
	return &AlreadyExists{
		Message: st.Message(),
		st:      st,
	}
}
