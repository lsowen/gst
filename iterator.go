//GstIterator — Object to retrieve multiple elements in a threadsafe way.
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

void gst_iterator_foreach_handler_dispatch(const GValue *item, gpointer user_data);

*/
import "C"

import (
	"github.com/lidouf/glib"
)

type IteratorResult C.GstIteratorResult

const (
	ITERATOR_DONE   = IteratorResult(C.GST_ITERATOR_DONE)
	ITERATOR_OK     = IteratorResult(C.GST_ITERATOR_OK)
	ITERATOR_RESYNC = IteratorResult(C.GST_ITERATOR_RESYNC)
	ITERATOR_ERROR  = IteratorResult(C.GST_ITERATOR_ERROR)
)

type IteratorItem C.GstIteratorItem

const (
	ITERATOR_ITEM_SKIP = IteratorItem(C.GST_ITERATOR_ITEM_SKIP)
	ITERATOR_ITEM_PASS = IteratorItem(C.GST_ITERATOR_ITEM_PASS)
	ITERATOR_ITEM_END  = IteratorItem(C.GST_ITERATOR_ITEM_END)
)

type Iterator C.GstIterator

func (i *Iterator) Free() {
	C.gst_iterator_free(i)
}

type IteratorCallbackFunction func(*glib.Value) bool

func (i *Iterator) forEach(callback IteratorCallbackFunction) IteratorResult {

	item := new(glib.Value)

	for {
		result := IteratorResult(C.gst_iterator_next(i, v2g(item)))
		switch result {
		case ITERATOR_OK:
			if !callback(item) {
				return result
			}
			break
		case ITERATOR_RESYNC:
			C.gst_iterator_resync(i)
			break
		case ITERATOR_ERROR:
		case ITERATOR_DONE:
			return result
		}
	}
}
