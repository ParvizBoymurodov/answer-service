
select c.id,c.name,a.questions,a.answers from categories c inner join answers a on c.id = a.id_category where c.removed =false and a.removed= false;
SELECT a.id, a.questions, a.answers, c.name FROM answers a inner join categories c on a.id_category = c.id WHERE questions  LIKE '%к' AND a.removed = FALSE and c.removed;
Update categories SET name = 'asasxasxa' WHERE id = 2;