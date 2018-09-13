package tray

/*
#cgo LDFLAGS: -L.
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

extern int clickCallback(int x, int y, bool rightClick);

#include "lib/tray.h"

static int initialise_tray(void * iconData) {
	struct tray tray = {
		.icon = iconData
	};

	if (tray_init(&tray) < 0) {
		return 1;
	}
	while (tray_loop(1) == 0) {}
	return 0;
}
*/
import "C"
import (
	"unsafe"
)

type ClickableIcon struct {
	IconData     []byte
	ClickHandler func(x, y int, rightClick bool)
}

var tray ClickableIcon

func (t *ClickableIcon) Initialise() {
	tray.ClickHandler = t.ClickHandler
	go C.initialise_tray(unsafe.Pointer(&t.IconData[0]))
}

//export clickCallback
func clickCallback(x, y C.int, rightClick C.bool) C.int {
	tray.ClickHandler(int(x), int(y), bool(rightClick))
	return 1
}
