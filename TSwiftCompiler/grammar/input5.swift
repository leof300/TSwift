// // Función recursiva para calcular el factorial de un número
// func factorial(_ n: Int) -> Int {
//     if n < 2 {
//         return 1
//     } else {
//         return n * factorial(n - 1)
//     }
// }
//
// // Función recursiva para calcular la función de Ackermann
// func ackermann( m: Int,  n: Int) -> Int {
//     if m == 0 {
//         return n + 1
//     } else if n == 0 {
//         return ackermann(m - 1, 1)
//     } else {
//         print(m,n)
//         return ackermann(m - 1, ackermann(m, n - 1))
//     }
// }
//
// func main(){
//     print("--------------------------")
//     print("----ARCHIVO RECURSIVOS----")
//     print("--------------------------")
//      print("Factorial de 6: ", factorial(6))
//      print("Factorial de 4: ", factorial(4))
//      print("Factorial de 3: ", factorial(3))
//      print("")
//      print("Ackerman de 3,0: ", ackermann(3, 0))
//     print("Ackerman de 2,8: ", ackermann(2, 8))
//     print("Ackerman de 2,1: ", ackermann(2, 1))
//}

// main()

func func1 () -> Int {
    return 3
}

func suma( num1 x : Int, num2 y: Int) -> Int {
    x+=1
    return x + y
}

func main(){
    var i:Int = 1
    print("************************")
    print(func1 ())
    print(suma(&i,3))
    print(i)
}

main()

/*
func resta(_ x : Int, _ y: Int) -> Int {
    return x - y
}

func mul(x: Int, y: Int) -> Int {
    return x * y
}

func duplicar(_ x: inout Int){
    x += x
}

func ejemplo2(verdura v: Int, v verdura: Int ) {
    print(v)
    print(verdura


    )
}*/