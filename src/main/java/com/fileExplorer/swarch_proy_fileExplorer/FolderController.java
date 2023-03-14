package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/folder")
@AllArgsConstructor
public class FolderController {

    private final FolderService folderService;

    @RequestMapping(value = "/newFolder", method = RequestMethod.PUT)
    public String newFolder(String diskName,String route, @ModelAttribute Folder folder){
        ArrayList<Folder> folders = new ArrayList<Folder>();
        folder.setFolders(folders);
        folderService.newFolder(diskName ,route, folder);
        return "";
    }

    @RequestMapping(value = "/newFile", method = RequestMethod.PUT)
    //@RequestParam("newFile") MultipartFile newFile
    public String newFile(String diskName, String route, String type, String name, double size) throws IOException {
        //System.out.println(newFile);
        folderService.newFile(diskName, route, type, name, size);
        return "";
    }

    @RequestMapping(value = "/getFolders", method = RequestMethod.GET)
    public List<Folder> getFolders(String diskName, String route){
        return folderService.getFolders(diskName, route);
    }

    @RequestMapping(value = "/getFiles", method = RequestMethod.GET)
    public List<File> getFiles(String diskName, String route){
        return folderService.getFiles(diskName, route);
    }
}
