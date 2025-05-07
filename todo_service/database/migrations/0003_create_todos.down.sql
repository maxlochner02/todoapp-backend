CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       completed BOOLEAN DEFAULT FALSE,
                       user_id INTEGER NOT NULL,
                       project_id INTEGER REFERENCES projects(id),
                       category_id INTEGER REFERENCES categories(id),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP
);