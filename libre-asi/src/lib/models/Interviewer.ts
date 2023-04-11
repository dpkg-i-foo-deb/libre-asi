interface Interviewer {
	ID?: number;
	CreatedAt?: Date;
	UpdatedAt?: Date;
	email: string;
	username: string;
	password: string;
	firstName: string;
	lastName: string;
	firstSurName: string;
	lastSurName: string;
	birthdate: Date;
	age: number;
	personalID: string;
}

export default Interviewer;
