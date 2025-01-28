import { Component, ViewChild } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AppService } from './app.service';
import { ToasterComponent } from './shared/toaster/toaster.component';
import { CommonModule } from '@angular/common';

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

  constructor(public appService: AppService) {
  }

  ngOnInit() {
    this.appService.openToasterEmit.subscribe(() => {
      this.toasterComponent?.onCreate();
    })
  }
}
