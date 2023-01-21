export interface Session {
	active: boolean;
	role: number;
}

export enum SessionRole {
	Admin,
	Interviewer,
	None
}
