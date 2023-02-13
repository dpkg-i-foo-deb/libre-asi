// This file was auto-generated by 'typesafe-i18n'. Any manual changes will be overwritten.
/* eslint-disable */
import type { BaseTranslation as BaseTranslationType, LocalizedString } from 'typesafe-i18n'

export type BaseTranslation = BaseTranslationType
export type BaseLocale = 'en'

export type Locales =
	| 'de'
	| 'en'

export type Translation = RootTranslation

export type Translations = RootTranslation

type RootTranslation = {
	welcome: {
		/**
		 * W​e​l​c​o​m​e​ ​t​o​ ​L​i​b​r​e​-​A​S​I
		 */
		TITLE: string
		/**
		 * L​i​b​r​e​-​A​S​I​ ​i​s​ ​a​ ​s​o​f​t​w​a​r​e​ ​t​o​o​l​ ​t​o​ ​a​s​s​i​s​t​ ​t​h​e​ ​i​n​t​e​r​v​i​e​w​ ​p​r​o​c​e​s​s​ ​t​h​a​t​ ​i​n​v​o​l​v​e​s​ ​t​h​e​ ​A​d​d​i​c​t​i​o​n​ ​S​e​v​e​r​i​t​y​ ​I​n​d​e​x​ ​f​o​r​m
		 */
		PARAGRAPH_1: string
		/**
		 * U​s​e​ ​t​h​e​ ​s​e​s​s​i​o​n​ ​m​e​n​u​ ​o​n​ ​t​h​e​ ​u​p​p​e​r​ ​r​i​g​h​t​ ​c​o​r​n​e​r​ ​t​o​ ​l​o​g​ ​i​n
		 */
		PARAGRAPH_2: string
		/**
		 * V​i​s​i​t​ ​t​h​e
		 */
		VISIT: string
		/**
		 * G​i​t​H​u​b​ ​r​e​p​o​s​i​t​o​r​y
		 */
		GITHUB_REPOSITORY: string
		/**
		 * t​o​ ​k​n​o​w​ ​m​o​r​e​ ​a​n​d​ ​r​e​a​d​ ​t​h​e​ ​d​o​c​u​m​e​n​t​a​t​i​o​n
		 */
		KNOW: string
	}
}

export type TranslationFunctions = {
	welcome: {
		/**
		 * Welcome to Libre-ASI
		 */
		TITLE: () => LocalizedString
		/**
		 * Libre-ASI is a software tool to assist the interview process that involves the Addiction Severity Index form
		 */
		PARAGRAPH_1: () => LocalizedString
		/**
		 * Use the session menu on the upper right corner to log in
		 */
		PARAGRAPH_2: () => LocalizedString
		/**
		 * Visit the
		 */
		VISIT: () => LocalizedString
		/**
		 * GitHub repository
		 */
		GITHUB_REPOSITORY: () => LocalizedString
		/**
		 * to know more and read the documentation
		 */
		KNOW: () => LocalizedString
	}
}

export type Formatters = {}