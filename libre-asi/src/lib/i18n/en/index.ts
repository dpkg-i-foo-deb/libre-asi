import type { BaseTranslation } from '../i18n-types';
import welcome from './welcome';
import settings from './settings';
import general from './general';

const en = {
	welcome,
	settings,
	general
} satisfies BaseTranslation;

export default en;
