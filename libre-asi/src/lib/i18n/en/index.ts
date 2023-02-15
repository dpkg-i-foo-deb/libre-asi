import type { BaseTranslation } from '../i18n-types';
import welcome from './welcome';
import settings from './settings';
import general from './general';
import navMenu from './navMenu';
import cannotConnect from './cannotConnect';

const en = {
	welcome,
	settings,
	general,
	navMenu,
	cannotConnect
} satisfies BaseTranslation;

export default en;
