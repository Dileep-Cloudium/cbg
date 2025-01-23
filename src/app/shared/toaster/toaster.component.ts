import { Component, Input, ViewChild } from '@angular/core';
import { ButtonComponent, ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { ToastCloseArgs, ToastComponent, ToastModule } from '@syncfusion/ej2-angular-notifications';
import { ToasterModel } from './toaster.model';

@Component({
  selector: 'app-toaster',
  templateUrl: './toaster.component.html',
  standalone: true,
  imports: [ToastModule, ButtonModule],
  styleUrls: ['./toaster.component.css']
})
export class ToasterComponent {

  @ViewChild('toasttype')
  private toastObj: ToastComponent | undefined;

  @ViewChild('hideTosat')
  private hidebtn: ButtonComponent | undefined;

  @Input() toasterObj: ToasterModel | undefined;

  public onCreate(): void {
    setTimeout(() => {
      const styleName = this.toasterObj?.type === "error" ? "danger" : this.toasterObj?.type;
      this.toastObj?.show({
        title: this.toasterObj?.title, position: this.toasterObj?.position, content: this.toasterObj?.message, cssClass: 'e-toast-' + styleName, icon: 'toast-icons e-' + this.toasterObj?.type
      });
    }, 100);
  }

  public onclose(e: ToastCloseArgs): void {
    if (e.toastContainer.childElementCount === 0) {
      // this.hidebtn.element.style.display = 'none';
    }
  }

  public onBeforeOpen(): void {
    // this.hidebtn.element.style.display = 'inline-block';
  }
}
