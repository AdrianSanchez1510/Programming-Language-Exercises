/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package Classes;

/**
 *
 * @author adrian
 */
public abstract class Grupo {

    /**
     * @return the group
     */
    public Grupo[] getGroup() {
        return group;
    }

    /**
     * @param group the group to set
     */
    public void setGroup(Grupo[] group) {
        this.group = group;
    }

    /**
     * @return the texts
     */
    public String[] getTexts() {
        return texts;
    }

    /**
     * @param texts the texts to set
     */
    public void setTexts(String[] texts) {
        this.texts = texts;
    }

    /**
     * @return the object
     */
    public ObjetoGeometrico[] getObject() {
        return object;
    }

    /**
     * @param object the object to set
     */
    public void setObject(ObjetoGeometrico[] object) {
        this.object = object;
    }
    private Grupo[] group;
    private String[] texts;
    private ObjetoGeometrico[] object;
}
