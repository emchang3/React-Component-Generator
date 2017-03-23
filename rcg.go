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
- Figure out how to handle nested proptypes (e.g.: shape).
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

func writeProptypes(w *bufio.Writer, name string, props []string) {
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

func writeImportStatements(w *bufio.Writer, importStatements []string)  {
  for _, v := range importStatements {
    importStatement := strings.Split(v, ":")
    myImport := fmt.Sprintf("import %v from '%v';\n", importStatement[0], importStatement[1])
    w.WriteString(myImport)
  }
  w.WriteString("\n")
}

func main() {
  fmt.Println("世界好！")
  file := flag.String("file", "myComponent.js", "File name.")
  imports := flag.String("imports", "React:react", "Imports.")
  mode := flag.String("mode", "func", "Function or class.")
  name := flag.String("name", "MyComponent", "Component name.")
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

  importStatements := strings.Split(*imports, ",")
  writeImportStatements(w, importStatements)

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
