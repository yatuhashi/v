package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func allShow(filename string) {
	/******************************************************
	  指定されたファイルを開き、全て標準出力する
	  *******************************************************/
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// エラー処理
		color.Red("ファイルが開けませんでした")
		return
	}
	fmt.Print(string(file))
}

func lineShow(filename string, start, end int, line bool) {
	/******************************************************
	  指定されたファイルを開き条件に応じて標準出力する
	  os.Stdout.Write([]byte("\033[38;5;39m"+sc.Text()+"\033[0m \n"))
	  39mのところをcolor.goより得れる値に変えてやると自分好みの色に変更出来る
	  *******************************************************/
	file, err := os.Open(filename)
	if err != nil {
		// エラー処理
		color.Red("ファイルが開けませんでした")
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	//行番号を表示するかどうかで回すものを変える
	if line == false {
		yellow := color.New(color.FgYellow).SprintFunc()
		for i := 1; sc.Scan(); i++ {
			if start > i || end < i {
				continue
			}
			if err := sc.Err(); err != nil {
				// エラー処理
				break
			}
			//fmt.Printf("%5s| %s\n",yellow(space(i)), sc.Text())
			fmt.Printf("%5s| ", yellow(space(i)))
			os.Stdout.Write([]byte("\033[38;5;39m" + sc.Text() + "\033[0m \n"))
		}
	} else {
		for i := 1; sc.Scan(); i++ {
			if start > i || end < i {
				continue
			}
			if err := sc.Err(); err != nil {
				// エラー処理
				break
			}
			//fmt.Printf("%s\n",sc.Text())
			os.Stdout.Write([]byte("\033[38;5;39m" + sc.Text() + "\033[0m \n"))
		}
	}
}
func space(i int) string {
	/*************************************
	  行番号の桁に合わせて空白文字数を変える
	  **************************************/
	ori := strconv.Itoa(i)
	return strings.Repeat(" ", (6-len(ori))) + ori
}

func insertString(filename, str string, col, row int, line, rp bool) {
	/******************************************************
	  指定されたファイルを開き条件に応じて標準出力する
	  os.Stdout.Write([]byte("\033[38;5;39m"+sc.Text()+"\033[0m \n"))
	  39mのところをcolor.goより得れる値に変えてやると自分好みの色に変更出来る
	  指定されたcol番号の所で文字処理を行い、文字列を挿入する
	  *******************************************************/
	file, err := os.Open(filename)
	if err != nil {
		// エラー処理
		color.Red("ファイルが開けませんでした")
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	//行番号を表示するかどうかで回すものを変える
	red := color.New(color.FgRed).SprintFunc()
	if line == false {
		yellow := color.New(color.FgYellow).SprintFunc()
		for i := 1; sc.Scan(); i++ {
			if col-5 > i || col+5 < i {
				continue
			}
			if err := sc.Err(); err != nil {
				// エラー処理
				break
			}
			if i == col {
				temp := sc.Text()
				fmt.Printf("%5s| ", yellow(space(i)))
				fmt.Printf(temp[:row])
				fmt.Printf("%s", red(str))
				if rp {
					fmt.Printf(temp[row+len(str):])
				} else {
					fmt.Printf(temp[row:])
				}
				fmt.Printf("\n")
				continue
			}
			fmt.Printf("%5s| ", yellow(space(i)))
			os.Stdout.Write([]byte("\033[38;5;39m" + sc.Text() + "\033[0m \n"))
		}
	} else {
		for i := 1; sc.Scan(); i++ {
			if col-5 > i || col+5 < i {
				continue
			}
			if err := sc.Err(); err != nil {
				// エラー処理
				break
			}
			if i == col {
				temp := sc.Text()
				fmt.Printf(temp[:row])
				fmt.Printf("%s", red(str))
				if rp {
					fmt.Printf(temp[row+len(str):])
				} else {
					fmt.Printf(temp[row:])
				}
				fmt.Printf("\n")
				continue
			}
			//fmt.Printf("%s\n",sc.Text())
			os.Stdout.Write([]byte("\033[38;5;39m" + sc.Text() + "\033[0m \n"))
		}
	}

}

/*
##表示機能
  - 全部表示
  - start,endを指定し表示
##カット機能
  - start,endを指定しカット表示
##追加機能
  - 指定した行数のあとに追加する
*/

func main() {
	filename := flag.String("f", "false", "help message for filename(string)")
	insert := flag.String("i", "", "help message for insert strings(string)")
	all := flag.Bool("a", false, "help message for all show(bool)")
	line := flag.Bool("l", false, "help message for without show line(bool)")
	lnum := flag.Bool("ln", false, "help message for without show line(bool)")
	rp := flag.Bool("rp", false, "help message for without show line(bool)")
	col := flag.Int("c", 0, "help message for insert col num(int)")
	row := flag.Int("r", 0, "help message for insert row num(int)")
	start := flag.Int("s", 0, "help message for line start num(int)")
	end := flag.Int("e", 65, "help message for line end num(int)")
	flag.Parse()

	if *filename == "false" {
		if len(flag.Args()) < 1 {
			color.Red("ファイルが開けませんでした")
			return
		}
		*filename = flag.Args()[0]
	}
	if *line && !*lnum {
		color.Green("0123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789\n")
	} else if !*lnum {
		color.Green("        0123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789\n")
	}
	if len(*insert) >= 1 {
		insertString(*filename, *insert, *col, *row, *line, *rp)
		return
	}
	if *all {
		allShow(*filename)
		return
	}
	lineShow(*filename, *start, *end, *line)
}
