import type QuestionOption from "./Option";

export interface Question {
    special_id?: string;
    statement?: string;
    order?: Number;
    type?: string;
    category?: string;
    options?: QuestionOption[];

}