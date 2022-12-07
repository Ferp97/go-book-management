CREATE TABLE book (
   "id" SERIAL PRIMARY KEY NOT NULL,
   "name" varchar(50) NOT NULL,
   "author" varchar(50) NOT NULL,
   "genre" varchar(20) NOT NULL
);

INSERT INTO public.book(
	name, author, genre)
	VALUES ('Book1', 'Author1', 'Genre1');
