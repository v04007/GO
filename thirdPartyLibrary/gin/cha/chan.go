package main

import (
	"encoding/json"
	"fmt"
)

//type user struct {
//	Name    string
//	Message string
//	Number  int
//}

type name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	//for k, _ := range b {
	//	name.Id = k
	//	name.Name = fmt.Sprintf("%d", k)
	//	b[k] = name
	//
	//	fmt.Println("name", name)
	//	fmt.Println("b[k]", b[k])
	//}

	b := make([]*name, 3)
	b = append(b, &name{
		Id:   0,
		Name: "一",
	})
	b = append(b, &name{
		Id:   1,
		Name: "二",
	})
	b = append(b, &name{
		Id:   2,
		Name: "三",
	})
	var name = new(name)

	for k, v := range b {
		fmt.Println(b[k], k, v)
		fmt.Println("ok", &(b[k]).Id)
		name.Id = k
		name.Name = fmt.Sprintf("%d", k)
		b[k] = name
	}
	marshalb, _ := json.Marshal(b)
	fmt.Println(string(marshalb))

	//r := gin.Default()
	//r.GET("/chan", func(c *gin.Context) {
	//	query := c.Query("data")
	//	c.JSON(http.StatusOK, gin.H{"type": "JSON", "status": 200, "data": query})
	//})
	//r.Run("localhost:8000")
}
