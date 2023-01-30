export default function emptyValidator(value: string): boolean {
	if (value === '' || value == undefined) {
		return false;
	}
	return true;
}
