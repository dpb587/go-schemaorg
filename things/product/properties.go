package product

import "github.com/dpb587/go-schemaorg"

var (
	// The manufacturer of the product.
	Manufacturer = schemaorg.NewProperty("manufacturer")

	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a
	// product or service, or the product to which the offer refers.
	Sku = schemaorg.NewProperty("sku")

	// An intended audience, i.e. a group for whom something was created.
	Audience = schemaorg.NewProperty("audience")

	// The Manufacturer Part Number (MPN) of the product, or the product to which
	// the offer refers.
	Mpn = schemaorg.NewProperty("mpn")

	// The height of the item.
	Height = schemaorg.NewProperty("height")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx">GTIN-8</a>
	// code of the product, or the product to which the offer refers. This code is
	// also known as EAN/UCC-8 or 8-digit EAN. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin8 = schemaorg.NewProperty("gtin8")

	// Review of the item.
	Reviews = schemaorg.NewProperty("reviews")

	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating = schemaorg.NewProperty("aggregateRating")

	// A pointer to another product (or multiple products) for which this product is
	// a consumable.
	IsConsumableFor = schemaorg.NewProperty("isConsumableFor")

	// An offer to provide this item&#x2014;for example, an offer to sell a product,
	// rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers = schemaorg.NewProperty("offers")

	// An award won by or for this item.
	Award = schemaorg.NewProperty("award")

	// A category for the item. Greater signs or slashes can be used to informally
	// indicate a category hierarchy.
	Category = schemaorg.NewProperty("category")

	// The width of the item.
	Width = schemaorg.NewProperty("width")

	// Awards won by or for this item.
	Awards = schemaorg.NewProperty("awards")

	// A property-value pair representing an additional characteristics of the
	// entitity, e.g. a product feature or another characteristic for which there is
	// no matching property in schema.org.</p>
	// 
	// <p>Note: Publishers should be aware that applications designed to use
	// specific schema.org properties (e.g. http://schema.org/width,
	// http://schema.org/color, http://schema.org/gtin13, ...) will typically expect
	// such data to be provided using those properties, rather than using the
	// generic property/value mechanism.
	AdditionalProperty = schemaorg.NewProperty("additionalProperty")

	// A pointer to another product (or multiple products) for which this product is
	// an accessory or spare part.
	IsAccessoryOrSparePartFor = schemaorg.NewProperty("isAccessoryOrSparePartFor")

	// An associated logo.
	Logo = schemaorg.NewProperty("logo")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx">GTIN-14</a>
	// code of the product, or the product to which the offer refers. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin14 = schemaorg.NewProperty("gtin14")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx">GTIN-13</a>
	// code of the product, or the product to which the offer refers. This is
	// equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes
	// can be converted into a GTIN-13 code by simply adding a preceeding zero. See
	// <a href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN
	// Summary</a> for more details.
	Gtin13 = schemaorg.NewProperty("gtin13")

	// The <a href="http://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx">GTIN-12</a>
	// code of the product, or the product to which the offer refers. The GTIN-12 is
	// the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item
	// Reference, and Check Digit used to identify trade items. See <a
	// href="http://www.gs1.org/barcodes/technical/idkeys/gtin">GS1 GTIN Summary</a>
	// for more details.
	Gtin12 = schemaorg.NewProperty("gtin12")

	// A material that something is made from, e.g. leather, wool, cotton, paper.
	Material = schemaorg.NewProperty("material")

	// The weight of the product or person.
	Weight = schemaorg.NewProperty("weight")

	// The depth of the item.
	Depth = schemaorg.NewProperty("depth")

	// A pointer to another, functionally similar product (or multiple products).
	IsSimilarTo = schemaorg.NewProperty("isSimilarTo")

	// The model of the product. Use with the URL of a ProductModel or a textual
	// representation of the model identifier. The URL of the ProductModel can be
	// from an external source. It is recommended to additionally provide strong
	// product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	Model = schemaorg.NewProperty("model")

	// The color of the product.
	Color = schemaorg.NewProperty("color")

	// A pointer to another, somehow related product (or multiple products).
	IsRelatedTo = schemaorg.NewProperty("isRelatedTo")

	// The product identifier, such as ISBN. For example: <code>meta
	// itemprop="productID" content="isbn:123-456-789"</code>.
	ProductID = schemaorg.NewProperty("productID")

	// A review of the item.
	Review = schemaorg.NewProperty("review")

	// The date the item e.g. vehicle was purchased by the current owner.
	PurchaseDate = schemaorg.NewProperty("purchaseDate")

	// A predefined value from OfferItemCondition or a textual description of the
	// condition of the product or service, or the products or services included in
	// the offer.
	ItemCondition = schemaorg.NewProperty("itemCondition")

	// The date of production of the item, e.g. vehicle.
	ProductionDate = schemaorg.NewProperty("productionDate")

	// The brand(s) associated with a product or service, or the brand(s) maintained
	// by an organization or business person.
	Brand = schemaorg.NewProperty("brand")

	// The release date of a product or product model. This can be used to
	// distinguish the exact variant of a product.
	ReleaseDate = schemaorg.NewProperty("releaseDate")
)
