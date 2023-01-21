import type { AuthCookie } from "./AuthCookie";

export interface JwtPair{
    access_token:AuthCookie
    refresh_token:AuthCookie
}