package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.Data;

import java.util.ArrayList;
import java.util.Objects;

@Data
public class Folder{

    private String name;
    private double size;
    private ArrayList<Folder> folders;
    private ArrayList<File> files;

    public Folder(ArrayList<Folder> folders, ArrayList<File> files, String name) {
        this.folders = folders;
        this.files = files;
        this.size = 0;
        this.name = name;
    }

    public ArrayList<String> getFolderNames(){
        ArrayList<String> names = new ArrayList<>();
        for(int i = 0; i < this.folders.size(); i++){
            names.add(this.folders.get(i).getName());
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

    public void addToFolders(Folder folder){
        this.folders.add(folder);
    }

    public void setFolders(ArrayList<Folder> folders) {
        this.folders = folders;
    }

    public void setFiles(ArrayList<File> files) {
        this.files = files;
    }

    public String getName() {
        return name;
    }
}
