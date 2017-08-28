#include <stdlib.h>
#include <gst/gst.h>
#include "_cgo_export.h"

GstPadProbeReturn gst_probe_callback_handler_dispatch(GstPad *pad, GstPadProbeInfo *info, gpointer user_data) {
  return gst_probe_callback_handler(pad, info, user_data);
}
