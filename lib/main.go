package main

import "app/storage"

type recordStorageImpl struct {
	records []*storage.Record
}

func NewRecordStorage() storage.RecordStorage {
	return &recordStorageImpl{
		records: []*storage.Record{
			{
				ID:   5,
				Name: "Hello",
			},
			{
				ID:   7,
				Name: "World",
			},
		},
	}
}

func (s *recordStorageImpl) FindOne(id uint) (*storage.Record, error) {
	for _, record := range s.records {
		if record.ID == id {
			return record, nil
		}
	}
	return nil, storage.ErrRecordNotFound
}

func (s *recordStorageImpl) FindMany() ([]*storage.Record, error) {
	return s.records, nil
}
