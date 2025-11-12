package service

const GetAllExpensesQuery string = `
	SELECT 
		expense_id, name, value, date, user_id, created_at, updated_at
	FROM 
		expenses;
`

const GetExpenseByIDQuery string = `
	SELECT 
		expense_id, name, value, date, user_id, created_at, updated_at 
	FROM 
		expenses 
	WHERE 
		expense_id = $1;
`

const CreateExpenseQuery = `
	INSERT INTO expenses 
		(expense_id, name, value, date, user_id, created_at, updated_at) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7);
`

const UpdateExpenseQuery string = `
	UPDATE expenses
	SET
		name = $2,
		value = $3,
		date = $4,
		user_id = $5,
		updated_at = $6
	WHERE
		expense_id = $1;
`

const DeleteExpenseQuery string = `
	DELETE FROM expenses
	WHERE expense_id = $1;
`
