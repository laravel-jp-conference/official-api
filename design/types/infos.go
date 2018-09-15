package types

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var PlaceInfoType = Type("PlaceInfoType", func() {
	Member("_links", HalLinkType, "navigation")
	Member("_embedded", EmbeddedPlaceResponseType, "navigation")
})
