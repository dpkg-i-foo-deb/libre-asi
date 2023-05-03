export interface Interview {
    id?: number;
    startDate?: Date;
    endDate?: Date;
    pauseAt?: Date;
    resumedAt?: Date;
    patientID?: Number;
    interviewerID?: Number;
    asiFormID?: Number;
    currentQuestion?: string;


}
export default Interview;