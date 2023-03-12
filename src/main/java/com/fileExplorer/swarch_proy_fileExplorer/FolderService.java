package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.AllArgsConstructor;
import org.bson.BsonBinarySubType;
import org.bson.types.Binary;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Optional;

@AllArgsConstructor
@Service
public class FolderService {
    //@Autowired
    private final DiskRepository diskRepository;
    //@Autowired
    private final FolderRepository folderRepository;
    //@Autowired
    private final FileRepository fileRepository;
    //private final MongoTemplate mongoTemplate;
    private final MongoOperations mongoOperations;

    public String newFolder(String diskName, String route, Folder folder){

        ArrayList<String> foldersList = new ArrayList<>(Arrays.asList(route.split("/")));

        Disk disk = diskRepository.findDiskByNamed(diskName);
        Folder folder1 = disk.folderByName(foldersList.get(0));
        //Folder folder2 = folderRepository.findOneByName(disk.folderByName(foldersList.get(0)).getName());
        //System.out.println(folder2);

        for(int i = 1; i< foldersList.size(); i++){
            folder1 = folder1.folderByName(foldersList.get(i));
        }

        if(!folder1.getFolderNames().contains(folder.getName())){
            ArrayList<Folder> folders = new ArrayList<Folder>();
            ArrayList <File> files = new ArrayList<>();
            folder.setFolders(folders);
            folder.setFiles(files);
            folder1.addToFolders(folder);
            folderRepository.save(folder1);
            mongoOperations.save(disk);

        }else{
            throw new IllegalStateException("There's already a folder with this name here");
        }

        return "";
    }

    public String newFile(String diskName, String route, File newFile) throws IOException{
        //File newFile = new File(file);
        //newFile.setFile(new Binary(BsonBinarySubType.BINARY,file.getBytes()));
        System.out.println(newFile);
        //fileRepository.insert(newFile);

        ArrayList<String> foldersList = new ArrayList<>(Arrays.asList(route.split("/")));

        Disk disk = diskRepository.findDiskByNamed(diskName);
        System.out.println("Hello world");
        Folder folder1 = disk.folderByName(foldersList.get(0));

        for(int i = 1; i< foldersList.size(); i++){
            folder1 = folder1.folderByName(foldersList.get(i));
        }

        if(!folder1.getFileNames().contains(newFile.getName())){

            folder1.addToFiles(newFile);
            folderRepository.save(folder1);
            mongoOperations.save(disk);
        }else{
            throw new IllegalStateException("There's already a file with this name here");
        }
        return "";
    }
}
