package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// INPUT
/*

10

.X..XX...X
X.........
.X.......X
..........
........X.
.X...XXX..
.....X..XX
.....X.X..
..........
.....X..XX
9 1 9 6

output 3
*/


type Pair [2]int

// Complete the minimumMoves function below.
func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {

    startPair := Pair{int(startX), int(startY)}
    endPair := Pair{int(goalX), int(goalY)}
        
    rowLength := len(grid)
    colLength := len(grid[0])
        
    matrix := make([][]rune,len(grid))
    visited := make([][]int,len(grid))

    for i,g := range grid {
        matrix[i] = []rune(g)
        visited[i] = make([]int,len(g))
    }

    queue := make([]Pair,0)
    queue = append(queue, startPair)
    visited[startPair[0]][startPair[1]] = 1


    dr := []int{-1,0,1,0}
    dc := []int{0,1,0,-1}
    
    for len(queue) != 0 {  

        currentPair := queue[0]
        queue = queue[1:]    

        if visited[endPair[0]][endPair[1]] != 0 {
            break
        }

        for i:= 0 ; i < 4 ; i ++ {
            workingPair := currentPair
            counter := visited[currentPair[0]][currentPair[1]]
            for {
                r := workingPair[0]+ dr[i]
                c := workingPair[1]+ dc[i]            
                if  r >= 0 && 
                    r < rowLength && 
                    c >=0 && 
                    c < colLength  && 
                    matrix[r][c] == rune('.') && 
                    visited[r][c] == 0 {
                    visited[r][c] = counter + 1
                    queue = append(queue, Pair{r,c})
                    workingPair[0] = r 
                    workingPair[1] = c
                } else {
                    break
                }
            }

        }
        
    }
    return int32(visited[endPair[0]][endPair[1]]-1)
}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var grid []string

    for i := 0; i < int(n); i++ {
        gridItem := readLine(reader)
        grid = append(grid, gridItem)
    }

    startXStartY := strings.Split(readLine(reader), " ")

    startXTemp, err := strconv.ParseInt(startXStartY[0], 10, 64)
    checkError(err)
    startX := int32(startXTemp)

    startYTemp, err := strconv.ParseInt(startXStartY[1], 10, 64)
    checkError(err)
    startY := int32(startYTemp)

    goalXTemp, err := strconv.ParseInt(startXStartY[2], 10, 64)
    checkError(err)
    goalX := int32(goalXTemp)

    goalYTemp, err := strconv.ParseInt(startXStartY[3], 10, 64)
    checkError(err)
    goalY := int32(goalYTemp)

    result := minimumMoves(grid, startX, startY, goalX, goalY)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
