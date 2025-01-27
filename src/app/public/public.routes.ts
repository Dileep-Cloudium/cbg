import { Routes } from '@angular/router';
import { PublicComponent } from './public.component';

export const routes: Routes = [
    {
      path: 'public',
      component: PublicComponent,
      children: [
        {
          path: 'pharmacies',
          loadComponent: () => import('../shared/pharmacies/pharmacies.component').then(m => m.PharmaciesComponent),
        },
        {
          path: 'policy',
          loadComponent: () => import('../shared/policies/policies.component').then(m => m.PoliciesComponent),
        },
        {
          path: 'request-pa',
          loadComponent: () => import('../shared/request-pa/request-pa.component').then(m => m.RequestPaComponent),
        },
        {
          path: '',
          loadComponent: () => import('./landing-page/landing-page.component').then(m => m.LandingPageComponent),
        }
      ]
    },
    {
      path: '',
      redirectTo: 'public',
      pathMatch: 'full'
    }
  ];