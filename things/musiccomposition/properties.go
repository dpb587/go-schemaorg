package musiccomposition

import "github.com/dpb587/go-schemaorg"

var (
	// The type of composition (e.g. overture, sonata, symphony, etc.).
	MusicCompositionForm = schemaorg.NewProperty("musicCompositionForm")

	// Smaller compositions included in this work (e.g. a movement in a symphony).
	IncludedComposition = schemaorg.NewProperty("includedComposition")

	// The person who wrote the words.
	Lyricist = schemaorg.NewProperty("lyricist")

	// An audio recording of the work.
	RecordedAs = schemaorg.NewProperty("recordedAs")

	// The date and place the work was first performed.
	FirstPerformance = schemaorg.NewProperty("firstPerformance")

	// An arrangement derived from the composition.
	MusicArrangement = schemaorg.NewProperty("musicArrangement")

	// The words in the song.
	Lyrics = schemaorg.NewProperty("lyrics")

	// The International Standard Musical Work Code for the composition.
	IswcCode = schemaorg.NewProperty("iswcCode")

	// The person or organization who wrote a composition, or who is the composer of
	// a work performed at some event.
	Composer = schemaorg.NewProperty("composer")

	// The key, mode, or scale this composition uses.
	MusicalKey = schemaorg.NewProperty("musicalKey")
)
