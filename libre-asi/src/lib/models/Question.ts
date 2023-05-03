import type QuestionOption from "./Option";

export interface Question {
    specialID?: string;
    statement?: string;
    order?: Number;
    type?: string;
    category?: string;
    options?: QuestionOption[];

}