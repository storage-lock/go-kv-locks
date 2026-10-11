package kv_locks

import (
	"context"

	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var (
	globalLockFactory = storage_lock_factory.NewStorageLockFactoryBeanFactory[string, any]()

	// newLockFactoryFunc 指向 NewLockFactory，包级变量以便单测注入构造失败，
	// 覆盖 GetOrInit initFunc 与 NewLock 的错误分支。
	newLockFactoryFunc = NewLockFactory
)

// NewLock 创建基于 KV 存储的分布式锁（进程内语义）
func NewLock(ctx context.Context, lockId string) (*storage_lock.StorageLock, error) {
	factory, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return factory.CreateLock(lockId)
}

// NewLockWithOptions 创建基于 KV 存储的分布式锁（自定义选项）
func NewLockWithOptions(ctx context.Context, options *storage_lock.StorageLockOptions) (*storage_lock.StorageLock, error) {
	factory, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return factory.CreateLockWithOptions(options)
}

// GetLockFactory 获取全局锁工厂（单例）
func GetLockFactory(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
	return globalLockFactory.GetOrInit(ctx, "x", func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
		factory, err := newLockFactoryFunc()
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
}

// GetGlobalLockFactory 返回全局工厂管理器
func GetGlobalLockFactory() *storage_lock_factory.StorageLockFactoryBeanFactory[string, any] {
	return globalLockFactory
}
