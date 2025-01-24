import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

/**
 * Injectable service provided at root level
 */
@Injectable()
export class AuthGuard implements CanActivate {

    /**
     * Injecting dependencies
     * @param router - Angular router for navigation
     * @param cookieService - Handle cookie storage values
     */
    constructor(private router: Router, private cookieService: CookieService) { }

    /**
     * Validate token existence in cookie storage
     * @returns boolean
     */
    async canActivate(): Promise<boolean> {
        const session = this.cookieService.get("access_token");
        if (session !== "") {
            return true;
        } else {
            this.router.navigate(['/login']);
            return false;
        }
    }

}