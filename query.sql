SELECT 
	u.id, 
	u.name,
	e.name parent_name
FROM 
	users u
LEFT JOIN
	users e ON u.parent_id=e.id
ORDER BY
	u.id;