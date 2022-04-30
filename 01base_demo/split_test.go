package base_demo

import (
	"fmt"
	"reflect"
	"testing"
)

/*
@author RandySun
@create 2022-04-30-14:57
*/

//
//  TestSplit
//  @Description:  测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//  @param t
//
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")         // 程序输出结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // slice不能直接比较,借助反射包中方法比较
		t.Errorf("expected: %v, got: %v", want, got) // 测试失败输出错误提示
	}
}

//
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestTypeSplit(t *testing.T) {
	// 定义测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "梦里有肉,我要吃肉", sep: "肉", want: []string{"梦里有", ",我要吃", ""}},
	}
	// 遍历切片, 逐一测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected: %#v, got: %#v", tc.want, got)
		}
	}
}

func TestChildrenSplit(t *testing.T) {
	// 定义测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义存储测试用例的切片
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "梦里有肉,我要吃肉", sep: "肉", want: []string{"梦里有", ",我要吃", ""}},
	}
	// 遍历切片, 逐一测试用例
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
			}
		})

	}
}

//// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("梦里有肉,我要吃肉", "肉")
	}

}

func BenchmarkSplitParallel(b *testing.B) {
	b.SetParallelism(8) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("梦里有肉,我要吃肉", "肉")
		}
	})
}

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func TestSubSetupDownSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	// 定义存储测试用例的切片
	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "梦里有肉,我要吃肉", sep: "肉", want: []string{"梦里有", ",我要吃", ""}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 示例函数示例
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("梦里有肉,我要吃肉", "肉"))
	// Output:
	// [a b c]
	// [梦里有 ,我要吃 ]
}

//func TestAssertSplit(t *testing.T) {
//	got := Split("a:b:c", ":")         // 程序输出结果
//	want := []string{"a", "b", "c"}    // 期望的结果
//	if !reflect.DeepEqual(want, got) { // slice不能直接比较,借助反射包中方法比较
//		t.Errorf("expected: %v, got: %v", want, got) // 测试失败输出错误提示
//	}
//	assert.Equal(t, got, want) // 使用assert提供的断言函数
//
//}
