package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

GstPadProbeReturn gst_probe_callback_handler_dispatch(GstPad *pad, GstPadProbeInfo *info, gpointer user_data);

static inline
GstPadProbeType CALL_MACRO_GST_PAD_PROBE_INFO_TYPE(GstPadProbeInfo *o) {
	return GST_PAD_PROBE_INFO_TYPE(o);
}

static inline
GstFlowReturn CALL_MACRO_GST_PAD_PROBE_INFO_FLOW_RETURN(GstPadProbeInfo *o) {
	return GST_PAD_PROBE_INFO_FLOW_RETURN(o);
}

*/
import "C"

import (
	"fmt"
	"github.com/lidouf/glib"
	"sync"
	"unsafe"
)

type PadProbeType C.GstPadProbeType

const (
	PAD_PROBE_TYPE_INVALID = PadProbeType(C.GST_PAD_PROBE_TYPE_INVALID)
	/* flags to control blocking */
	PAD_PROBE_TYPE_IDLE  = PadProbeType(C.GST_PAD_PROBE_TYPE_IDLE)
	PAD_PROBE_TYPE_BLOCK = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK)
	/* flags to select datatypes */
	PAD_PROBE_TYPE_BUFFER           = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER)
	PAD_PROBE_TYPE_BUFFER_LIST      = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER_LIST)
	PAD_PROBE_TYPE_EVENT_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM)
	PAD_PROBE_TYPE_EVENT_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_FLUSH      = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_FLUSH)
	PAD_PROBE_TYPE_QUERY_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_DOWNSTREAM)
	PAD_PROBE_TYPE_QUERY_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_UPSTREAM)
	/* flags to select scheduling mode */
	PAD_PROBE_TYPE_PUSH = PadProbeType(C.GST_PAD_PROBE_TYPE_PUSH)
	PAD_PROBE_TYPE_PULL = PadProbeType(C.GST_PAD_PROBE_TYPE_PULL)
	/* flag combinations */
	PAD_PROBE_TYPE_BLOCKING         = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCKING)
	PAD_PROBE_TYPE_DATA_DOWNSTREAM  = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_DOWNSTREAM)
	PAD_PROBE_TYPE_DATA_UPSTREAM    = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_UPSTREAM)
	PAD_PROBE_TYPE_DATA_BOTH        = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_BOTH)
	PAD_PROBE_TYPE_BLOCK_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_DOWNSTREAM)
	PAD_PROBE_TYPE_BLOCK_UPSTREAM   = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_BOTH       = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_BOTH)
	PAD_PROBE_TYPE_QUERY_BOTH       = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_BOTH)
	PAD_PROBE_TYPE_ALL_BOTH         = PadProbeType(C.GST_PAD_PROBE_TYPE_ALL_BOTH)
	PAD_PROBE_TYPE_SCHEDULING       = PadProbeType(C.GST_PAD_PROBE_TYPE_SCHEDULING)
)

type PadProbeReturn C.GstPadProbeReturn

const (
	GST_PAD_PROBE_DROP    = PadProbeReturn(C.GST_PAD_PROBE_DROP)
	GST_PAD_PROBE_OK      = PadProbeReturn(C.GST_PAD_PROBE_OK)
	GST_PAD_PROBE_REMOVE  = PadProbeReturn(C.GST_PAD_PROBE_REMOVE)
	GST_PAD_PROBE_PASS    = PadProbeReturn(C.GST_PAD_PROBE_PASS)
	GST_PAD_PROBE_HANDLED = PadProbeReturn(C.GST_PAD_PROBE_HANDLED)
)

type PadProbeInfo C.GstPadProbeInfo

func (p *PadProbeInfo) g() *C.GstPadProbeInfo {
	return (*C.GstPadProbeInfo)(p)
}

func (p *PadProbeInfo) AsPadProbeInfo() *PadProbeInfo {
	return p
}

func (p *PadProbeInfo) Type() PadProbeType {
	return PadProbeType(C.CALL_MACRO_GST_PAD_PROBE_INFO_TYPE(p.g()))
}

func (p *PadProbeInfo) Id() uint64 {
	return uint64(p.g().id)
}

func (p *PadProbeInfo) Data() unsafe.Pointer {
	return unsafe.Pointer(p.g().data)
}

//Returns
//The GstBuffer from the probe.
func (p *PadProbeInfo) Buffer() *Buffer {
	b := p.Data()
	if b == nil {
		return nil
	}
	r := new(Buffer)
	r.SetPtr(glib.Pointer(b))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstBufferList from the probe.
func (p *PadProbeInfo) GetBufferList() *BufferList {
	c := C.gst_pad_probe_info_get_buffer_list(p.g())
	if c == nil {
		return nil
	}
	r := new(BufferList)
	r.SetPtr(glib.Pointer(c))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstEvent from the probe.
func (p *PadProbeInfo) GetEvent() *Event {
	c := C.gst_pad_probe_info_get_event(p.g())
	if c == nil {
		return nil
	}
	r := new(Event)
	r.SetPtr(glib.Pointer(c))
	return r
}

//Parameters
//info
//a GstPadProbeInfo
//Returns
//The GstQuery from the probe.
func (p *PadProbeInfo) GetQuery() *Query {
	c := C.gst_pad_probe_info_get_query(p.g())
	if c == nil {
		return nil
	}
	r := new(Query)
	r.SetPtr(glib.Pointer(c))
	return r
}

func (p *PadProbeInfo) Offset() uint64 {
	return uint64(p.g().offset)
}

func (p *PadProbeInfo) Size() uint {
	return uint(p.g().size)
}

func (p *PadProbeInfo) FlowRet() FlowReturn {
	return FlowReturn(C.CALL_MACRO_GST_PAD_PROBE_INFO_FLOW_RETURN(p.g()))
}

//Remove the probe with id from pad .
//MT safe.
//Parameters
//pad
//the GstPad with the probe
//id
//the probe id to remove
func (p *Pad) RemoveProbe(event_type EventType, id uint64) {
	C.gst_pad_remove_probe(p.g(), C.gulong(id))
}

type PadProbeCallback func(*Pad, *PadProbeInfo) PadProbeReturn

type probeCallbackId uint64

type probeCallbackCounter struct {
	sync.Mutex
	current probeCallbackId
}

func (c *probeCallbackCounter) Next() probeCallbackId {
	c.Lock()
	defer c.Unlock()
	id := c.current
	c.current += 1
	return id
}

var currentCallbackId = &probeCallbackCounter{}

var probe_handlers = make(map[probeCallbackId]PadProbeCallback)

//export gst_probe_callback_handler
func gst_probe_callback_handler(cpad *C.GstPad, cinfo *C.GstPadProbeInfo, user_data C.gpointer) PadProbeReturn {
	callback_id := *(*probeCallbackId)(user_data)

	if callback, ok := probe_handlers[callback_id]; ok {
		pad := new(Pad)
		pad.SetPtr(glib.Pointer(cpad))

		info := (*PadProbeInfo)(cinfo)
		result := callback(pad, info)
		if result == GST_PAD_PROBE_REMOVE {
			delete(probe_handlers, callback_id)
		}
		return result
	}

	fmt.Errorf("No handler found for callback %d\n", callback_id)
	return GST_PAD_PROBE_REMOVE
}

func (p *Pad) AddProbe(mask PadProbeType, callback PadProbeCallback) {

	callback_id := currentCallbackId.Next()
	probe_handlers[callback_id] = callback
	C.gst_pad_add_probe(p.g(), C.GstPadProbeType(mask), (C.GstPadProbeCallback)(C.gst_probe_callback_handler_dispatch), C.gpointer(unsafe.Pointer(&callback_id)), nil)

}
