package main

import (
	"fmt"
	"reflect"
)

//refValue := reflect.ValueOf(&orange) // value
//1
//则需要首先解引用指针，取得指针指向的对象：
//
//refValue = refValue.Elem()

func SetField(obj interface{}, name string, value interface{}) error {

	// won't work if I remove .Elem()
	fmt.Println(reflect.ValueOf(obj))
	structValue := reflect.ValueOf(obj).Elem()

	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()

	// won't work either if I add .Elem() to the end
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {

		return fmt.Errorf("Provided value %v type %v didn't match obj field type %v", val, val.Type(), structFieldType)
	}

	structFieldValue.Set(val)
	return nil
}

type AA struct {
	Name string
	Agio int
}

func main() {
	A := &AA{
		Name: "ss",
		Agio: 1,
	}

	SetField(A, "Name", "ss")
}

///////////下面是反射查找字段和值，FieldByName
//func (u *User) SayHello() {
//	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
//}
//func caseInsenstiveFieldByName(v reflect.Value, name string) reflect.Value {
//	name = strings.ToLower(name)
//	return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
//}
//func main() {
//	tonydon := &User{"TangXiaodong", 100, "0000123"}
//	object := reflect.ValueOf(tonydon)
//
//	myref := object.Elem()
//	typeOfType := myref.Type()
//	for i := 0; i < myref.NumField(); i++ {
//		field := myref.Field(i)
//		fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
//	}
//	tonydon.SayHello()
//	v := object.MethodByName("SayHello")
//	//fmt.Println("fffffff", reflect.Indirect(object).FieldByName("Name"))
//	fmt.Println("fffffff", caseInsenstiveFieldByName(reflect.Indirect(object), "name"))
//	v.Call([]reflect.Value{})

// fmt.Println("字段是否存在", reflect.Indirect(object).FieldByName("Name").IsValid())
//}

//反射操作切片

//func delete(slice interface{}, i int) {
//	v := reflect.ValueOf(slice).Elem()
//	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
//}
//
//func insert(slice interface{}, i int, val interface{}) {
//	v := reflect.ValueOf(slice).Elem()
//	v.Set(reflect.AppendSlice(v.Slice(0, i+1), v.Slice(i, v.Len())))
//	v.Index(i).Set(reflect.ValueOf(val))
//}
//
//func delete_copy(slice interface{}, i int) {
//	v := reflect.ValueOf(slice).Elem()
//	tmp := reflect.MakeSlice(v.Type(), 0, v.Len()-1)
//	v.Set(
//		reflect.AppendSlice(
//			reflect.AppendSlice(tmp, v.Slice(0, i)),
//			v.Slice(i+1, v.Len())))
//}
//
//func insert_copy(slice interface{}, i int, val interface{}) {
//	v := reflect.ValueOf(slice).Elem()
//	tmp := reflect.MakeSlice(v.Type(), 0, v.Len()+1)
//	v.Set(reflect.AppendSlice(
//		reflect.AppendSlice(tmp, v.Slice(0, i+1)),
//		v.Slice(i, v.Len())))
//	v.Index(i).Set(reflect.ValueOf(val))
//}
//
//func main() {
//	arr := []int{0, 1, 2, 3, 4, 5, 6}[:6]
//	brr := arr
//
//	fmt.Println("arr:", arr, "brr:", brr)
//	insert(&arr, 2, 8)
//	fmt.Println("arr:", arr, "brr:", brr)
//	delete(&arr, 5)
//	fmt.Println("arr:", arr, "brr:", brr)
//
//	fmt.Println("\nCopy Version\n")
//
//	arr = []int{0, 1, 2, 3, 4, 5, 6}[:6]
//	brr = arr
//
//	fmt.Println("arr:", arr, "brr:", brr)
//	insert_copy(&arr, 2, 8)
//	fmt.Println("arr:", arr, "brr:", brr)
//	fmt.Println("brr was unchanged, setting to arr")
//	brr = arr
//	delete_copy(&arr, 5)
//	fmt.Println("arr:", arr, "brr:", brr)
//
//}

//func getNestFile(i interface{}) {
//	objA := reflect.ValueOf(i).Elem()
//	objB := objA.Field(2).Elem()
//	objC := objB.Field(1).Elem()
//	idC := objC.Field(1).Int()
//	fmt.Println(idC)
//}
//
//func main() {
//	c := C{name: "namec", id: 100}
//	b := B{value: "bvalue", c: &c}
//	a := &A{num: 10, level: "high", b: &b}
//
//	getNestFile(a) // 100
//}
