package routine

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

func TestCreateInheritedMap(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		thd := currentThread(true)
		assert.NotNil(t, thd)
		assert.Nil(t, thd.inheritableThreadLocals)
		thd.inheritableThreadLocals = &threadLocalMap{}
		assert.Nil(t, thd.inheritableThreadLocals.table)
		assert.Nil(t, createInheritedMap())
		//
		wg.Done()
	}()
	wg.Wait()
}

func TestCreateInheritedMapNil(t *testing.T) {
	tls := NewInheritableThreadLocal()
	tls.Set(nil)
	srcValue := tls.Get()
	assert.Nil(t, srcValue)
	assert.True(t, srcValue == nil)

	mp := createInheritedMap()
	assert.NotNil(t, mp)
	getValue := mp.get(tls)
	assert.Nil(t, getValue)
	assert.True(t, getValue == nil)

	mp2 := createInheritedMap()
	assert.NotNil(t, mp2)
	assert.NotSame(t, mp, mp2)
	getValue2 := mp2.get(tls)
	assert.Nil(t, getValue2)
	assert.True(t, getValue2 == nil)
}

func TestCreateInheritedMapValue(t *testing.T) {
	tls := NewInheritableThreadLocal()
	value := rand.Uint64()
	tls.Set(value)
	srcValue := tls.Get()
	assert.NotSame(t, &value, &srcValue)
	assert.Equal(t, value, srcValue)

	mp := createInheritedMap()
	assert.NotNil(t, mp)
	getValue := mp.get(tls)
	assert.NotSame(t, &value, &getValue)
	assert.Equal(t, value, getValue)

	mp2 := createInheritedMap()
	assert.NotNil(t, mp2)
	assert.NotSame(t, mp, mp2)
	getValue2 := mp2.get(tls)
	assert.NotSame(t, &value, &getValue2)
	assert.Equal(t, value, getValue2)
}

func TestCreateInheritedMapStruct(t *testing.T) {
	tls := NewInheritableThreadLocal()
	value := personCloneable{Id: 1, Name: "Hello"}
	tls.Set(value)
	srcValue := tls.Get()
	assert.NotSame(t, &value, &srcValue)
	assert.Equal(t, value, srcValue)

	mp := createInheritedMap()
	assert.NotNil(t, mp)
	getValue := mp.get(tls)
	assert.NotSame(t, &value, &getValue)
	assert.Equal(t, value, getValue)

	mp2 := createInheritedMap()
	assert.NotNil(t, mp2)
	assert.NotSame(t, mp, mp2)
	getValue2 := mp2.get(tls)
	assert.NotSame(t, &value, &getValue2)
	assert.Equal(t, value, getValue2)
}

func TestCreateInheritedMapPointer(t *testing.T) {
	tls := NewInheritableThreadLocal()
	value := &person{Id: 1, Name: "Hello"}
	tls.Set(value)
	srcValue := tls.Get().(*person)
	assert.Same(t, value, srcValue)
	assert.Equal(t, *value, *srcValue)

	mp := createInheritedMap()
	assert.NotNil(t, mp)
	getValue := mp.get(tls).(*person)
	assert.Same(t, value, getValue)
	assert.Equal(t, *value, *getValue)

	mp2 := createInheritedMap()
	assert.NotNil(t, mp2)
	assert.NotSame(t, mp, mp2)
	getValue2 := mp2.get(tls).(*person)
	assert.Same(t, value, getValue2)
	assert.Equal(t, *value, *getValue2)
}

func TestCreateInheritedMapCloneable(t *testing.T) {
	tls := NewInheritableThreadLocal()
	value := &personCloneable{Id: 1, Name: "Hello"}
	tls.Set(value)
	srcValue := tls.Get().(*personCloneable)
	assert.Same(t, value, srcValue)
	assert.Equal(t, *value, *srcValue)

	mp := createInheritedMap()
	assert.NotNil(t, mp)
	getValue := mp.get(tls).(*personCloneable)
	assert.NotSame(t, value, getValue)
	assert.Equal(t, *value, *getValue)

	mp2 := createInheritedMap()
	assert.NotNil(t, mp2)
	assert.NotSame(t, mp, mp2)
	getValue2 := mp2.get(tls).(*personCloneable)
	assert.NotSame(t, value, getValue2)
	assert.Equal(t, *value, *getValue2)
}

func TestFill(t *testing.T) {
	a := make([]Any, 6)
	fill(a, 4, 5, unset)
	for i := 0; i < 6; i++ {
		if i == 4 {
			assert.True(t, a[i] == unset)
		} else {
			assert.Nil(t, a[i])
			assert.True(t, a[i] != unset)
		}
	}
}
