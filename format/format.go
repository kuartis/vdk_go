package format

import (
	"github.com/kuartis/deepch_mas/av/avutil"
	"github.com/kuartis/deepch_mas/format/aac"
	"github.com/kuartis/deepch_mas/format/flv"
	"github.com/kuartis/deepch_mas/format/mp4"
	"github.com/kuartis/deepch_mas/format/rtmp"
	"github.com/kuartis/deepch_mas/format/rtsp"
	"github.com/kuartis/deepch_mas/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
