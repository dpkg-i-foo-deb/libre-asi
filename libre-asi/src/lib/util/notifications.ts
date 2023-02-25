import { notifications } from '$lib/stores/notificationStore';
export function sendSuccess(title: string, subtitle: string) {
  notifications.set({
    kind: 'success',
    title: title,
    subtitle: subtitle,
    caption: new Date().toLocaleDateString(),
    timeout: 10000,
    visible: true
  });
}

export function sendError(title: string, subtitle: string) {
  notifications.set({
    kind: 'error',
    title: title,
    subtitle: subtitle,
    caption: new Date().toLocaleDateString(),
    timeout: 10000,
    visible: true
  });
}

export function sendInfo(title: string, subtitle: string) {
  notifications.set({
    kind: 'info',
    title: title,
    subtitle: subtitle,
    caption: new Date().toLocaleDateString(),
    timeout: 10000,
    visible: true
  })
}
