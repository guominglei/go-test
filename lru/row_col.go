//package main
package lru
import "fmt"

/*
	二维数组遍历
*/

func RightLoop(row_col *[][]int) {

	row :=0
	col := len(row_col[0])
	row_start := 0
	row_end := len(row_col)
	col_start := 0
	col_start := col
	var state = "right"
	for {
		// 右
		if state == "right":
			for i:=row; i <col ; i++{
				value := row_col[row][i]
				if value != -1:
					fmt.Println(value)
					row_col[row][i] = -1
				}else{
					state = 'down'
					col = i-1
					break
				}
			col = i 
			state = "down"
		// 下
		if state == "down"{
				for r:=row+1; r<row_end; r++{
					value = row_col[r][col]
					if value != -1{
						fmt.Println(value)
					}else{
						row_col[r]
					}
				}
			}
		// 左
		if state == "left"{
			
			}
		// 上
		if state == "up"{
			
	}
}

func main_() {

}
