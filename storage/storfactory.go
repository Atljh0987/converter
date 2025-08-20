package storage

import f "converter/functions"

type StorageFactory interface {
	CreateStorage(StorageType) Storage
}

type DefaultFactory struct{}

func (d DefaultFactory) CreateStorage(st StorageType) f.Either[f.Error, Storage] {

	switch st {
	case MapStorageType:
		return f.Right[f.Error, Storage](NewMapStorage())
	default:
		return f.Left[f.Error, Storage](f.NewError("Unknown StorageType"))
	}
}
