export interface Question {
    specialID?: string;
    statement?: string;
    order?: Number;
    type?: string;
    category?: string;
    options?: Array<typeof Option>;

}