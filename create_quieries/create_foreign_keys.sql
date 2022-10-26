ALTER TABLE "Group_client"
    ADD FOREIGN KEY ("subscription_id")
        REFERENCES "Client"("subscription_id");

ALTER TABLE "Group_client"
    ADD FOREIGN KEY ("group_id")
        REFERENCES "Group"("group_id");

ALTER TABLE "Group"
    ADD FOREIGN KEY ("trainer_id")
        REFERENCES "Trainer"("trainer_id");

ALTER TABLE "Group"
    ADD FOREIGN KEY ("program_id")
        REFERENCES "Program"("program_id");

ALTER TABLE "Times"
    ADD FOREIGN KEY ("program_id")
        REFERENCES "Program"("program_id");

ALTER TABLE "Timetable"
    ADD FOREIGN KEY ("time_id")
        REFERENCES "Times"("time_id");

ALTER TABLE "Timetable"
    ADD FOREIGN KEY ("group_id")
        REFERENCES "Group"("group_id");
