package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.AllArgsConstructor;
import org.springframework.data.mongodb.MongoDatabaseFactory;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/disk")
@AllArgsConstructor
public class DiskController {
    private final DiskService diskService;

    @GetMapping(value="/all")
    public List<Disk> fetchAllDisks(){
        return diskService.getAllDisks();
    }

    @RequestMapping(value="/new", method = RequestMethod.POST)
    public String newDisk(@ModelAttribute Disk disk) {
        ArrayList<Folder> folders = new ArrayList<Folder>();
        disk.setFolders(folders);
        diskService.newDisk(disk.getName(), disk);
        return "Inserted";
    }

    @RequestMapping(value="/newFolder", method=RequestMethod.POST)
    public String newFolder(String diskName,@ModelAttribute Folder folder){
        //Folder folder = new Folder(new ArrayList<Folder>(),new ArrayList<File>(),name);
        folder.setFiles(new ArrayList<File>());
        folder.setFolders(new ArrayList<Folder>());
        diskService.newFolder(diskName, folder );
        return "";
    }


}
