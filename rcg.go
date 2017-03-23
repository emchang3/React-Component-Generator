package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println("世界好！")
  name := flag.String("name", "MyComponent", "Component name.")
  file := flag.String("file", "myComponent.js", "File name.")
  mode := flag.String("mode", "func", "Function or class.")

  flag.Parse()

  props := flag.Args()

  fo, err := os.Create(*file)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err := fo.Close(); err != nil {
      panic(err)
    }
  }()

  w := bufio.NewWriter(fo)

  w.WriteString("import React from 'react';\n\n")

  if *mode == "class" {
    classDec := fmt.Sprintf("class %v extends React.Component {\n", *name)
    w.WriteString(classDec)

    constructor := "  constructor(props) {\n    super(props);\n  }\n\n"
    w.WriteString(constructor)

    render := "  render() {\n    return (\n      <div>\n      </div>\n    );\n  }\n"
    w.WriteString(render)
  } else {
    funcDec := fmt.Sprintf("const %v = ({", *name)
    w.WriteString(funcDec)

    myProps := ""
    for _, v := range props {
      myProps += fmt.Sprintf(" %v,", strings.Split(v, ":")[0])
    }
    myProps = strings.TrimSuffix(myProps, ",")
    myProps = fmt.Sprintf("%v }) => {\n", myProps)
    w.WriteString(myProps)

    ret := "  return (\n    <div>\n    </div>\n  );\n"
    w.WriteString(ret)
  }

  close := "}\n\n"
  w.WriteString(close)

  propTypes := fmt.Sprintf("%v.propTypes = {\n", *name)
  for _, v := range props {
    propType := strings.Split(v, ":")
    propTypes += fmt.Sprintf("  %v: React.PropTypes.%v.isRequired,\n", propType[0], propType[1])
  }
  propTypes = strings.TrimSuffix(propTypes, ",\n")
  propTypes += "\n}\n\n"
  w.WriteString(propTypes)

  export := fmt.Sprintf("export default %v;\n", *name)
  w.WriteString(export)

  w.Flush()
}
