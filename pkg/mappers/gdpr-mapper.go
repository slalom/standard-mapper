package mappers

import (
	"github.com/tredfield/standard-mapper/pkg/types"
)

// MapGdpr maps from Gdpr to Standard
func MapGdpr(gdpr types.Gdpr) (*types.Standard, error) {
	var standard types.Standard

	for _, division := range gdpr.Division {
		// todo map division to standard

		for _, article := range division.ARTICLE {
			var child types.Children
			child.ID = article.Identifier

			standard.Children = append(standard.Children, child)

			// todo finish mapping article to standard
		}
	}

	return &standard, nil
}
