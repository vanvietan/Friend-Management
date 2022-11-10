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

INSERT INTO public.relationships
(id, requester_id, addressee_id, "type")
VALUES
    (1, 101, 102, 'Friend'),
    (2, 102, 101, 'Friend'),
    (3, 103, 101, 'Subscribed'),
    (4, 101, 104, 'Blocked'),
    (5, 104, 101, 'Blocked');

