package gotest

import "testing"

func Test_Divistion_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没有通过")
	} else {
		t.Log("第一个测试通过了")
	}
}

func Test_Divistion_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("Division did not work as expected.")
	} else {
		t.Log("one test passed.", e)
	}
}

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func BenchmarkTimeConsumingFunction(b *testing.B) {
	// 调用该函数停止压力测试的时间技术
	b.StopTimer()

	// 做一些初始化的工作，例如读取文件数据， 数据库连接之类的
	// 这样这些时间不影响我们测试函数本身的性能

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
