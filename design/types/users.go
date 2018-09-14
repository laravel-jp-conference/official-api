package types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var SpeakerResponseType = Type("SpeakerResponseType", func() {
	Member("_links", HalLinkSpeakerType, "navigation")
	Member("name", String, "Speaker Name")
	Member("image", String, "Speaker Image URL")
})
