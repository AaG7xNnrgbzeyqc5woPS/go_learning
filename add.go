package main

import ("fmt"
        "math"
        "math/rand"
)

//------------------------------------------------------------------
//打印宽度控制，每个数字的宽度一致，表格对齐
func w_num(i int) int {
  n := int(math.Log10(float64(i))) + 1
  return n
}

//-----------------------------------------------------------------
//输出一个算式的“格式化串”
func s_fmt(i,j int) string {
  var sfmt string
  
 //wi, wj, ws 表示max_i,max+j以及他们之和的数字位数
  //用于控制输出宽度，让表格更整齐
  wi := w_num(i)
  wj := w_num(j)
  ws := w_num(i+j)

  //fmt.Println("被加数,加数,和的最大数的位数：", wi, " ", wj," ", ws)

  sfmt =  fmt.Sprintf("%%%dd + %%%dd = %%%dd     ", wi,wj,ws)
  //fmt.Println("动态格式化串",sfmt)

  return sfmt
}


//-------------------------------------------------------------------
/*
换行规则：
0. 定义：逻辑行，逻辑列指加法表逻辑上的行和列数，比如9×9加法表，99×99加法表，这里的第一个数字是逻辑行，第二个数字是逻辑列。
   很明显，99×99这样的大表格无法简单的在屏幕上简单的输出，一个逻辑行要在屏幕上输出多个屏幕行。
1. 每行输出算式有个最大值，这里假定为10,达到这个最大值，换行
2. 条件1 不满足时，如果逻辑行输出结束，也换行
3. 如果一个逻辑行多于一个输出行， 逻辑行之间有一个换行符号

*/

//打印 逻辑行间回车，已经逻辑行结束的时候的回车
//col: 当前列号
//col_max: 打印时候，一行最大列数
//total_col: 逻辑行的总列数

func is_print_ret(col, col_max, total_col   int) (ret bool) {
  ret = false
  if col % col_max == 0 {  //整除一行最大的列数的时候换行
    ret = true
  } else {
    if col == total_col {  //不整除只有到逻辑行结束才换行
      ret = true
    }
  }
  return ret
}


//-----------------------------------------------------------
// 逻辑行数大于显示列数，说明是一个逻辑行分多行显示
// 这时希望加一个空行区分逻辑行
func is_add_ret(col_max, total_col int) bool {
  if total_col > col_max  {
    return true
  }
  return false

}


//------------------------------------------------------------
func add_list(max_i, max_j int, ){
  const maxnum_perline = 10
  
  var sfmt string
  sfmt = s_fmt(max_i, max_j)

  for i := 1; i <= max_i; i++ {
    for j := 1; j <= max_j; j++ {
      //s := fmt.Sprintf("%1d + %1d = %2d   ",i,j,i+j)
      s := fmt.Sprintf(sfmt, i, j, i+j )
      fmt.Print(s)
      if is_print_ret(j, maxnum_perline, max_j) {
         fmt.Println()
      }
    }

    if is_add_ret(maxnum_perline, max_j) {
      fmt.Println()
    }
  }
  
}

//-------------------------------------------------

func add_list_9() {
  fmt.Println("9×9 以内的加法口诀表")
  add_list(9,9)
}

func add_list_19() {
  fmt.Println("19×19以内的加法口诀表")
  add_list(19,19)
}


func add_list_99() {
  fmt.Println("99×99以内的加法口诀表")
  add_list(99,99)
}

func add_list_8_23() {
  fmt.Println("8×23以内的加法口诀表")
  add_list(8,23)
}

//=========================================================
//随机加法题
func add_rand(max int) {
  var a,b int
  a = rand.Intn(100)
  b = rand.Intn(100)
  fmt.Println(a,"+",b,"=",a+b) 
}

//--------------------------------------------------
func main() {
  add_list_9()
  println()
  //println("------------------------------------------------------------------------------------")
  //add_list_19()
  //println()
  //println("------------------------------------------------------------------------------------")
  //add_list_99()

  //println("-------------------------------------------------------------------------------------")
  //add_list_8_23()

  add_rand(100)
}
