CREATE TABLE IF NOT EXISTS Goal(
    id serial primary key NOT NULL,
    name text NOT NULL,
    complete boolean NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS Entry(
    id serial primary key NOT NULL,
    description text,    
    complete boolean NOT NULL DEFAULT false,
    created_at timestamp DEFAULT 'now',
    goal_id int references Goal(id) ON DELETE CASCADE NOT NULL
);
