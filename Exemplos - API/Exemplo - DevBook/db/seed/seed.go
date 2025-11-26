package seed

func Create() {
	query := `
		create table users (
			id int auto_increment primary key,
			name varchar(50) not null,
			email varchar(50) not null
			) engine=innodb;
		`
}
