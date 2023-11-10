/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Classes;

import java.util.*;
/**
 *
 * @author adrian
 */
public class Prestamo {
    
    private Socio newSocio;
    private List<Libro> newBook = new ArrayList<Libro>();
    private int code;
    
    public Prestamo(Socio a, int b)
    {
        this.newSocio = a;
        this.code = b;
    }
    
    /**
     * @return the newSocio
     */
    public Socio getNewSocio() {
        return newSocio;
    }

    /**
     * @param newSocio the newSocio to set
     */
    public void setNewSocio(Socio newSocio) {
        this.newSocio = newSocio;
    }

    /**
     * @return the newBook
     */
    public List<Libro> getNewBook() {
        return newBook;
    }

    /**
     * @param newBook the newBook to set
     */
    public void addNewBook(Libro e) {
        newBook.add(e);
        newSocio.addBook(e);
    }

    /**
     * @return the code
     */
    public int getCode() {
        return code;
    }

    /**
     * @param code the code to set
     */
    public void setCode(int code) {
        this.code = code;
    }
    
    public void CrearPrestamo(Libro e)
    {
        if(e.isAvailability())
        {
            e.setAvailability(false);
            this.newSocio.addBook(e);
            return;
        }
        System.out.println("This book is already owned");
    }
}
