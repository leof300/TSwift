
var bol = true
print("hola ",bol) 
print("hola ",bol) 
print(j) 

// Función recursiva para calcular el factorial de un número
func factorial(_ n: Int) -> Int {
    if n < 2 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}