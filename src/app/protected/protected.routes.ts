import { Routes } from '@angular/router';

export const routes: Routes = [
    {
        path: '',
        loadComponent: () => import('./protected.component').then(m => m.ProtectedComponent),
        children: [
          {
            path: 'dashboard',
            loadComponent: () => import('../shared/dashboard/dashboard.component').then(m => m.DashboardComponent),
          },
          {
            path: 'policy',
            loadComponent: () => import('../shared/policies/policies.component').then(m => m.PoliciesComponent),
          },
          {
            path: 'users',
            loadComponent: () => import('../shared/policies/policies.component').then(m => m.PoliciesComponent),
          },
          {
            path: '',
            redirectTo: 'dashboard',
            pathMatch: 'full'
          }
        ]
      },
]