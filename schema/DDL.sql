CREATE TABLE `bookingtogo`.`customer` ( 
    `cst_id` INT NOT NULL AUTO_INCREMENT ,  
    `nationality_id` INT NOT NULL ,  
    `cst_name` CHAR(50) NOT NULL ,  
    `cst_dob` DATE NOT NULL ,  
    `cst_phoneNum` VARCHAR(20) NOT NULL ,  
    `cst_email` VARCHAR(50) NOT NULL ,    
    PRIMARY KEY  (`cst_id`)
) ENGINE = InnoDB;

CREATE TABLE `bookingtogo`.`family_list` ( 
    `fl_id` INT NULL AUTO_INCREMENT ,  
    `cst_id` INT NOT NULL ,  
    `fl_relation` VARCHAR(50) NOT NULL , 
    `fl_name` VARCHAR(20) NOT NULL ,  
    `fl_dob` DATE NOT NULL ,    
    PRIMARY KEY  (`fl_id`)
) ENGINE = InnoDB;

CREATE TABLE `bookingtogo`.`nationality` ( 
    `nationality_id` INT NULL AUTO_INCREMENT ,
    `nationality_name` VARCHAR(50) NOT NULL , 
    `nationality_code` CHAR(2) NOT NULL ,      
    PRIMARY KEY  (`nationality_id`)
) ENGINE = InnoDB;