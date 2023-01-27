export function getCookie(name: string) {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) {
		const popped = parts.pop() ?? '';

		return popped.split(';').shift();
	}
}
