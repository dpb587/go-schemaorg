package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dpb587/go-schemaorg/datatypes/datetime"
	"github.com/dpb587/go-schemaorg/datatypes/float"
	"github.com/dpb587/go-schemaorg/datatypes/integer"
	"github.com/dpb587/go-schemaorg/datatypes/text"
	"github.com/dpb587/go-schemaorg/datatypes/url"
	"github.com/dpb587/go-schemaorg/things/brand"
	"github.com/dpb587/go-schemaorg/things/creativework"
	"github.com/dpb587/go-schemaorg/things/imageobject"
	"github.com/dpb587/go-schemaorg/things/mediaobject"
	"github.com/dpb587/go-schemaorg/things/offer"
	"github.com/dpb587/go-schemaorg/things/organization"
	"github.com/dpb587/go-schemaorg/things/product"
	"github.com/dpb587/go-schemaorg/things/productmodel"
	"github.com/dpb587/go-schemaorg/things/quantitativevalue"
	"github.com/dpb587/go-schemaorg/things/thing"
)

func main() {
	subject := product.New().
		Add(thing.Identifier, url.New("https://www.theloopyewe.com/Hedgehog-Fibres-Merino-Lace-2A42E7E3?p=2D7606E5-Rusty-Nail#product")).
		Add(thing.Image, imageobject.New().
			Add(mediaobject.ContentUrl, url.New("https://www.theloopyewe.com/asset/catalog-entry-photo/462bc06d-774e-fb77-6ede-f493846f958b~v2-702x702.jpg")).
			Add(mediaobject.ContentSize, text.New("142412")).
			Add(imageobject.Thumbnail, imageobject.New().
				Add(mediaobject.ContentUrl, url.New("https://www.theloopyewe.com/asset/catalog-entry-photo/462bc06d-774e-fb77-6ede-f493846f958b~v2-96x96.jpg")).
				Add(mediaobject.ContentSize, text.New("4863")).
				Add(mediaobject.Height, quantitativevalue.New().
					Add(quantitativevalue.Value, integer.New(96)),
				).
				Add(mediaobject.Width, quantitativevalue.New().
					Add(quantitativevalue.Value, integer.New(96)),
				).
				Add(creativework.CopyrightHolder, organization.New().
					Add(thing.Identifier, url.New("https://www.theloopyewe.com/#organization")).
					Add(thing.Name, text.New("The Loopy Ewe")),
				).
				Add(creativework.CopyrightYear, integer.New(2016)).
				Add(creativework.DateCreated, datetime.MustParse(time.RFC3339, "2016-11-21T10:31:18Z")),
			).
			Add(mediaobject.Height, quantitativevalue.New().
				Add(quantitativevalue.Value, integer.New(702)),
			).
			Add(mediaobject.Width, quantitativevalue.New().
				Add(quantitativevalue.Value, integer.New(702)),
			).
			Add(creativework.CopyrightHolder, organization.New().
				Add(thing.Identifier, url.New("https://www.theloopyewe.com/#organization")).
				Add(thing.Name, text.New("The Loopy Ewe")),
			).
			Add(creativework.CopyrightYear, integer.New(2016)).
			Add(creativework.DateCreated, datetime.MustParse(time.RFC3339, "2016-11-21T10:31:18Z")),
		).
		Add(thing.Name, text.New("Hedgehog Fibres, Merino Lace [Rusty Nail]")).
		Add(thing.Description, text.New(`This extra fine merino is soft and springy. This wonderful yarn will work great for a lace project.

The colors of the Hedgehog Fibres Merino Lace shift beautifully from one to another. Randomly dyed to prevent striping or pooling as much as possible, the yarn creates a variegated effect that looks good in plain stockinette or advanced stitch motifs.

The dye is applied randomly, therefore no two hanks are the same.`)).
		Add(product.Brand, brand.New().
			Add(thing.Name, text.New("Hedgehog Fibres")).
			Add(thing.Identifier, url.New("https://www.theloopyewe.com/Hedgehog-Fibres-A1A1CE0M#brand")).
			Add(thing.SameAs, url.New("https://shop.hedgehogfibres.com/#brand")),
		).
		Add(product.Color, text.New("Rusty Nail")).
		Add(product.Material, text.New("Merino Wool")).
		Add(product.Model, productmodel.New().
			Add(thing.Identifier, url.New("https://www.theloopyewe.com/Hedgehog-Fibres-Merino-Lace-2A42E7E3#productmodel")).
			Add(thing.Name, text.New("Hedgehog Fibres, Merino Lace")),
		).
		Add(product.Offers, offer.New().
			Add(offer.Availability, url.New("http://schema.org/InStock")).
			Add(offer.InventoryLevel, quantitativevalue.New().
				Add(quantitativevalue.Value, integer.New(6)).
				Add(quantitativevalue.UnitText, text.New("skein")),
			).
			Add(offer.PriceCurrency, text.New("USD")).
			Add(offer.Price, float.New(37.00)).
			Add(offer.ItemCondition, url.New("http://schema.org/NewCondition")),
		).
		Add(product.Sku, text.New("2D7606E5"))

	bytes, err := json.MarshalIndent(subject, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bytes)
}
