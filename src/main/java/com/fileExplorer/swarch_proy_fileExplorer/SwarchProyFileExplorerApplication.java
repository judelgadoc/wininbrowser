package com.fileExplorer.swarch_proy_fileExplorer;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;

//import javax.management.Query;
import java.util.ArrayList;
import java.util.List;

@SpringBootApplication
public class SwarchProyFileExplorerApplication {

	public static void main(String[] args) {

			SpringApplication.run(SwarchProyFileExplorerApplication.class, args);
			System.out.println("Running");

	}

	@Bean
	CommandLineRunner runner(DiskRepository repository, MongoTemplate mongoTemplate){


		return args -> {
			String name = "F";
			ArrayList diskFolders = new ArrayList<Folder>();
			Disk disk = new Disk(512,name,diskFolders);
			/*usingMongoTemplate(repository,mongoTemplate,name, disk)

			repository.findDiskByName(name)
					.ifPresentOrElse(s -> {
						System.out.println("There's already a disk with this name");
					},()->{
						repository.insert(disk);
					});*/
		};

	}


	private void usingMongoTemplate(DiskRepository repository, MongoTemplate mongoTemplate, String name, Disk disk){
		Query query = new Query();

		query.addCriteria(Criteria.where("name").is(name));

		List<Disk>disks= mongoTemplate.find(query, Disk.class);

		if(!disks.isEmpty()){
			//throw new IllegalStateException("Repeated name in new disk");
			System.out.println("There's already a disk with this name");
		}else{
			repository.insert(disk);
		}
	}



}
