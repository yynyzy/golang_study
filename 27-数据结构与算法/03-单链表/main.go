package main

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string // name
	nickname string
	next     *HeroNode //这个表示指向下一个结点
}

//给链表插入一个结点
//编写第一种插入方法,在单链表的最后加入.
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1．先找到该链表的最后这个结点
	//2．创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next // 让temp不断的指向下一个结点
	}
}
func main() {
}
