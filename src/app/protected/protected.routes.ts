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
            loadComponent: () => import('./pharmacies/pharmacies.component').then(m => m.PharmaciesComponent),
          },
          {
            path: 'policy',
            loadComponent: () => import('./policies/policies.component').then(m => m.PoliciesComponent),
          },
          {
            path: 'home',
            loadComponent: () => import('./home/home.component').then(m => m.HomeComponent),
          },
          {
            path: '',
            redirectTo: 'home',
            pathMatch: 'full',
          },
        ]
      },
]