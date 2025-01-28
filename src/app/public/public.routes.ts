import { Routes } from '@angular/router';
import { PublicComponent } from './public.component';

export const routes: Routes = [
    {
      path: 'public',
      component: PublicComponent,
      children: [
        {
          path: 'pharmacies',
          loadComponent: () => import('./pharmacies/pharmacies.component').then(m => m.PharmaciesComponent),
        },
        {
          path: 'policy',
          loadComponent: () => import('./policies/policies.component').then(m => m.PoliciesComponent),
        },
        {
          path: 'request-pa',
          loadComponent: () => import('./request-pa/request-pa.component').then(m => m.RequestPaComponent),
        },
        {
          path: '',
          loadComponent: () => import('./home/home.component').then(m => m.HomeComponent),
        }
      ]
    },
    {
      path: '',
      redirectTo: 'public',
      pathMatch: 'full'
    }
  ];