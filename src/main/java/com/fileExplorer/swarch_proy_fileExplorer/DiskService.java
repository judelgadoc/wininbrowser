package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@AllArgsConstructor
@Service
public class DiskService {

    private final DiskRepository diskRepository;

    public List<Disk> getAllDisks() {
        return diskRepository.findAll();
    }
}
