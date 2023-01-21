export interface AuthCookie {
    name:string,
    value:string,
    path:string,
    domain:string,
    max_age:number,
    expires:Date,
    secure:boolean,
    http_only:boolean,
    same_site:string,
    session_only:boolean,
}