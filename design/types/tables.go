package types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TimetableType = Type("TimetableType", func() {
	Member("_links", HalLinkType, "navigation")
	Member("title", String, "講演タイトル")
	Member("body", String, "講演内容")
	Member("_embedded", HalLinkSpeakerType)
	Member("from_to", String, "講演時間")
	Member("place", String, "ホール")
})

var VndErrorType = Type("VndErrorType", func(){
	Member("message", String)
	Member("path", String)
	Member("logref", Integer)
	Member("_links", VndLinkType)
})
