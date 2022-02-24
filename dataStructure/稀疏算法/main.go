package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//稀疏算法

type ValNode struct {
	row int
	col int
	val interface{}
}
func sparseAlgorithm()  (data []ValNode,err error) {
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2]=1//黑子
	chessMap[2][3]=2//白子
//输出看看原始的数组
	for _, value := range chessMap {
		for _,v := range value {
			fmt.Printf("%d\t",v)
		}
		fmt.Println()
	}
	//转换为稀疏数组:
	//1	遍历chessMap,如果发现有一个元素不为0,创建要给node结构体
	//2    将其放入对应切片即可
	//标值稀疏数组还应该有一个记录元素二维数组的规模(行\列\默认值)
	var saprseArr []ValNode
	node := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	saprseArr = append(saprseArr, node)
	for key, value := range chessMap {
		for k,v := range value {
			if v != 0 {
				node:=ValNode{
					row: key,
					col: k,
					val: v,
				}
				saprseArr=append(saprseArr,node)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("输出稀疏数组\n")
	for i, node := range saprseArr {
		fmt.Printf("%d:%d %d %d\n", i, node.row, node.col, node.val)
	}
	return saprseArr, nil
}

func save(data []ValNode)(err error) {
	file, err := os.OpenFile("./chessMap.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 06666)
	if err != nil {
		fmt.Println("os.OpenFile failed:",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i, node := range data {
		_, err = writer.WriteString(fmt.Sprintf("%d:%d %d %d\n", i, node.row, node.col, node.val))
		if err != nil {
			fmt.Println("writer.WriteString failed",err)
			return
		}
	}
	writer.Flush() //将缓存中的内容写入文件
	return
}
func Open()  {
	file, err := ioutil.ReadFile("./chessMap.txt")
	if err != nil {
		fmt.Println("ioutil.ReadFile failed:",err)
		return
	}

	fmt.Println("存储文件中读取:")
	fmt.Println(string(file))
}
func main()  {
	saprseArr, err := sparseAlgorithm()
	if err != nil {
		fmt.Println("sparseAlgorithm failed",err)
		return
	}
	err = save(saprseArr)
	if err != nil {
		fmt.Println("save failed",err)
		return
	}
	Open()
}