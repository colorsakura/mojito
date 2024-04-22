package mpv

//#include <mpv/client.h>
import "C"

// Errors mpv_error
const (
	ERROR_SUCCESS Error = C.MPV_ERROR_SUCCESS

	MPV_ERROR_EVENT_QUEUE_FULL Error = C.MPV_ERROR_EVENT_QUEUE_FULL
)
