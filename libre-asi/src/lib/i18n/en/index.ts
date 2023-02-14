import type { BaseTranslation } from '../i18n-types';
import welcome from './welcome';
import settings from './settings';
import general from './general';
import navMenu from './navMenu';

const en = {
	welcome,
	settings,
	general,
	navMenu
} satisfies BaseTranslation;

export default en;
