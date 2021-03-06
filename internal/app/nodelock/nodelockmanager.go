package nodelock

import "sync"

type LockManager interface {
	GetLock(node string) sync.Locker
}

type DefaultLockManager struct {
	permits uint
	maps    sync.Map
}

func NewLockManager(permits uint) LockManager {
	return DefaultLockManager{permits: permits, maps: sync.Map{}}
}

func (m DefaultLockManager) GetLock(node string) sync.Locker {
	maps := m.maps

	value, ok := maps.Load(node)

	if ok {
		return value.(sync.Locker)
	}

	var locker sync.Locker
	locker = NewLocker(m.permits)
	maps.Store(node, locker)

	return locker
}
