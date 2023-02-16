import type { BaseTranslation } from '../i18n-types';
import welcome from './welcome';
import settings from './settings';
import general from './general';
import navMenu from './navMenu';
import cannotConnect from './cannotConnect';
import login from './login'

const en = {
	welcome,
	settings,
	general,
	navMenu,
	cannotConnect,
	login
} satisfies BaseTranslation;

export default en;
