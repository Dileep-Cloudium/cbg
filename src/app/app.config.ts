import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { registerLicense } from '@syncfusion/ej2-base';
import { environment } from '../environments/environment.develop';
import { HTTP_INTERCEPTORS, provideHttpClient } from '@angular/common/http';
import { ErrorHandler } from '@angular/core';
import { GlobalErrorHandler } from './core/handler/globalErrorHandler';
import { AuthGuard } from './core/guards/auth.guard';
import { AppService } from './app.service';
import { Interceptor } from './core/interceptor/interceptor.interceptor';

registerLicense(environment.syncfusionLicenseKey);
export const appConfig: ApplicationConfig = {
  providers: [
    AuthGuard,
    AppService,
    provideRouter(routes),
    provideHttpClient(),
    provideZoneChangeDetection({ eventCoalescing: true }),
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
    { provide: ErrorHandler, useClass: GlobalErrorHandler },
  ]
};
