    CREATE TABLE spaceship (
        id int not null auto_increment,
        name varchar(255) not null,
        class varchar(255) not null,
        crew int not null,
        image varchar(255) not null,
        value decimal(65,2) not null,
        status enum('operational', 'in repair', 'destroyed') not null,
        armaments text not null,
        PRIMARY KEY (id)
    );
