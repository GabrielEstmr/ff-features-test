package ff

import ff_resources "ff-features-go/ff/resources"

type RegisterMethods interface {
	RegisterFeatures(features ff_resources.Features) error
}
