package kv_locks

import (
	"context"

	kv_storage "github.com/storage-lock/go-kv-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

// LockFactory 基于 KV 存储的锁工厂
type LockFactory struct {
	*storage_lock_factory.StorageLockFactory[any]
}

var newKVStorage = kv_storage.NewKVStorage

// NewLockFactory 创建锁工厂
func NewLockFactory() (*LockFactory, error) {
	s, err := newKVStorage(context.Background())
	if err != nil {
		return nil, err
	}
	factory := storage_lock_factory.NewStorageLockFactory[any](s, nil)
	return &LockFactory{
		StorageLockFactory: factory,
	}, nil
}
