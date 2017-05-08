package main

import (
"bufio"
"fmt"
"os"
"regexp"
"strconv"
"strings"
)

type Arbol struct {
izq   *Arbol
valor string
der   *Arbol
}

type Pila struct {
dato *Arbol
sig  *Pila
}

func push(p *Pila, a *Arbol) *Pila {
if p != nil {
nuevo := &Pila{a, p}
p = nuevo
} else {
p = &Pila{a, nil}
}
return p
}

func pop(p *Pila) (*Arbol, *Pila) {
salida := p
p = p.sig
return salida.dato, p
}

func expToArbol(p *Pila, s []string) *Pila {
cant := 0
var arbAux *Arbol
for _, dato := range s {
var a = &Arbol{valor: dato}
if dato != "+" && dato != "-" && dato != "*" && dato != "/" && dato!=":="{
p = push(p, a)
cant++
} else {
if cant < 2 {
fmt.Println("Expresion no valida")
return nil
} else {
arbAux, p = pop(p)
a.der = arbAux
arbAux, p = pop(p)
a.izq = arbAux
p = push(p, a)
cant--
}
}
}
if cant != 1 {
fmt.Println("Expresion no valida")
return nil
}
return p
}

func recorrerInorden(t *Arbol) {
if t == nil {
return
}
recorrerInorden(t.izq)
fmt.Print(t.valor, " ")
recorrerInorden(t.der)
}

func recorrerPreorder(t *Arbol){
  if t == nil {
    return
  }

  fmt.Print(tipoExpresion(t.valor))
  fmt.Println(t.valor)
  if t.der!=nil && t.izq!=nil{
    fmt.Print("Izquierda :")
    recorrerInorden(t.izq)
    fmt.Print("Derecha :")
    recorrerInorden(t.der)
    fmt.Println("")
    fmt.Println("--------------------------------------")
    recorrerPreorder(t.izq)
    recorrerPreorder(t.der)
  }
}

func calcular(t *Arbol) int {
if t != nil {
switch t.valor {
case "*":
return calcular(t.izq) * calcular(t.der)
case "/":
return calcular(t.izq) / calcular(t.der)
case "+":
return calcular(t.izq) + calcular(t.der)
case "-":
return calcular(t.izq) - calcular(t.der)
default:
i, _ := strconv.Atoi(t.valor)
return i
}
}
return -1
}

func evaluar(t *Arbol) {
recorrerInorden(t)
fmt.Println("=", calcular(t))
}

func comprobar(t *Arbol, er *string) int {
if t != nil {
_, e := strconv.Atoi(t.valor)
if t.der == nil && t.izq == nil {
if e != nil {
*er += t.valor + " debe ser un valor numerico\n"
return 1
}
return 0
} else if t.der != nil && t.izq != nil {
if e != nil {
if t.valor == "*" || t.valor == "/" || t.valor == "+" || t.valor == "-" {
return comprobar(t.der, er) + comprobar(t.izq, er)
}
*er += t.valor + " operacion no valida\n"
return comprobar(t.der, er) + comprobar(t.izq, er) + 1
}
*er += t.valor + " debe ser una operacion, no un numero\n"
return comprobar(t.der, er) + comprobar(t.izq, er) + 1
} else {
return 1
}
}
return 0
}
func prettyMatches(m []string) string {
s := "["
for i, e := range m {
s += e
if i < len(m)-1 {
s += "|"
}
}
s += "]"
return s
}

func prettySubmatches(m [][]string) string {
s := "[\n"
for _, e := range m {
s += "    " + prettyMatches(e) + "\n"
}
s += "]"
return s
}

var (
operadoresLog = []string{"&", "|", "!"}
text          = `& | !`
)

var (
comparativos = []string{"==", "!=", "<", ">", ">=", "<="}
text2        = `& `
)

var (
expresionesLog = []string{"true", "false"}
text3          = `true false `
)

var (
expresionesAsig = []string{":="}
)

var (
expresionesOperadores = []string{"\\+", "\\-", "\\*", "\\/"}
)

var (
expresionesConstantes = []string{"0", "1", "2", "3", "4", "5"}
)

var (
expresionesVariables = []string{"x"}
)

func main() {

var p *Pila
var a *Arbol
//var e string

sc := bufio.NewScanner(os.Stdin)
//var cadena string
fmt.Println("Escriba la expresion postfija(cada termino separado por espacio)")
sc.Scan()
cadena := sc.Text()
//cadena=sc.Text()
//var cadena string = "4 + 5 - 2 - 3 := x"
//op := "4 5 + 3 2 - +"
p = expToArbol(p, strings.Split(cadena, " "))
if p != nil {
  a, p = pop(p)
  recorrerPreorder(a)
//
//b := comprobar(a, &e)
//if b == 0 {
//  fmt.Println("Jose es lindo")
//evaluar(a)
//}
}

//cadena := op

}


func tipoExpresion(cadena string) string  {

  for _, e := range operadoresLog {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena){
      return "operador logico: "
    }
  }

  for _, e := range comparativos {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena) {
      return "operador comparativo: "
    }
  }

  for _, e := range expresionesLog {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena){
      return "operador logico: "
    }
  }

  for _, e := range expresionesAsig {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena) {
      return "operador asignación: "
    }
  }

  for _, e := range expresionesOperadores {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena) {
      return "operador operacion: "
    }
  }

  for _, e := range expresionesConstantes {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena) {
      return "operador constante: "
    }
  }

  for _, e := range expresionesVariables {
    re := regexp.MustCompile(e)
    if cadena == re.FindString(cadena) {
      return "operador variable: "
    }
  }

  return ""
}
