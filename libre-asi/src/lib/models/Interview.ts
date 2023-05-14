export interface Interview {
	id?: number;
	startDate?: Date;
	endDate?: Date;
	pauseAt?: Date;
	resumedAt?: Date;
	patientID?: number;
	interviewerID?: number;
	asiFormID?: number;
	currentQuestion?: string;
}
export default Interview;
