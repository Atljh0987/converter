package storage

type StorageType string

var (
	MapStorageType StorageType = "mapstorage"
)

type Storage interface{
	Create()
	Get()
	Update()
	Delete()
}

type MapStorage struct{
	storage map[string]Entity
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		storage: make(map[string]Entity),
	}
}

func (m MapStorage) Create() {

}

func (m MapStorage) Get() {
	
}

func (m MapStorage) Update() {
	
}

func (m MapStorage) Delete() {
	
}