package com.fileExplorer.swarch_proy_fileExplorer;


import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/folder")
@AllArgsConstructor
public class FolderController {

    public String newFolder(@ModelAttribute Folder folder){

        return "";
    }
}
