CREATE TABLE "Group" (
	"group_id" SERIAL PRIMARY KEY NOT NULL,
	"group_name" VARCHAR(50) NOT NULL,
	"program_id" VARCHAR(50) NOT NULL,
	"notes" VARCHAR(255) NOT NULL
);
