package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.Data;

import java.util.ArrayList;

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
}
