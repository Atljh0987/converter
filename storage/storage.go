package storage

import mapstorage "converter/storage/memorystorage"

// var storage = mapstorage.NewMemoryStorage()
var linkstorage = mapstorage.NewMemoryLinkStorage()

type LinkStorage interface {
	GetCount() int
	Get(i int) string
	Create(link string)
	Delete(i int)
}

type Storage interface {
	Create()
	Get()
	Update()
	Delete()
}

func GetLinkStorage() LinkStorage {
	return linkstorage
}
