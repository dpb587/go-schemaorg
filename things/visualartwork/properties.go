package visualartwork

import "github.com/dpb587/go-schemaorg"

var (
	// The height of the item.
	Height = schemaorg.NewProperty("height")

	// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble,
	// Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil,
	// Mixed Media, etc.)
	ArtMedium = schemaorg.NewProperty("artMedium")

	// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage,
	// etc.
	Artform = schemaorg.NewProperty("artform")

	// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board,
	// etc.
	ArtworkSurface = schemaorg.NewProperty("artworkSurface")

	// The number of copies when multiple copies of a piece of artwork are produced
	// - e.g. for a limited edition of 20 prints, 'artEdition' refers to the total
	// number of copies (in this example "20").
	ArtEdition = schemaorg.NewProperty("artEdition")

	// The width of the item.
	Width = schemaorg.NewProperty("width")

	// A material used as a surface in some artwork, e.g. Canvas, Paper, Wood,
	// Board, etc.
	Surface = schemaorg.NewProperty("surface")

	// The depth of the item.
	Depth = schemaorg.NewProperty("depth")
)
