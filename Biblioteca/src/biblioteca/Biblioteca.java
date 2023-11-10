/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package biblioteca;

import java.util.*;
import Classes.*;
/**
 *
 * @author adrian
 */
public class Biblioteca {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        List<Libro> biblio = new ArrayList<Libro>();
        List<Socio> society = new ArrayList<Socio>(); //I know that Society doesn't make sense but it was funny
        List<Prestamo> pre = new ArrayList<Prestamo>();
        
        biblio.add(new Libro("A01G","Dias de mi lecho de muerte","Adrian Sanchez","Gol"));
        biblio.add(new Libro("G02T","Cien años de Soledad","Gabriel Garcia Marquez","Tom"));
        biblio.add(new Libro("I03G","Hija de la Fortuna","Isabella Allende","Gol"));
        biblio.add(new Libro("Z04Z","Star Wars: Convergence","Zarazoida Cordova","Zar"));
        
        society.add(new Socio(2,"Jordy Rodriguez", "Ka"));
        society.add(new Socio(8,"Alberto Sanchez", "Ra"));
        society.add(new Socio(4,"Kimberly De Jesus", "Lol"));
        
        pre.add(new Prestamo(society.get(2), 4));
        pre.get(0).CrearPrestamo(biblio.get(2));
        pre.get(0).CrearPrestamo(biblio.get(3));
        pre.get(0).CrearPrestamo(biblio.get(1));
        
        society.get(2).seeBooks();
        
    }
}
