package types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var EmbeddedPlaceResponseType = Type("EmbeddedPlaceResponseType", func() {
	Member("places", PlaceCollectType, "embedded")
	Required("places")
})

var PlaceCollectType = Type("PlaceCollectType", func() {
	Member("_links", HalLinkType, "navigation")
	Member("name", String, "会場名")
	Member("address", String, "住所")
	Member("latitude", String, "緯度")
	Member("longitude", String, "軽度")
	Member("note", String, "備考")
	Required("_links", "name", "address", "latitude", "longitude")
})
