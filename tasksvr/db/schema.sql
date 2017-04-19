create table tasks (
	ID serial primary key,
	title varchar(255),
	createdAt timestamp,
	modifiedAt timestamp,
	complete bool
);

create table tags (
	taskID integer references tasks(ID),
	tag varchar(64)
);
