package mosir_sdk_go

// MediaFileProfile is the enum for media file profiles.
type MediaFileProfile string

const (
	MediaFileProfileAnimatedCompatible MediaFileProfile = "ANIMATED_COMPATIBLE"
	MediaFileProfileAnimatedThumbnail  MediaFileProfile = "ANIMATED_THUMBNAIL"
	MediaFileProfileCompatible         MediaFileProfile = "COMPATIBLE"
	MediaFileProfileQuality            MediaFileProfile = "QUALITY"
	MediaFileProfileThumbnail          MediaFileProfile = "THUMBNAIL"
)

// MediaFileMetadata represents a media file's metadata.
type MediaFileMetadata struct {
	ContentType string
	ID          string
	MediaID     string
	Profile     MediaFileProfile
	URL         string
}

// MediaMetadata represents the metadata for a piece of media.
type MediaMetadata struct {
	AccountID   string
	AspectRatio *struct {
		Denominator int
		Numerator   int
	}
	BlurHash   string
	DurationMs *int
	Files      []MediaFileMetadata
	ID         string
	Status     string
	Type       string
}
