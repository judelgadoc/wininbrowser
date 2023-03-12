package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

@Data
@Document
public class Disk {
    @Id
    private String id;

    @Indexed(unique = true)
    private String name;
    private int maximumSize;

    private ArrayList<Folder> folders;
    private double size;

    public Disk(int maximumSize, String name, ArrayList<Folder> folders) {
        //this.disk_id = disk_id;
        this.maximumSize = maximumSize;
        this.name = name;
        this.folders= folders;
    }

    public ArrayList<String> getFolderNames(){
        ArrayList<String> names = new ArrayList<>();
        for (Folder folder : this.folders) {
            names.add(folder.getName());
        }
        return names;
    }

    public Folder folderByName(String name){
        for (Folder folder : this.folders) {
            if (Objects.equals(folder.getName(), name)) {
                return folder;
            }
        }
        return null;
    }

    public void setFolders(ArrayList<Folder> folders) {
        this.folders = folders;
    }

    public void addToFolders(Folder folder){
        this.folders.add(folder);
    }

    public String getName() {
        return name;
    }

    public ArrayList<Folder> getFolders() {return folders;}

}
