export default interface Patient {

    ID?: number;
    CreatedAt?: Date;
    UpdatedAt?: Date;
    email: string;
    username: string;
    password: string;
    firstName: string;
    lastName: string;
    firstSurname: string;
    secondSurname: string;
    birthdate: Date;
    personalID: string;

}