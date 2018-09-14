package types

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var HalLinkType = Type("HalLinkType", func() {
	Member("self", HalLinkUnit, "The URI of the main resource being represented ('/path/to') expressed with a self link")
	Required("self")
})

var HalLinkUnit = Type("HalLinkUnitType", func() {
	Member("href", String, func() {
		Description("for Hyper Reference")
		Default("")
	})
	Required("href")
})

var HalLinkSpeakerType = Type("HalLinkSpeakerType", func() {
	Member("self-twitter", HalLinkUnit, "The URI of the main resource being represented ('/path/to') expressed with a self link")
	Member("self-github", HalLinkType)
	Member("self-facebook", HalLinkType)
})

var VndLinkType = Type("VndLinkType", func() {
	Member("about", HalLinkUnit, "The URI of the main resource being represented ('/path/to') expressed with a self link")
	Required("about")
})
