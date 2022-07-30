// Go 为时间（time）和时间段（duration）提供了大量的支持；这儿有是一些例子。
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 从获取当前时间时间开始。
	now := time.Now()
	p(now) // 2022-07-30 22:45:16.5114109 +0800 CST m=+0.005224801

	// 通过提供年月日等信息，你可以构建一个 time。 时间总是与 Location 有关，也就是时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then) // 2009-11-17 20:34:58.651387237 +0000 UTC

	// 你可以提取出时间的各个组成部分。
	p(then.Year())       // 2009
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     // 34
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	// 支持通过 Weekday 输出星期一到星期日。
	p(then.Weekday()) // Tuesday

	// 这些方法用来比较两个时间，分别测试一下是否为之前、之后或者是同一时刻，精确到秒。
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	//方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff) // 111306h10m17.860023663s

	// 我们可以用各种单位来表示时间段的长度。
	p(diff.Hours())       // 111306.17162778435
	p(diff.Minutes())     // 6.678370297667061e+06
	p(diff.Seconds())     // 4.007022178600237e+08
	p(diff.Nanoseconds()) // 400702217860023663

	// 你可以使用 Add 将时间后移一个时间段，或者使用一个 - 来将时间前移一个时间段。
	p(then.Add(diff))  // 2022-07-30 14:45:16.5114109 +0000 UTC
	p(then.Add(-diff)) // 1997-03-08 02:24:40.791363574 +0000 UTC
}

// $ go run time.go
// 输出：
// 2022-07-30 22:45:16.5114109 +0800 CST m=+0.005224801
// 2009-11-17 20:34:58.651387237 +0000 UTC
// 2009
// November
// 17
// 20
// 34
// 58
// 651387237
// UTC
// Tuesday
// true
// false
// false
// 111306h10m17.860023663s
// 111306.17162778435
// 6.678370297667061e+06
// 4.007022178600237e+08
// 400702217860023663
// 2022-07-30 14:45:16.5114109 +0000 UTC
// 1997-03-08 02:24:40.791363574 +0000 UTC
