// Copyright 2021 the Dupi authors
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

//go:build windows
// +build windows

package lock

const (
	reserved = 0
	allBytes = ^uint32(0)
)

func (f *File) Lock() error {
	ol := new(syscall.Overlapped)
	h := syscall.Handle(f.handle.Fd())
	return windows.LockFileEx(h, reserved, allBytes, allBytes, ol)
}

func (f *File) Unlock() error {
	ol := new(syscall.Overlapped)
	h := syscall.Handle(f.handle.Fd())
	return windows.UnlockFileEx(h, reserved, allBytes, allBytes, ol)
}

func (f *File) LockShared() error {
	return f.Lock()
}
