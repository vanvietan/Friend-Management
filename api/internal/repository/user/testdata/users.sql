TRUNCATE table users CASCADE;
TRUNCATE table relationships CASCADE;


INSERT INTO public.users
(id, email)
VALUES
    (101, 'van1@gmail.com'),
    (102, 'van2@gmail.com'),
    (103, 'van3@gmail.com'),
    (104, 'van4@gmail.com'),
    (105, 'van5@gmail.com');

