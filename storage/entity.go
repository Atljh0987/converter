package storage

type Extension interface {
	String() string
}
type VideoExtension string
type ImageExtension string

func (v VideoExtension) String() string {
	return string(v)
}

func (i ImageExtension) String() string {
	return string(i)
}

const (
	WEBM VideoExtension = "webm"
	MP4  VideoExtension = "mp4"
)

func VideoExtensions() []VideoExtension {
	return []VideoExtension{WEBM, MP4}
}

func VideoExtensionsStrings() []string {
	extensions := VideoExtensions()
	result := make([]string, 0, len(extensions))

	for _, ext := range extensions {
		result = append(result, ext.String())
	}

	return result
}

const (
	WEBP ImageExtension = "webp"
	JPEG ImageExtension = "jpeg"
)

type Entity struct {
	uri           string
	extensionFrom Extension
	extensionTo   Extension
}
