CREATE TABLE article (
	id uuid DEFAULT uuid_generate_v4 (),
	author TEXT,
	title TEXT UNIQUE NOT NULL,
	body TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL
);

INSERT INTO article (id,author,title,body,created_at) VALUES
	 ('e8357d90-252f-47cb-ac66-48852f2b3cc8','Leonardo','Hello','World','2022-06-23 20:00:02'),
	 ('d6650ba6-0cc5-4da9-abad-1e664a0e3cec','Leonardo','Hi','Hello','2022-06-23 20:00:25'),
	 ('7f300573-2489-434c-b057-31049a7bef54','Testing','Article','Lorem Ipsum','2022-06-23 20:00:44');
