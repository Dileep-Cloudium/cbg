import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { registerLicense } from '@syncfusion/ej2-base';
import { environment } from '../environments/environment.develop';
import { provideHttpClient } from '@angular/common/http';

registerLicense(environment.syncfusionLicenseKey);
export const appConfig: ApplicationConfig = {
  providers: [provideZoneChangeDetection({ eventCoalescing: true }), provideRouter(routes),
    provideHttpClient()
  ]
};
