import { Routes } from '@angular/router';
import { ProtectedComponent } from './protected.component';

export const routes: Routes = [
    {
        path: '',
        component: ProtectedComponent,
        children: [
          {
            path: 'rx-claims',
            loadComponent: () => import('./rx-claims/rx-claims.component').then(m => m.RxClaimsComponent),
          },
          {
            path: 'benefits',
            loadComponent: () => import('./benifits/benifits.component').then(m => m.BenifitsComponent),
          },
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
            redirectTo: 'rx-claims',
            pathMatch: 'full'
          }
        ]
      },
]