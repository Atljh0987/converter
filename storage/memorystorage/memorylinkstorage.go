package memorystorage

type MemoryLinkStorage struct {
	storage []string
}

func NewMemoryLinkStorage() *MemoryLinkStorage {
	return &MemoryLinkStorage{
		storage: []string{},
	}
}

func (o *MemoryLinkStorage) Get(i int) string {
	return o.storage[i]
}

func (o *MemoryLinkStorage) GetCount() int {
	return len(o.storage)
}

func (o *MemoryLinkStorage) Create(link string) {
	for _, item := range o.storage {
        if item == link {
            return
        }
    }

	o.storage = append(o.storage, link)
}

func (o *MemoryLinkStorage) Delete(i int) {
	o.storage = append(o.storage[:i], o.storage[i+1:]...)
}
