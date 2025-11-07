package service

const CreateUserQuery string = `
	INSERT into users 
		(user_id, name, email, username, password_hash, created_at, updated_at)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7);
`

const GetUserByIDQuery string = `
	SELECT 
		user_id, name, email, username, created_at, updated_at
	FROM 
		users
	WHERE
		user_id = $1;
`

const GetUserForLoginQuery string = `
	SELECT 
		user_id, name, email, username, password_hash, created_at, updated_at
	FROM 
		users
	WHERE
		username = $1
	OR
		email = $1;
`
