package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.AllArgsConstructor;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface DiskRepository extends MongoRepository<Disk, String> {

    Optional<Disk> findDiskByName(String name);

    @Query("{name:?0}")
    Disk findDiskByNamed(String name); //Maybe this needs a query

    @Query("{name:?0}")
    Disk findByFolderName(String name);



}
