package engine

import "embed"

var embeddedFiles embed.FS

type Engine interface{
	convert()
}

type FFMpeg struct{}

func (f FFMpeg) convert() {
	
}