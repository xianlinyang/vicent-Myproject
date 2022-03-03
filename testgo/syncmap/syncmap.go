package syncmap

import (
	"fmt"
	"sync"
)

type SyncMapExt struct {
	sync.Map
}

func (s *SyncMapExt) SyncMapGet(syncMap *sync.Map, keys ...interface{}) (value interface{}, ok bool) {
	for _, v := range keys {
		value, ok = syncMap.Load(v)
		switch value.(type) {
		default:
			//fmt.Printf("unexpected type %T", t)
		case *sync.Map:
			syncMap = value.(*sync.Map)
		}
	}
	return
}

func (s *SyncMapExt) SyncMapSet(syncMap *sync.Map, value interface{}, keys ...interface{}) bool {
	mapList := make([]*sync.Map, len(keys))
	mapList[0] = syncMap
	var key interface{}
	for i, k := range keys {
		value, ok := syncMap.Load(k)
		if ok {
			switch value.(type) {
			case *sync.Map:
				mapList[i+1] = value.(*sync.Map)
			}
		}
		key = k
	}

	mapList[len(mapList)-1].Store(key, value)
	return true
}

func (s *SyncMapExt) SyncMapGet1(keys ...interface{}) (value interface{}, ok bool) {
	var syncMap *SyncMapExt
	syncMap = s
	for _, v := range keys {
		value, ok = syncMap.Load(v)
		switch value.(type) {
		default:
			//fmt.Printf("unexpected type %T", t)
		case *SyncMapExt:
			syncMap = value.(*SyncMapExt)
		}
	}
	return
}

func (s *SyncMapExt) SyncMapSet1(value interface{}, keys ...interface{}) bool {
	mapList := make([]*SyncMapExt, len(keys))
	mapList[0] = s
	var key interface{}
	for i, k := range keys {
		value, ok := s.Load(k)
		if ok {
			switch value.(type) {
			case *SyncMapExt:
				mapList[i+1] = value.(*SyncMapExt)
			}
		}
		key = k
	}

	mapList[len(mapList)-1].Store(key, value)
	return true
}

func SyncTest() {
	var pyram SyncMapExt
	var rootMap *SyncMapExt

	//rootMap.Store("overlay1", "1")
	//rootMap.Store("overlay2", SC{
	//	Name: "1",
	//	Sex:  "2",
	//})
	//pyram.Store("rootcid", &rootMap)

	if k, ok := pyram.Load("rootcid"); ok {
		rootMap = k.(*SyncMapExt)
	} else {
		rootMap = &SyncMapExt{}
	}

	rootMap.Store("1", SC{Name: "账上", Sex: "男"})
	pyram.Store("rootcid", rootMap)

	pyram.SyncMapSet1(SC{Name: "账上", Sex: "男"}, "rootcid", "overlay1")
	value, ok := pyram.SyncMapGet1("rootcid", "1")
	if ok {
		fmt.Println(value)
	}

	//if value, ok := SyncMapGet(&pyram, "rootcid", "overlay2"); ok {
	//	sc := value.(SC)
	//	sc.Name = "张珊"
	//
	//}
	//if ok {
	//	fmt.Println(value)
	//}
	//
	//if v, ok := pyram.Load("rootcid"); ok {
	//	rootMap := v.(*sync.Map)
	//	v, ok := rootMap.Load("overlay2")
	//	if ok {
	//		f := v.(SC)
	//		f.Name = "张散"
	//		rootMap.Store("overlay2", f)
	//	}
	//}

	//SyncMapSet(&pyram, SC{
	//	Name: "账上",
	//	Sex:  "男",
	//}, "rootcid", "overlay2")

	//pyram.Range(func(key, value interface{}) bool {
	//	value.(*SyncMapExt).Range(func(key, value interface{}) bool {
	//		fmt.Println(key, value)
	//		return true
	//	})
	//	return true
	//})
}
