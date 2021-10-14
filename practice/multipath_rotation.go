/*
N            <-人数
s_1
...
s_N
M            <-パス回しの回数
a_1 b_1 x_1
...
a_M b_M x_M  <-誰から、誰に、何個

入力例：
3
10
5
8
3
1 3 5
3 2 3
2 1 10
出力例：
13
0
10

手順：
1.標準入力で受け取る
2.一行目と一行目の数+一行目の数字をAtoiで変換し人数とパス回数を取得
  受け取った文字列が正しいかエラー処理
3.人数分mapを作成する
4.mapにボールの数を記録していく
5.パス回しによるボール個数の変化を処理
6.5をパス回数分whileループする
7.最後にmapを表示していく
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := CreateSliceFromStdin()
	num, path_len, err := GetNumAndPathlen(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	//create map
	maps := make(map[int]int)
	for i := 1; i <= num; i++ {
		maps[i], err = strconv.Atoi(slice[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//pathによるボール数の変化
	for j := 1; j <= path_len; j++ {
		path_inf := strings.Split(slice[num+1+j], " ")
		from, _ := strconv.Atoi(path_inf[0])
		to, _ := strconv.Atoi(path_inf[1])
		many, _ := strconv.Atoi(path_inf[2])
		if maps[from] < many {
			many = maps[from]
		}
		maps[from] -= many
		maps[to] += many
	}

	for i := 1; i <= len(maps); i++ {
		fmt.Println(maps[i])
	}
}

func CreateSliceFromStdin() []string {
	args := bufio.NewScanner(os.Stdin)
	slice := []string{}
	for {
		ret := args.Scan()
		if !ret {
			break
		}
		slice = append(slice, args.Text())
	}
	return slice
}

func GetNumAndPathlen(slice []string) (int, int, error) {
	num, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, err
	}
	path_len, err := strconv.Atoi(slice[num+1])
	if err != nil {
		return 0, 0, err
	}
	return num, path_len, nil
}
