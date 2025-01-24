import { Component, ViewChild } from '@angular/core';
import { NavigationEnd, NavigationError, NavigationStart, Router, RouterOutlet } from '@angular/router';
import { AppService } from './app.service';
import { ToasterComponent } from './shared/toaster/toaster.component';
import { CommonModule } from '@angular/common';
import { filter } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  standalone: true,
  imports: [RouterOutlet, ToasterComponent, CommonModule]
})
export class AppComponent {

  /**
 * Toaster component from core module
 */
  @ViewChild('Toaster') public toasterComponent!: ToasterComponent;

  // constructor(public appService: AppService) { }
  constructor(public appService: AppService) {
  }
  ngOnInit() {
    console.log('AppComponent ngOnInit');
    this.appService.openToasterEmit.subscribe(() => {
      this.toasterComponent?.onCreate();
    })
  }
}
