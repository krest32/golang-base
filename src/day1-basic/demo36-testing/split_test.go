// split/split_test.go

package demo36_testing

import (
	"reflect"
	"testing"
)

// 测试函数名必须以Test开头，必须接收一个 *testing. T类型参数
func TestSplit(t *testing.T) {
	// 程序输出的结果
	got := Split("a")
	// 期望的结果
	want := "abbb"
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
