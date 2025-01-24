import { Routes } from '@angular/router';
import { AuthGuard } from './core/guards/auth.guard';

export const routes: Routes = [
    {
      path: '',
      loadChildren: () => import('./protected/protected.routes').then(m => m.routes),
      canActivate: [AuthGuard]
    },
    {
      path: '',
      loadChildren: () => import('./public/public.routes').then(m => m.routes),
    },
    {
      path: '**',
      redirectTo: '',
    },
  ]; 