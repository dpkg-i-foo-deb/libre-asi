export async function fetchNoRefresh(url: string, options?: RequestInit): Promise<Response> {
    const headers = new Headers(options?.headers);

    headers.append('Accept', 'application/json');
    headers.append('Content-Type', 'application/json');

    options = {
        ...options,
        headers,
        credentials: 'include',
        mode: 'cors'
    };

    try {
        const response = await fetch(url, options);

        return response;
    } catch (e) {
        return new Response('Service unavailable', { status: 503 });
    }
}