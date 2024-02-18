  CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
  );

  CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
  );

  CREATE TABLE news_articles(
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    content VARCHAR NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
  );

  CREATE TABLE custom_pages(
    id BIGSERIAL PRIMARY KEY,
    custom_url VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
  );

  CREATE TABLE comments(
    id BIGSERIAL PRIMARY KEY,
    news_id BIGINT NOT NULL,
    name varchar NOT NULL,
    content VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
  );



  INSERT INTO users(username, email, password, created_at, updated_at)
  VALUES
  ('Steph', 'steph@gmail.com', '$2y$10$NMyl.lizYxdBp3L38z6Y7O6PxESHIbTVNAlYpIdndN.W7j9qu5Oia', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('Dray', 'dray@gmail.com', '$2y$10$bQqFIJRCsqYvQI58qL3AsOOAoN29TU9qaYxgZoNE5gZb0zo/MeRma', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('Klay', 'klay@gmail.com', '$2y$10$mns19dBq6TbiUdihhb4L9ufew5wyxw6TGbXClghgKveB9nErQG5RG', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('Wemby', 'wemby@gmail.com', '$2y$10$vBZM/2AX0JUCLFycTvj/z.uD52ruFso.CTgXCqFFt/RgHJrPjB6FW', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('James', 'james@gmail.com', '$2y$10$Fpa12nyRB6qVRv1YTaPUeOuXLcKKGGVelJmlxeV2FWIm3z5JoxZpe', '2022-06-10 15:00:00', '2022-06-10 15:00:00');

  INSERT INTO categories(name, created_at, updated_at)
  VALUES
  (`politics`, '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('football', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('basketball', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('game', '2022-06-10 15:00:00', '2022-06-10 15:00:00'),
  ('celebrity', '2022-06-10 15:00:00', '2022-06-10 15:00:00');