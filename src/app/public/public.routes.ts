import { Routes } from '@angular/router';

export const routes: Routes = [
    {
      path: 'login',
      loadComponent: () => import('./login/login.component').then(m => m.LoginComponent),
    },
    {
      path: 'register',
      loadComponent: () => import('./register/register.component').then(m => m.RegisterComponent),
    },
    {
      path: '',
      loadComponent: () => import('./public.component').then(m => m.PublicComponent),
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
          path: '',
          redirectTo: 'dashboard',
          pathMatch: 'full'
        }
      ]
    }
  ];