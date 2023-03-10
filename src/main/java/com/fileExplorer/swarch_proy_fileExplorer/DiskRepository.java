package com.fileExplorer.swarch_proy_fileExplorer;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface DiskRepository extends MongoRepository<Disk, String> {
    Optional<Disk> findDiskByName(String name);

}
