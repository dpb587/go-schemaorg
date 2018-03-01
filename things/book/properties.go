package book

import "github.com/dpb587/go-schemaorg"

var (
	// The ISBN of the book.
	Isbn = schemaorg.NewProperty("isbn")

	// The edition of the book.
	BookEdition = schemaorg.NewProperty("bookEdition")

	// The format of the book.
	BookFormat = schemaorg.NewProperty("bookFormat")

	// The number of pages in the book.
	NumberOfPages = schemaorg.NewProperty("numberOfPages")

	// The illustrator of the book.
	Illustrator = schemaorg.NewProperty("illustrator")
)
