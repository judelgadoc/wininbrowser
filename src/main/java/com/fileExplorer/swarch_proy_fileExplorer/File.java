package com.fileExplorer.swarch_proy_fileExplorer;

import lombok.Data;
import org.bson.BsonBinarySubType;
import org.bson.types.Binary;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;

@Data
@Document
public class File{
    @Id
    private String id;
    private String type;
    private String name;
    private double size;
    //private Binary file;

    public File(String type, String name, double size) {
        this.type = type;
        this.name = name;
        this.size = size;
    }

    /*
    public File(String type) throws IOException {
        this.type = file.getContentType();
        this.name = file.getOriginalFilename();
        this.size = file.getSize();
        //this.file = new Binary(BsonBinarySubType.BINARY, file.getBytes());
    }*/

    public String getName() {
        return name;
    }

    public void setFile(Binary file) {
        //this.file = file;
    }
}
