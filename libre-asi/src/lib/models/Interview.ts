export interface Interview {
    ID?: number;
    StartDate?: Date;
    EndDate?: Date;
    PauseAt?: Date;
    ResumedAt?: Date;
    PatientID?: number;
    InterviewerID?: number;
    AsiFormID?: number;
    CurrentQuestion?: string;


}
export default Interview;