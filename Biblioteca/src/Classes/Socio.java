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
public class Socio {

    /**
     * @return the pr
     */
    public List<Prestamo> getPr() {
        return pr;
    }

    /**
     * @param pr the pr to set
     */
    public void setPr(Prestamo e) {
        this.pr.add(e);
    }
    
    private int number;
    private String name;
    private String address;
    private List<Libro> book = new ArrayList<Libro>();
    private List<Prestamo> pr = new ArrayList<Prestamo>();
    
    public Socio(int number, String name, String address)
    {
        this.number = number;
        this.name = name;
        this.address = address;
    }
    
    /**
     * @return the number
     */
    public int getNumber() {
        return number;
    }

    /**
     * @param number the number to set
     */
    public void setNumber(int number) {
        this.number = number;
    }

    /**
     * @return the name
     */
    public String getName() {
        return name;
    }

    /**
     * @param name the name to set
     */
    public void setName(String name) {
        this.name = name;
    }

    /**
     * @return the address
     */
    public String getAddress() {
        return address;
    }

    /**
     * @param address the address to set
     */
    public void setAddress(String address) {
        this.address = address;
    }

    /**
     * @return the book
     */
    public List<Libro> getBook() {
        return book;
    }

    /**
     * @param book the book to set
     */
    public void addBook(Libro book) {
        this.book.add(book);
    }
    
    public void seeBooks()
    {
        if(this.book.size() > 2)
        {
            System.out.println(this.name + " has the following books: ");
            
            for(int i = 0; i < this.book.size(); i++)
            {
                System.out.println((i+1) + ". " + this.book.get(i).getTitle());
            }
            return;
        }
        System.out.println("Not enough Amount of books for " + this.name);
    }
    
    
}
