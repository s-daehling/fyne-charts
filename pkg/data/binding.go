package data

import (
	"errors"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type base struct {
	listeners []binding.DataListener
	lock      sync.RWMutex
}

// AddListener allows a data listener to be informed of changes to this item.
func (b *base) AddListener(l binding.DataListener) {
	fyne.Do(func() {
		b.listeners = append(b.listeners, l)
		l.DataChanged()
	})
}

// RemoveListener should be called if the listener is no longer interested in being informed of data change events.
func (b *base) RemoveListener(l binding.DataListener) {
	fyne.Do(func() {
		for i, listener := range b.listeners {
			if listener == l {
				// Delete without preserving order:
				lastIndex := len(b.listeners) - 1
				b.listeners[i] = b.listeners[lastIndex]
				b.listeners[lastIndex] = nil
				b.listeners = b.listeners[:lastIndex]
				return
			}
		}
	})
}

func (b *base) trigger() {
	fyne.Do(b.triggerFromMain)
}

func (b *base) triggerFromMain() {
	for _, listen := range b.listeners {
		listen.DataChanged()
	}
}

type ConstrainedList[T any] interface {
	binding.DataList

	Add(v T) (err error)
	Get() (l []T, err error)
	GetValue(index int) (v T, err error)
	Remove(index int) (err error)
	Set(l []T) (err error)
	SetValue(index int, v T) (err error)
}

func NewConstrainedList[T any]() (l ConstrainedList[T]) {
	t := newList[T]()
	l = &t
	return
}

type boundList[T any] struct {
	base
	val   *[]T
	items []binding.DataItem
}

func newList[T any]() boundList[T] {
	return boundList[T]{val: new([]T)}
}

// GetItem returns the DataItem at the specified index.
func (bl *boundList[T]) GetItem(i int) (binding.DataItem, error) {
	bl.lock.RLock()
	defer bl.lock.RUnlock()

	if i < 0 || i >= len(bl.items) {
		return nil, errors.New("index out of bounds")
	}

	return bl.items[i], nil
}

// Length returns the number of items in this data list.
func (bl *boundList[T]) Length() int {
	bl.lock.RLock()
	defer bl.lock.RUnlock()

	return len(bl.items)
}

func (bl *boundList[T]) appendItem(i binding.DataItem) {
	bl.items = append(bl.items, i)
}

func (bl *boundList[T]) deleteItem(i int) {
	bl.items = append(bl.items[:i], bl.items[i+1:]...)
}

func (bl *boundList[T]) Add(val T) (err error) {
	bl.lock.Lock()
	*bl.val = append(*bl.val, val)

	trigger, err := bl.doReload()
	bl.lock.Unlock()

	if trigger {
		bl.trigger()
	}
	return
}

func (bl *boundList[T]) Get() (l []T, err error) {
	bl.lock.RLock()
	defer bl.lock.RUnlock()

	return *bl.val, nil
}

func (bl *boundList[T]) GetValue(index int) (val T, err error) {
	bl.lock.RLock()
	defer bl.lock.RUnlock()

	if index < 0 || index >= bl.Length() {
		return *new(T), errors.New("index out of bounds")
	}

	return (*bl.val)[index], nil
}

func (bl *boundList[T]) Remove(index int) (err error) {
	bl.lock.Lock()

	v := *bl.val
	if len(v) == 0 {
		bl.lock.Unlock()
		return nil
	}
	if index == 0 {
		*bl.val = v[1:]
	} else if index == len(v)-1 {
		*bl.val = v[:len(v)-1]
	} else {
		*bl.val = append(v[:index], v[index+1:]...)
	}

	trigger, err := bl.doReload()
	bl.lock.Unlock()

	if trigger {
		bl.trigger()
	}
	return
}

func (bl *boundList[T]) Set(l []T) (err error) {
	bl.lock.Lock()
	*bl.val = l
	trigger, err := bl.doReload()
	bl.lock.Unlock()

	if trigger {
		bl.trigger()
	}
	return
}

func (bl *boundList[T]) SetValue(index int, val T) (err error) {
	bl.lock.RLock()
	len := bl.Length()
	bl.lock.RUnlock()

	if index < 0 || index >= len {
		return errors.New("index out of bounds")
	}

	bl.lock.Lock()
	(*bl.val)[index] = val
	bl.lock.Unlock()

	item, err := bl.GetItem(index)
	if err != nil {
		return err
	}
	return item.(binding.Item[T]).Set(val)
}

func (bl *boundList[T]) doReload() (trigger bool, retErr error) {
	oldLen := len(bl.items)
	newLen := len(*bl.val)
	if oldLen > newLen {
		for i := oldLen - 1; i >= newLen; i-- {
			bl.deleteItem(i)
		}
		trigger = true
	} else if oldLen < newLen {
		for i := oldLen; i < newLen; i++ {
			item := bindListItem(bl.val, i)
			bl.appendItem(item)
		}
		trigger = true
	}

	for i, item := range bl.items {
		if i > oldLen || i > newLen {
			break
		}

		err := item.(*boundListItem[T]).doSet((*bl.val)[i])
		if err != nil {
			retErr = err
		}
	}
	return trigger, retErr
}

type boundListItem[T any] struct {
	base

	val   *[]T
	index int
}

func bindListItem[T any](v *[]T, i int) binding.Item[T] {
	return &boundListItem[T]{val: v, index: i}
}

func (bli *boundListItem[T]) Get() (T, error) {
	bli.lock.Lock()
	defer bli.lock.Unlock()

	if bli.index < 0 || bli.index >= len(*bli.val) {
		return *new(T), errors.New("index out of bounds")
	}

	return (*bli.val)[bli.index], nil
}

func (bli *boundListItem[T]) Set(val T) error {
	return bli.doSet(val)
}

func (bli *boundListItem[T]) doSet(val T) error {
	bli.lock.Lock()
	(*bli.val)[bli.index] = val
	bli.lock.Unlock()

	bli.trigger()
	return nil
}
