package com.fileExplorer.swarch_proy_fileExplorer;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface FolderRepository extends MongoRepository<Folder, String>{

    @Query("{name:?0}")
    Folder findOneByName(String name);
}
