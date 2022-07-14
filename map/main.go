/**
 * @Author: yangyi
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/7/14 上午9:36
 */

package main

import "fmt"

func main()  {
    tom1 := make([]interface{},0)
    tom2 := make([]interface{},0)
    tom3 := make([]interface{},0)

    tom1 = append(tom1,1,2,3,4)
    tom2 = append(tom2,1,2,3,4)
    tom3 = append(tom3,1,2,3,4)

    a := map[string][]interface{}{
        "tom1": tom1,
        "tom2": tom2,
        "tom3": tom3,
    }
    fmt.Println("=====a===",a)

}
