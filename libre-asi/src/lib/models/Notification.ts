export interface Notification {
	kind: 'info' | 'error' | 'warning' | 'success';
	timeout: number;
	title: string;
	subtitle: string;
	caption: string;
}
