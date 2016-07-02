package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
)

type PadLinkReturn C.GstPadLinkReturn

const (
	PAD_LINK_OK              = PadLinkReturn(C.GST_PAD_LINK_OK)
	PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
	PAD_LINK_WAS_LINKED      = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
	PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
	PAD_LINK_NOFORMAT        = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
	PAD_LINK_NOSCHED         = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
	PAD_LINK_REFUSED         = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

func (p PadLinkReturn) String() string {
	switch p {
	case PAD_LINK_OK:
		return "PAD_LINK_OK"
	case PAD_LINK_WRONG_HIERARCHY:
		return "PAD_LINK_WRONG_HIERARCHY"
	case PAD_LINK_WAS_LINKED:
		return "PAD_LINK_WAS_LINKED"
	case PAD_LINK_WRONG_DIRECTION:
		return "PAD_LINK_WRONG_DIRECTION"
	case PAD_LINK_NOFORMAT:
		return "PAD_LINK_NOFORMAT"
	case PAD_LINK_NOSCHED:
		return "PAD_LINK_NOSCHED"
	case PAD_LINK_REFUSED:
		return "PAD_LINK_REFUSED"
	}
	panic("Wrong value of PadLinkReturn variable")
}

type PadDirection C.GstPadDirection

const (
	PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
	PAD_SRC     = PadDirection(C.GST_PAD_SRC)
	PAD_SINK    = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) g() C.GstPadDirection {
	return C.GstPadDirection(p)
}

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type PadPresence C.GstPadPresence

const (
	PAD_ALWAYS    = PadPresence(C.GST_PAD_ALWAYS)
	PAD_SOMETIMES = PadPresence(C.GST_PAD_SOMETIMES)
	PAD_REQUEST   = PadPresence(C.GST_PAD_REQUEST)
)

func (p PadPresence) g() C.GstPadPresence {
	return C.GstPadPresence(p)
}

func (p PadPresence) String() string {
	switch p {
	case PAD_ALWAYS:
		return "PAD_ALWAYS"
	case PAD_SOMETIMES:
		return "PAD_SOMETIMES"
	case PAD_REQUEST:
		return "PAD_REQUEST"
	}
	panic("Wrong value of PadPresence variable")
}

type Pad struct {
	GstObj
}

//LiD: add GstPad Type() interface
func (p *Pad) Type() glib.Type {
	return glib.TypeFromName("GstPad")
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

func (p *Pad) CanLink(sink_pad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sink_pad.g()) != 0
}

func (p *Pad) Link(sink_pad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sink_pad.g()))
}

func (p *Pad) IsLinked() bool {
	return C.gst_pad_is_linked(p.g()) != 0
}

func (p *Pad) QueryCaps() *Caps {
	return (*Caps)(C.gst_pad_query_caps(p.g(), nil))
}

func (p *Pad) HasCurrentCaps() bool {
	return C.gst_pad_has_current_caps(p.g()) != 0
}

func (p *Pad) GetCurrentCaps() *Caps {
	return (*Caps)(C.gst_pad_get_current_caps(p.g()))
}

func (p *Pad) GetAllowedCaps() *Caps {
	return (*Caps)(C.gst_pad_get_allowed_caps(p.g()))
}

type GhostPad struct {
	Pad
}

func (p *GhostPad) g() *C.GstGhostPad {
	return (*C.GstGhostPad)(p.GetPtr())
}

func (p *GhostPad) AsGhostPad() *GhostPad {
	return p
}

func (p *GhostPad) SetTarget(new_target *Pad) bool {
	return C.gst_ghost_pad_set_target(p.g(), new_target.g()) != 0
}

func (p *GhostPad) GetTarget() *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_ghost_pad_get_target(p.g())))
	return r
}

func (p *GhostPad) Construct() bool {
	return C.gst_ghost_pad_construct(p.g()) != 0
}

func NewGhostPad(name string, target *Pad) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new(s, target.g())))
	return p
}

func NewGhostPadNoTarget(name string, dir PadDirection) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new_no_target(s, dir.g())))
	return p
}