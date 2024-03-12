INSERT INTO teachers (email)
VALUES ('teacherken%40gmail.com'),
('teacherjoe%40gmail.com'),
('teacherken@gmail.com'),
('john.doe@example.com'),
('emily.smith@example.com'),
('michael.johnson@example.com'),
('sarah.wilson@example.com'),
('david.brown@example.com'),
('lisa.jones@example.com'),
('matthew.davis@example.com'),
('laura.miller@example.com'),
('james.taylor@example.com'),
('jessica.white@example.com');

INSERT INTO students (email, is_suspended)
VALUES ('studentjon@gmail.com', false),
('studenthon@gmail.com', false),
('commonstudent1@gmail.com', false),
('commonstudent2@gmail.com', false),
('student_only_under_teacher_ken@gmail.com', false),
('studentmary@gmail.com', false),
('studentbob@gmail.com', false),
('studentagnes@gmail.com', false),
('studentmiche@gmail.com', false);

INSERT INTO teacher_student (teacher_id, student_id)
VALUES (1,1),
(2,2),
(1,2),
(1,3),
(2,3);
