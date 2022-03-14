package format

import (
	"github.com/kuartis/deepch_kua/av/avutil"
	"github.com/kuartis/deepch_kua/format/aac"
	"github.com/kuartis/deepch_kua/format/flv"
	"github.com/kuartis/deepch_kua/format/mp4"
	"github.com/kuartis/deepch_kua/format/rtmp"
	"github.com/kuartis/deepch_kua/format/rtsp"
	"github.com/kuartis/deepch_kua/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
