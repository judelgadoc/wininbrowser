package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.AllArgsConstructor;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Optional;

@AllArgsConstructor
@Service
public class FolderService {
    private final DiskRepository diskRepository;
    private final FolderRepository folderRepository;
    //private final MongoTemplate mongoTemplate;
    private final MongoOperations mongoOperations;

    public String newFolder(String diskName, String route, Folder folder){

        System.out.println(route);
        ArrayList<String> foldersList = new ArrayList<>(Arrays.asList(route.split("/")));
        System.out.println(foldersList);

        Disk disk = diskRepository.findDiskByNamed(diskName);
        System.out.println(disk.folderByName(foldersList.get(0)));
        Folder folder1 = disk.folderByName(foldersList.get(0));
        //Folder folder2 = folderRepository.findOneByName(disk.folderByName(foldersList.get(0)).getName());
        //System.out.println(folder2);

        for(int i = 1; i< foldersList.size(); i++){
            folder1 = folder1.folderByName(foldersList.get(i));
        }

        if(!folder1.getFolderNames().contains(folder.getName())){
            ArrayList<Folder> folders = new ArrayList<Folder>();
            folder.setFolders(folders);
            folder1.addToFolders(folder);
            //System.out.println(mongoOperations.save(folder1));
            folderRepository.save(folder1);
            mongoOperations.save(disk);

        }else{
            throw new IllegalStateException("There's already a folder with this name here");
        }

        return "";
    }
}
