import type { Answer } from './Answer';
import type QuestionOption from './Option';

export interface Question {
	id?: number;
	special_id?: string;
	statement?: string;
	order?: number;
	type?: string;
	category?: string;
	options?: QuestionOption[];
	answers: Answer[];
	valid: boolean;
}
