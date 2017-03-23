package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "strings"
)

/*
To-Do:
- Convert props to a flag.
- Accommodate multiple imports.
*/

func writeClass(w *bufio.Writer, name string) {
  classDec := fmt.Sprintf("class %v extends React.Component {\n", name)
  w.WriteString(classDec)

  constructor := "  constructor(props) {\n    super(props);\n  }\n\n"
  w.WriteString(constructor)

  render := "  render() {\n    return (\n      <div>\n      </div>\n    );\n  }\n"
  w.WriteString(render)
}

func writeFunc(w *bufio.Writer, name string, props []string) {
  if props[0] != "" {
    funcDec := fmt.Sprintf("const %v = ({", name)
    w.WriteString(funcDec)

    myProps := ""
    for _, v := range props {
      myProps += fmt.Sprintf(" %v,", strings.Split(v, ":")[0])
    }
    myProps = strings.TrimSuffix(myProps, ",")
    myProps = fmt.Sprintf("%v }) => {\n", myProps)
    w.WriteString(myProps)
  } else {
    funcDec := fmt.Sprintf("const %v = () => {\n", name)
    w.WriteString(funcDec)
  }

  ret := "  return (\n    <div>\n    </div>\n  );\n"
  w.WriteString(ret)
}

func writeProptypes(w *bufio.Writer, name string, props []string)  {
  if props[0] != "" {
    propTypes := fmt.Sprintf("%v.propTypes = {\n", name)
    for _, v := range props {
      propType := strings.Split(v, ":")
      propTypes += fmt.Sprintf("  %v: React.PropTypes.%v.isRequired,\n", propType[0], propType[1])
    }
    propTypes = strings.TrimSuffix(propTypes, ",\n")
    propTypes += "\n}\n\n"
    w.WriteString(propTypes)
  }
}

func main() {
  fmt.Println("世界好！")
  name := flag.String("name", "MyComponent", "Component name.")
  file := flag.String("file", "myComponent.js", "File name.")
  mode := flag.String("mode", "func", "Function or class.")
  proptypes := flag.String("props", "", "prop:type,prop:type")

  flag.Parse()

  props := strings.Split(*proptypes, ",")

  fo, err := os.Create(*file)
  if err != nil {
    panic(err)
  }

  // Why panic at the end? Who knows! I'm cargo-culting, here.
  defer func() {
    if err := fo.Close(); err != nil {
      panic(err)
    }
  }()

  w := bufio.NewWriter(fo)

  imports := "import React from 'react';\n\n"
  w.WriteString(imports)

  if *mode == "class" {
    writeClass(w, *name)
  } else {
    writeFunc(w, *name, props)
  }

  close := "}\n\n"
  w.WriteString(close)

  writeProptypes(w, *name, props)

  export := fmt.Sprintf("export default %v;\n", *name)
  w.WriteString(export)

  w.Flush()
}
