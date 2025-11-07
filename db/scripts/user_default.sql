INSERT INTO users (user_id, name, email, username, password_hash, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'Master User',
    'system@local',
    'master_user',
    '$2b$12$p2JcrWa/8GnnVhplE460JuXeFH6N555Axe7yOanWfv8Z8VSgnZA4O',
    NOW(),
    NOW()
);