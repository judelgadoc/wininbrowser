package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;

@RestController
@RequestMapping("/folder")
@AllArgsConstructor
public class FolderController {

    private final FolderService folderService;

    @RequestMapping("/new")
    public String newFolder(String diskName,String route, @ModelAttribute Folder folder){
        System.out.println("here");
        ArrayList<Folder> folders = new ArrayList<Folder>();
        folder.setFolders(folders);
        folderService.newFolder(diskName ,route, folder);
        return "";
    }
}
