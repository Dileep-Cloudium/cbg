import { Component } from '@angular/core';
import { AppBarModule } from '@syncfusion/ej2-angular-navigations';
import { ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { Router, RouterModule } from '@angular/router';

@Component({
  selector: 'app-public',
  standalone: true,
  imports: [AppBarModule, ButtonModule, RouterModule],
  templateUrl: './public.component.html',
  styleUrl: './public.component.css'
})
export class PublicComponent {

  constructor(private router: Router) {}

  navigate(path: string) {
    console.log('navigate', path);
    
    this.router.navigate([path]);
  }
}
