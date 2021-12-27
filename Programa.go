package main 

import "fmt"	
import "time"

func main(){
	
	fmt.Println("****** MI PROGRAMA CON GO *****")


	/*
	fmt.Println("Hola" +os.Args[1] +" Bienvenido al pgorama en GO ") 

	edad,_ := strconv.Atoi(os.Args[2])
	

	if edad >= 18  && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad")
	}else if edad == 25 {
		fmt.Println("tienes 25 años")
	}else if edad == 99 {
		fmt.Println("Moriras pronto")
	}else{
		fmt.Println("Eres menor de edad")
	}
	*/


	/*
	numero,_ := strconv.Atoi(os.Args[1])
	if numero%2 == 0{

		fmt.Println("El numero es par")
	}else{
		fmt.Println("El numero es impar")
	} 
	*/


	
	//peliculas := [] string{"Pelicula1", "El club de la lucha", "A todo Gas", "El gran torino"}


	/*	
	for i := 0; i < len(peliculas); i ++ {
	if i%2 == 0{
	fmt.Println("La pelicula --> "+peliculas [i] +"es par:",  i)
	}else{
		fmt.Println("La pelicula --> "+peliculas [i] +"es impar:",  i)
	} 
	}
	*/

	/*
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
	*/

	momento := time.Now();
	hoy := momento.Weekday();

	switch hoy{
	case 0:
		fmt.Println("hoy es domingo")
	case 1:
		fmt.Println("hoy es lunes")
	case 2:
		fmt.Println("hoy es martes")
	case 3:
		fmt.Println("hoy es miercoles")
	default: 
		fmt.Println("hoy es otro dia de la semana")
	}
}