const emailRegex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

export default function validateEmail(value: string): boolean {
	return emailRegex.test(value);
}
