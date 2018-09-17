package types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var SponsorResponseType = Type("SponsorResponseType", func() {
	Member("_links", HalLinkType, "navigation")
	Member("name", String, "Sponsor Name")
	Member("image", String, "Sponsor Image URL")
	Member("typeID", Integer, "Sponsor Type ID")
	Member("typeName", Integer, "Sponsor Type Name")
})
