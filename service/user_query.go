package service

const CreateUserQuery string = `
	INSERT into users 
		(user_id, name, email, user_name, password, created_at, modified_at)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7);
`

const GetUserByIDQuery string = `
	SELECT 
		user_id, name, email, user_name, created_at, modified_at
	FROM 
		users
	WHERE
		user_id = $1;
`

const GetUserForLoginQuery string = `
	SELECT 
		user_id, name, email, user_name, created_at, modified_at
	FROM 
		users
	WHERE
		user_name = $1
	OR
		email = $1;
`
