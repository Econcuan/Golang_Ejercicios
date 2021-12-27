package main

import "fmt"
import "io/ioutil"
import "os"

func main(){

	fmt.Println("Lector:");

	nuevo_texto := os.Args[1]+"\n"
	//escribir := ioutil.WriteFile("Fechero.txt", nuevo_texto, 0755)

	fichero, err := os.OpenFile("fechero.txt", os.O_APPEND, 0777)
	showError(err)	

	escribir, err := fichero.WriteString(nuevo_texto);
	fmt.Println(escribir)
	showError(err)

	fichero.Close();

	texto, errorDeFichero := ioutil.ReadFile("fechero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))

}

func showError(e error){

	if (e != nil){
		panic(e)
	}
}