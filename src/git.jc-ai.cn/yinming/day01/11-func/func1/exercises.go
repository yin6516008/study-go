package func1

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))

	coinsMap = map[string]int{
		"e": 1,
		"E": 1,
		"i": 2,
		"I": 2,
		"o": 3,
		"O": 3,
		"u": 4,
		"U": 4,
	}
)

func integral(name string) int {
	var count int
	for key, coin := range coinsMap {
		count += strings.Count(name, key) * coin
	}
	return count
}

// DispatchCoin is a func
func DispatchCoin() int {
	var counts = make(map[string]int, 20)

	for _, name := range users {
		coin := integral(name)
		coins -= coin
		counts[name] = coin
	}
	fmt.Println(counts)
	return coins
}
