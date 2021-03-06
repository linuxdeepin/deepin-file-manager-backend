/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package desktop

// #cgo pkg-config: gio-unix-2.0 gdk-3.0
// void run_in_terminal(char* dir, char* executable);
// void free(void*);
import "C"
import "unsafe"

func runInTerminal(dir string, executable string) {
	cDir := C.CString(dir)
	if dir == "" {
		C.free(unsafe.Pointer(cDir))
		cDir = C.CString("\x00")
	}
	defer C.free(unsafe.Pointer(cDir))

	cExecutable := C.CString(executable)
	if executable == "" {
		C.free(unsafe.Pointer(cExecutable))
		cExecutable = C.CString("\x00")
	}
	defer C.free(unsafe.Pointer(cExecutable))

	C.run_in_terminal(cDir, cExecutable)
}
