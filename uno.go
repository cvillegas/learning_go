package main 

import "fmt"

func main(){
    var mensaje string = "Hola mundo" 
    mensajeFacil := "Hola mundanos usando :=" 
    fmt.Println(mensaje)
    fmt.Println(mensajeFacil)

    //Comentarios
    a := 10.
    var b float64 = 3
    fmt.Println(( a / b))
    var c int = 10 
    d := 3 
    fmt.Println(c / d)
    x := true
    y := false 
    fmt.Println((x || y))
    fmt.Println((x && y))
    fmt.Println((!x))

}

func privada() {
    fmt.Println("funcion que no necesita ser exportada")
}

func Publica() {
    fmt.Println("funcion que puedo exportar")
}

func imprimirHola() {
    defer fmt.Println("Mundo")
    fmt.Println("Hola")
}
