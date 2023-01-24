export interface Notification {
	kind: undefined | 'info' | 'warning' | 'error' | 'success';
	timeout: number;
	title: string;
	subtitle: string;
	caption: string;
	visible: boolean;
}
