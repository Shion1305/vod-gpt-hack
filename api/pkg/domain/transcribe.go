package domain

type TranscriptionPartial struct {
	ID      int
	Content string
	Start   float32
	End     float32
}

type TranscriptionGrouped struct {
	Content string
	Start   float32
	End     float32
}
