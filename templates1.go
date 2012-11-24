package main

import (
    "os"
    "text/template"
)

type Person struct {
    Name string // 可导出的字段
}

func main() {
    t := template.New("Hello Template") // 创建一个名为 Hello Template 的模板
    t, _ = t.Parse("Hello {{.Name}}") // 分析模板并创建一个模板
    p := Person{Name:"Mary"} // 定义一个实例
    t.Execute(os.Stdout, p) // 转换模板t，填充了p中的数据
}
