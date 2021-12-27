package main

import "fmt"

type gorra struct{

	marca string 
	color string 
	precio float32
	plana bool

}


func main() {

/*
	var gorra_negra = gorra { "nike", "negro", 25.50, false}

		fmt.Println(gorra_negra.marca)
		fmt.Println(gorra_negra.precio)

*/

/*
	var numero1 float32 = 10
	var numero2 float32 = 6

	fmt.Println("calculadora1");
	calculadora (numero1, numero2);


	fmt.Println("-------------------------")
	var numero3 float32 = 44;
	var numero4 float32 = 7;
	fmt.Println("calculadora2");
	calculadora(numero3, numero4);
*/

	/*
	fmt.Println("pedido1----->")
	fmt.Println (gorras(45, "Q"));

	fmt.Println("------------------------")

	fmt.Println("pedido2------>")
	fmt.Println (gorras(24, "$"));
 	*/

 	
 	/*
 	var peliculas [3][2] string;
 	peliculas[0][0] = "La verdad duele"
 	peliculas[0][1] = "Mientras duermes"

 	peliculas[1][0] = "Ciudadano ejemplar"
	peliculas[1][1] = "El señor de los anillos"

 	peliculas[2][0] = "Gran torino"
 	peliculas[2][1] = "A todo Gas"
 	*/

 	
 	peliculas := [] string{
 		"la verdad duele", 
 		"Ciudadano Ejemplar", 
 		"Batman",
 		"Super Man"}
 	
 		peliculas = append(peliculas, "Sin limites");
 		peliculas = append(peliculas, "Sin miedo");
 		peliculas = append(peliculas, "A todo gas");


 	fmt.Println(peliculas[0:3])
 	


 	}

 	 	func pantalon(caracteriscas ...string){

 	 		for _, caracteriscas := range caracteriscas{
 	 			fmt.Println(caracteriscas);
 	 			
 	 		}

}	

	func gorras(pedido float32, modena string) (string, float32, string){

		precio := func()float32{
			return pedido * 7

		}

		return "El precio del pedido:", precio(), modena
	}



	func devolverTexto () (dato1 string, dato2 string){

		dato1 = "Wilmer"
		dato2 = "Gomez"

		return dato1, dato2	



	}



	func holaMundo(){

		fmt.Println("Hola Mundo!!");

	}
	
	func operacion(n1 float32, n2 float32, op string) float32 {
		var resultado float32;
		if (op =="+"){
			resultado = n1 + n2
		} 

		if (op =="-"){
			resultado = n1 - n2
		}

		if (op =="*"){
			resultado = n1 * n2
		}

		if (op =="/"){
			resultado = n1 / n2
		}

		return resultado

	}	

	func calculadora (numero1 float32, numero2 float32){

	//suma
	fmt.Print("la suma es: " ) 
	fmt.Println(operacion(numero1, numero2, "+"));
	
	//resta
	fmt.Print("la resta es: " ) 
	fmt.Println(operacion(numero1, numero2, "-"));

	//multiplicacion
	fmt.Print("la multiplicacion es: " ) 
	fmt.Println(operacion(numero1, numero2, "*"));

	//division
	fmt.Print("la division es: " ) 
	fmt.Println(operacion(numero1, numero2, "/"));


	}




