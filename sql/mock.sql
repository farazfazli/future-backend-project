INSERT INTO member(id, first_name, last_name) VALUES('2d793328-0c42-4f10-8475-a0c10ae3fdb7', 'John', 'Smith');
INSERT INTO member(id, first_name, last_name) VALUES('9a8671d9-78ad-4f2a-80a3-45bb087f8be3', 'Testfirst', 'Testlast');
INSERT INTO trainer(id, first_name, last_name) VALUES('1391b756-cf01-4fd4-b0ae-eae8cfe2f46f', 'Marty', 'McFly');
INSERT INTO trainer(id, first_name, last_name) VALUES('f7ba4208-fd31-4128-b733-5c471cee820c', 'Jane', 'Smith');
-- Our table constraints ensure that sample values fall within the specified guidelines (M-F, 8-5, 30 min duration)
INSERT INTO appointments(user_id, trainer_id, starts_at, ends_at) VALUES('2d793328-0c42-4f10-8475-a0c10ae3fdb7', '1391b756-cf01-4fd4-b0ae-eae8cfe2f46f', '2022-11-15T09:00:00-08:00', '2022-11-15T09:30:00-08:00');
INSERT INTO appointments(user_id, trainer_id, starts_at, ends_at) VALUES('9a8671d9-78ad-4f2a-80a3-45bb087f8be3', 'f7ba4208-fd31-4128-b733-5c471cee820c', '2022-11-15T09:00:00-08:00', '2022-11-15T09:30:00-08:00');
INSERT INTO appointments(user_id, trainer_id, starts_at, ends_at) VALUES('9a8671d9-78ad-4f2a-80a3-45bb087f8be3', 'f7ba4208-fd31-4128-b733-5c471cee820c', '2022-11-15T09:30:00-08:00', '2022-11-15T10:00:00-08:00');
