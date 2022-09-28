CREATE TABLE "Trainer" (
	"trainer_id" SERIAL PRIMARY KEY ,
	"trainer_name" VARCHAR(50) NOT NULL,
	"trainer_second_name" VARCHAR(50) NOT NULL,
	"trainer_third_name" VARCHAR(50) NOT NULL,
	"trainer_phone" VARCHAR(20) NOT NULL,
	"speciality" VARCHAR(50) NOT NULL
);