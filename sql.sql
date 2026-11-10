-- database yaratish
CREATE DATABASE library_db;

-- table yaratish
CREATE TABLE books (
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

--book qo'shish
INSERT INTO books (title, author, price) 
VALUES ('Clean Code', 'Robert C. Martin', 150000.00);

-- bookni o'chirish
DELETE FROM books WHERE id = 2;

-- Bookni yangilash
UPDATE books SET price = 200000 WHERE id = 1;

-- price bo'yicha saralash (O'sish tartibida)
SELECT * FROM books ORDER BY price ASC;

-- price bo'yicha saralash (Kamayish tartibida)
SELECT * FROM books ORDER BY price DESC;

-- eng qimmat kitobni chiqarish
SELECT * FROM books ORDER BY price DESC LIMIT 1;

-- eng qimmat kitoblardan dastlabki 10 tasini tashlab, keyingi 5 tasini chiqarish
SELECT * FROM books ORDER BY price DESC OFFSET 10 LIMIT 5;