CREATE TABLE "Client" (
	"subscription_id" SERIAL PRIMARY KEY NOT NULL,
	"client_name" VARCHAR(50) NOT NULL,
	"client_second_name" VARCHAR(50) NOT NULL,
	"client_third_name" VARCHAR(50) NOT NULL,
	"sex" VARCHAR(10) NOT NULL,
	"birthdate" DATE NOT NULL,
	"height" FLOAT NOT NULL,
	"weight" FLOAT NOT NULL,
	"subscription_begin" DATE NOT NULL,
	"subscription_end" DATE NOT NULL
);
