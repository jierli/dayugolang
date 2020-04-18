package main

import "fmt"

//汉诺塔游戏
//n个盘子 start(开始) end(终结) temp(借助)
//n start-> temp-> end
//n-1 start -> end -> temp
//start -> end
//n-1 temp->start->end
//终止条件1 -> start -> end

func tower(start string, end string, temp string, layer int) {
	if layer == 1 {
		fmt.Println(start, "->", end)
	}
	tower(start, temp, end, layer-1)
	fmt.Println(start, "->", end)
	tower(temp, start, end, layer-1)
	fmt.Println(start, "->", end)
}

func main() {
	tower("塔1", "塔2", "塔3", 3)
}
