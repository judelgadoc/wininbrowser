package com.fileExplorer.swarch_proy_fileExplorer;

import com.mongodb.client.result.UpdateResult;
import lombok.AllArgsConstructor;
//import lombok.var;
import org.springframework.boot.context.config.ConfigDataNotFoundException;
import org.springframework.boot.context.config.ConfigDataResourceNotFoundException;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.data.mongodb.core.query.UpdateDefinition;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@AllArgsConstructor
@Service
public class DiskService {

    private final DiskRepository diskRepository;
    //private final MongoTemplate mongoTemplate;
    private final MongoOperations mongoOperations;

    public List<Disk> getAllDisks() {
        return diskRepository.findAll();
    }

    public String newDisk(String name, Disk disk){
        diskRepository.findDiskByName(name)
            .ifPresentOrElse(s -> {
                System.out.println("There's already a disk with this name");
                throw new IllegalStateException("There's already a disk with this name");
            },()->{
                diskRepository.insert(disk);

            });
        return "";
    }


    public String newFolder(String diskName, Folder folder){
        Query query = new Query();
        query.addCriteria(Criteria.where("name").is(diskName));
        Disk disk = mongoOperations.findOne(query, Disk.class);
        if(!disk.getFolderNames().contains(folder.getName())){
            ArrayList<Folder> folders = new ArrayList<Folder>();
            folder.setFolders(folders);
            disk.addToFolders(folder);
            mongoOperations.save(disk);
        }else{
            throw new IllegalStateException("There's already a folder with this name here");
        }

        /*
        Update update = new Update().set("folders",folder);
        UpdateResult updateResult = mongoTemplate.upsert(query, update,Folder.class);
        //var diskRepository1 = diskRepository.findByItem(name);
        System.out.println(updateResult);

        //diskRepository.findBy({ name:"C" });
        //diskRepository.save();*/
        return "";
    }
}
