ALTER TABLE categories
ADD CONSTRAINT unique_category_name  UNIQUE (name);
