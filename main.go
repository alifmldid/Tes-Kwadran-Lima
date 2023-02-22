package main

import (
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.POST("/max", max)
	r.POST("/urut", urut)
	r.Run(":8000")
}

type Input struct{
	Bil1 int
	Bil2 int
	Bil3 int
	Bil4 int
	Bil5 int
}

func max(c *gin.Context){
	var input Input
	c.ShouldBindJSON(&input)
	
	var max = 0

	if input.Bil1 > input.Bil2 && input.Bil1 > input.Bil3 && input.Bil1 > input.Bil4 && input.Bil1 > input.Bil5{
		max = input.Bil1
	}else if input.Bil2 > input.Bil3 && input.Bil2 > input.Bil4 && input.Bil2 > input.Bil5{
		max = input.Bil2
	}else if input.Bil3 > input.Bil4 && input.Bil3 > input.Bil5{
		max = input.Bil3
	}else if input.Bil4 > input.Bil5{
		max = input.Bil4
	}else{
		max = input.Bil5
	}

	c.JSON(200, fmt.Sprintf("Angka terbesar: %d", max))
}

func urut(c *gin.Context){
	var input Input
	c.ShouldBindJSON(&input)

	urut := []int{input.Bil1, input.Bil2, input.Bil3, input.Bil4, input.Bil5}

	sort.Ints(urut)

	c.JSON(200, urut)
}