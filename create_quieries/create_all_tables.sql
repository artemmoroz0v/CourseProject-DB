CREATE TABLE Client (
    subscription_id serial NOT NULL,
    client_second_name varchar(50) NOT NULL,
    client_name varchar(50) NOT NULL,
    client_third_name varchar(50) NOT NULL,
    sex varchar(10) NOT NULL,
    birthdate DATE NOT NULL,
    height FLOAT NOT NULL,
    weight FLOAT NOT NULL,
    subscription_begin DATE NOT NULL,
    subscription_end DATE NOT NULL,
    CONSTRAINT Client_pk PRIMARY KEY (subscription_id)
);



CREATE TABLE Trainer (
    trainer_id serial NOT NULL,
    trainer_second_name varchar(50) NOT NULL,
    trainer_name varchar(50) NOT NULL,
    trainer_third_name varchar(50) NOT NULL,
    trainer_phone varchar(20) NOT NULL,
    CONSTRAINT Trainer_pk PRIMARY KEY (trainer_id)
);



CREATE TABLE Program (
    program_id serial NOT NULL,
    program_name varchar(50) NOT NULL,
    CONSTRAINT Program_pk PRIMARY KEY (program_id)
);



CREATE TABLE FC_Group (
    group_id serial NOT NULL,
    program_id integer NOT NULL,
    notes varchar(255) NOT NULL,
    trainer_id integer NOT NULL,
    clients_amount integer NOT NULL,
    CONSTRAINT Group_pk PRIMARY KEY (group_id)
);



CREATE TABLE Timetable (
    time_id integer NOT NULL UNIQUE,
    group_id integer NOT NULL
);



CREATE TABLE Group_client (
    group_id integer NOT NULL,
    subscription_id integer NOT NULL
);



CREATE TABLE Times (
    time_id serial NOT NULL,
    weekday integer NOT NULL,
    training_time TIME NOT NULL,
    program_id integer NOT NULL,
    CONSTRAINT Times_pk PRIMARY KEY (time_id)
);

