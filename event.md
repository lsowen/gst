**progress**
0/69
```
    #define 	GST_EVENT_MAKE_TYPE()
    #define 	GST_EVENT_TYPE()
    #define 	GST_EVENT_TYPE_NAME()
    #define 	GST_EVENT_TIMESTAMP()
    #define 	GST_EVENT_SEQNUM()
    #define 	GST_EVENT_IS_UPSTREAM()
    #define 	GST_EVENT_IS_DOWNSTREAM()
    #define 	GST_EVENT_IS_SERIALIZED()
    #define 	GST_EVENT_IS_STICKY()
    GstEventTypeFlags 	gst_event_type_get_flags ()
    const gchar * 	gst_event_type_get_name ()
    GQuark 	gst_event_type_to_quark ()
    GstEvent * 	gst_event_ref ()
    void 	gst_event_unref ()
    gboolean 	gst_event_replace ()
    GstEvent * 	gst_event_copy ()
    GstEvent * 	gst_event_steal ()
    gboolean 	gst_event_take ()
    #define 	gst_event_is_writable()
    #define 	gst_event_make_writable()
    GstStructure * 	gst_event_writable_structure ()
    GstEvent * 	gst_event_new_custom ()
    const GstStructure * 	gst_event_get_structure ()
    gboolean 	gst_event_has_name ()
    guint32 	gst_event_get_seqnum ()
    void 	gst_event_set_seqnum ()
    gint64 	gst_event_get_running_time_offset ()
    void 	gst_event_set_running_time_offset ()
    GstEvent * 	gst_event_new_flush_start ()
    GstEvent * 	gst_event_new_flush_stop ()
    void 	gst_event_parse_flush_stop ()
    GstEvent * 	gst_event_new_eos ()
    GstEvent * 	gst_event_new_gap ()
    void 	gst_event_parse_gap ()
    GstEvent * 	gst_event_new_stream_start ()
    void 	gst_event_parse_stream_start ()
    void 	gst_event_set_stream_flags ()
    void 	gst_event_parse_stream_flags ()
    void 	gst_event_set_group_id ()
    gboolean 	gst_event_parse_group_id ()
    GstEvent * 	gst_event_new_segment ()
    void 	gst_event_parse_segment ()
    void 	gst_event_copy_segment ()
    GstEvent * 	gst_event_new_tag ()
    void 	gst_event_parse_tag ()
    GstEvent * 	gst_event_new_buffer_size ()
    void 	gst_event_parse_buffer_size ()
    GstEvent * 	gst_event_new_qos ()
    void 	gst_event_parse_qos ()
    GstEvent * 	gst_event_new_seek ()
    void 	gst_event_parse_seek ()
    GstEvent * 	gst_event_new_navigation ()
    GstEvent * 	gst_event_new_latency ()
    void 	gst_event_parse_latency ()
    GstEvent * 	gst_event_new_step ()
    void 	gst_event_parse_step ()
    GstEvent * 	gst_event_new_sink_message ()
    void 	gst_event_parse_sink_message ()
    GstEvent * 	gst_event_new_reconfigure ()
    GstEvent * 	gst_event_new_caps ()
    void 	gst_event_parse_caps ()
    GstEvent * 	gst_event_new_toc ()
    void 	gst_event_parse_toc ()
    GstEvent * 	gst_event_new_toc_select ()
    void 	gst_event_parse_toc_select ()
    GstEvent * 	gst_event_new_segment_done ()
    void 	gst_event_parse_segment_done ()
    GstEvent * 	gst_event_new_protection ()
    void 	gst_event_parse_protection ()
```