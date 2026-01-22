# 📔 Go Quick Notes - Cheatsheet

## 🏗️ Estructura Normal

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

## 💎 Tipos de Datos Primitivos

### 📝 Cadenas de Texto (String)

```go
var mistring string
// String zero value is ""
fmt.Println(mistring)

mistring = "Hola y bienvenido a code and learn"
fmt.Println(mistring)
```

**Características:**
- Zero value: `""`
- Se declaran con `var nombre string`

### 🔢 Números Enteros (Int)

```go
// Entero estándar
var miint int
// Zero value for number types is 0
fmt.Println(miint)
miint = 22
fmt.Println(miint)
```

**Tipos de enteros:**
- **Unsigned (solo positivos):** `uint8, uint16, uint32, uint64`
- **Signed (positivos y negativos):** `int8, int16, int32, int64`
- **int:** Tipo estándar, tamaño depende del SO

**Ejemplo con tipos específicos:**

```go
// uint8: solo números positivos (0-255)
var Valor_positivo_pequenyo uint8
Valor_positivo_pequenyo = 10

// int8: números positivos y negativos (-128 a 127)
var Int_positivo_negativo_pequenyo int8
Int_positivo_negativo_pequenyo = -10

// Número más grande
var myint int = 2400898
fmt.Println(myint)
```

**Conversión de tipos (Type Casting):**

```go
myint = int(Int_positivo_negativo_pequenyo)
// Sintaxis: int(valor)
```

### 🔤 Byte y Rune

```go
// Byte: alias de uint8, representa un carácter ASCII
var mybyte byte = 'A'
fmt.Println(mybyte)  // Imprime: 65 (código ASCII)

// Rune: alias de int32, representa un carácter Unicode
var myrune rune = 'A'
fmt.Println(myrune)  // Imprime: 65

myrune = '😊'
fmt.Println(myrune)  // Imprime: 128522 (código Unicode del emoji)
```

**Diferencias:**
- **byte:** Se usa para datos sin procesar, valor 0-255
- **rune:** Se usa para caracteres Unicode, incluye emojis y caracteres especiales

### ⚖️ Booleanos (Bool)

```go
var miboleano bool
// Zero value for bool is false
fmt.Println(miboleano)

miboleano = true
fmt.Println(miboleano)
```

**Características:**
- Solo dos valores: `true` o `false`
- Zero value: `false`

## 🔑 Conceptos Clave

| Concepto | Descripción |
|----------|-------------|
| **Zero Value** | Valor por defecto cuando se declara una variable sin inicializar |
| **Type Casting** | Conversión de un tipo de dato a otro (ej: `int(variable)`) |
| **Unicode** | Estándar que permite representar caracteres de cualquier idioma |
| **ASCII** | Estándar para representar caracteres en inglés (0-127) |
