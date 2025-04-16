ALTER TABLE users
ADD COLUMN email VARCHAR(255) DEFAULT 'unknown@example.com';

-- Ensure all rows have unique email values
UPDATE users
SET email = CONCAT('user_', id, '@example.com')
WHERE email IS NULL OR email = 'unknown@example.com';

-- Add the NOT NULL constraint
ALTER TABLE users
ALTER COLUMN email SET NOT NULL;

-- Add the UNIQUE constraint
ALTER TABLE users
ADD CONSTRAINT users_email_key UNIQUE (email);
