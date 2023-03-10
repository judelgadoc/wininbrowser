package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("disk/")
@AllArgsConstructor
public class DiskController {
    private final DiskService diskService;

    @GetMapping
    public List<Disk> fetchAllDisks(){
        return diskService.getAllDisks();
    }
}
