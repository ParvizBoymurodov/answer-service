-- CREATE TABLE if not exists categories (
--       id BIGSERIAL PRIMARY KEY,
--       name TEXT NOT NULL,
--       removed BOOLEAN DEFAULT FALSE
-- );
--
-- CREATE TABLE if not exists answers (
--        id BIGSERIAL PRIMARY KEY,
--        questions    TEXT NOT NULL,
--        answers    TEXT NOT NULL ,
--        id_category BIGSERIAL references categories(id),
--        removed BOOLEAN DEFAULT FALSE
-- );


-- SELECT * FROM answers WHERE questions LIKE '%' and removed =false;
--
--
-- SELECT id FROM answers WHERE questions='Как  зовут?' AND answers='Меня зовут Парвиз';
-- CREATE TABLE if not exists answers (
--   id BIGSERIAL PRIMARY KEY,
--   questions    TEXT NOT NULL,
--   answers    TEXT NOT NULL ,
--   id_category BIGSERIAL references categories(id),
--   removed BOOLEAN DEFAULT FALSE
-- );
-- CREATE TABLE if not exists categories (
--                                           id BIGSERIAL PRIMARY KEY,
--                                           name TEXT NOT NULL,
--                                           removed BOOLEAN DEFAULT FALSE
-- );
--


select c.id,c.name,a.questions,a.answers from categories c inner join answers a on c.id = a.id_category where c.removed =false and a.removed= false;
SELECT a.id, a.questions, a.answers, c.name FROM answers a inner join categories c on a.id_category = c.id WHERE questions  LIKE '%к' AND a.removed = FALSE and c.removed;
Update categories SET name = 'asasxasxa' WHERE id = 2;